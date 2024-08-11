package gapi

import (
	"context"
	"database/sql"
	"fmt"

	"reflect"
	"testing"

	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/zizdlp/zbook/db/mock"
	"github.com/zizdlp/zbook/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	mockwk "github.com/zizdlp/zbook/worker/mock"
)

type eqCreateUserTxParamsMatcher struct {
	arg      db.CreateUserTxParams
	password string
}

func (expected eqCreateUserTxParamsMatcher) Matches(x interface{}) bool {
	actualArg, ok := x.(db.CreateUserTxParams) // 这个是真的外部传入的值。 expected是期待的值
	if !ok {
		return false
	}

	err := util.CheckPassword(expected.password, actualArg.HashedPassword)
	if err != nil {
		return false
	}

	expected.arg.HashedPassword = actualArg.HashedPassword
	return reflect.DeepEqual(expected.arg.CreateUserParams, actualArg.CreateUserParams)
}

func (e eqCreateUserTxParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserTxParams(arg db.CreateUserTxParams, password string) gomock.Matcher {
	return eqCreateUserTxParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := createRandomUser(t)

	testCases := []struct {
		name          string
		req           *rpcs.CreateUserRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *rpcs.CreateUserResponse, err error)
	}{
		{
			name: "InternalError",
			req: &rpcs.CreateUserRequest{
				Username: user.Username,
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.CreateUserTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *rpcs.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "OK",
			req: &rpcs.CreateUserRequest{
				Username: user.Username,
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserTxParams{
					CreateUserParams: db.CreateUserParams{
						Username: user.Username,
						Email:    user.Email,
					},
				}
				store.EXPECT().
					CreateUserTx(gomock.Any(), EqCreateUserTxParams(arg, password)).
					Times(1).
					Return(db.CreateUserTxResult{User: user}, nil)
			},
			checkResponse: func(t *testing.T, res *rpcs.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)

			},
		},
		{
			name: "DuplicateUsername",
			req: &rpcs.CreateUserRequest{
				Username: user.Username,
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.CreateUserTxResult{}, db.ErrUniqueViolation)

			},
			checkResponse: func(t *testing.T, res *rpcs.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.AlreadyExists, st.Code())
			},
		},
		{
			name: "InvalidEmail",
			req: &rpcs.CreateUserRequest{
				Username: user.Username,
				Password: password,
				Email:    "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, res *rpcs.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)

			taskCtrl := gomock.NewController(t)
			defer taskCtrl.Finish()
			taskDistributor := mockwk.NewMockTaskDistributor(taskCtrl)

			config, err := util.LoadConfig(".")
			if err != nil {
				return
			}
			redisClient := redis.NewClient(&redis.Options{
				Addr: config.RedisAddress, // Redis 服务器地址和端口
			})
			minioClient, err := storage.GetMinioClient()
			if err != nil {
				return
			}
			tc.buildStubs(store)
			server := newTestServer(t, store, taskDistributor, redisClient, minioClient)

			res, err := server.CreateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

func createRandomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}
	return
}
