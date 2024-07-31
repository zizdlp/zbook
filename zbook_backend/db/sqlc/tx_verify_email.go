package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (store *SQLStore) VerifyEmailTx(ctx context.Context, VerificationUrl string) (User, error) {
	var user User

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		verification, err := q.GetVerification(ctx, VerificationUrl)
		if err != nil {
			return status.Errorf(codes.Internal, "get verification failed:%v", err)
		}
		if verification.VerificationType != util.VerifyTypeVerifyEmail {
			return status.Errorf(codes.Internal, "invalid VerificationType")
		}

		_, err = q.MarkVerificationAsUsed(ctx, VerificationUrl)
		if err != nil {
			return status.Errorf(codes.Internal, "mark verification used failed: %s", err)
		}
		arg := UpdateUserBasicInfoParams{
			Username: verification.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		user, err = q.UpdateUserBasicInfo(ctx, arg)
		if err != nil {
			return status.Errorf(codes.Internal, "update user failed: %s", err)
		}
		return nil
	})
	return user, err
}
