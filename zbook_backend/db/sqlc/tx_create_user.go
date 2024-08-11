package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zizdlp/zbook/util"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		if arg.CreateUserParams.Username == "admin" {
			arg_update_user_role := UpdateUserBasicInfoParams{
				Username: "admin",
				UserRole: pgtype.Text{String: util.AdminRole, Valid: true},
			}
			result.User, err = q.UpdateUserBasicInfo(ctx, arg_update_user_role)
			if err != nil {
				return err
			}
		}
		return arg.AfterCreate(result.User)
	})

	return result, err
}
