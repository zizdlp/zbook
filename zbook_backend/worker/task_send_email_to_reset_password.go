package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/util"
)

const TaskResetPassword = "task:reset_password"

type PayloadResetPassword struct {
	Email string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskResetPassword(
	ctx context.Context,
	payload *PayloadResetPassword,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskResetPassword, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskResetPassword(ctx context.Context, task *asynq.Task) error {
	var payload PayloadResetPassword
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	verification, err := processor.store.CreateVerification(ctx, db.CreateVerificationParams{
		VerificationID:   uuid.New(),
		UserID:           user.UserID,
		VerificationType: util.VerifyTypeResetPassword,
	})
	if err != nil {
		return fmt.Errorf("failed to create verify code:%w", err)
	}
	subject := "点击如下链接重置密码"
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}
	verifyUrl := fmt.Sprintf("%s/reset_password?verification_id=%s", config.HOMEADDRESS, util.UUIDToString(verification.VerificationID))
	htmlFilePath := "./email_reset_template.html"
	content, err := os.ReadFile(htmlFilePath)
	if err != nil {
		return fmt.Errorf("failed to get email verify template")
	}
	// 使用 fmt.Sprintf 插入变量
	finalContent := fmt.Sprintf(string(content), user.Username, verifyUrl, verifyUrl, verifyUrl)
	to := []string{user.Email}

	err = processor.mailer.SendEmail(subject, finalContent, to, nil, nil, nil, config.SmtpAuthAddress, config.SmtpServerAddress)
	if err != nil {
		return fmt.Errorf("failed to send Email to reset password: %w", err)
	}
	// TODO: send email to user
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str(" email ", user.Email).Msg("processed task")
	return nil
}
