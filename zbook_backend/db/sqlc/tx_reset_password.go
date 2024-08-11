package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResetPasswordTxParams struct {
	VerificationUrl string
	Password        string
	Email           string
}

func (store *SQLStore) ResetPasswordTx(ctx context.Context, arg ResetPasswordTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		verification, err := q.GetVerification(ctx, arg.VerificationUrl)
		if err != nil {
			return status.Errorf(codes.Internal, "get verification failed:%v", err)
		}
		if verification.VerificationType != util.VerifyTypeResetPassword {
			return status.Errorf(codes.InvalidArgument, "invalid VerificationType: %s", verification.VerificationType)
		}

		_, err = q.MarkVerificationAsUsed(ctx, arg.VerificationUrl)
		if err != nil {
			return status.Errorf(codes.Internal, "mark verify as used failed:%v", err)
		}

		hashedPassword, err := util.HashPassword(arg.Password)
		if err != nil {
			return status.Errorf(codes.Internal, "password cannot be hashed: %s", err)
		}
		if arg.Email != verification.Email {
			return status.Errorf(codes.PermissionDenied, "email is not belong to this verification")
		}
		arg := UpdateUserBasicInfoParams{
			Username:       verification.Username,
			HashedPassword: pgtype.Text{String: hashedPassword, Valid: true},
		}

		_, err = q.UpdateUserBasicInfo(ctx, arg)
		if err != nil {
			if errors.Is(err, ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "user not exist: %s", err)
			}
			return status.Errorf(codes.Internal, "reset password failed: %s", err)
		}
		return err

	})
	return err
}
