package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/util"
)

const TaskVerifyEmail = "task:verify_email"

type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskVerifyEmail(
	ctx context.Context,
	payload *PayloadVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskVerifyEmail(ctx context.Context, task *asynq.Task) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}
	var payload PayloadVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	verification, err := processor.store.CreateVerification(ctx, db.CreateVerificationParams{
		VerificationUrl:  util.RandomString(32),
		UserID:           user.UserID,
		VerificationType: util.VerifyTypeVerifyEmail,
	})
	if err != nil {
		return fmt.Errorf("failed to create verify code:%w", err)
	}

	subject := "Welcome to ZBook!"

	Title := "Verify Your Email Address"
	emailSubject := "Thank you for registering with us! Please verify your email address by clicking the button below:"
	verifyUrl := fmt.Sprintf("%s/verify_email?verification_url=%s", config.HOMEADDRESS, verification.VerificationUrl)
	buttonText := "Verify Email"
	additionalText := "If you did not register for an account, please ignore this email or contact support if you have any questions."
	base64Image, err := util.ReadImageBytesToBase64("icons/logo.png")
	if err != nil {
		return fmt.Errorf("failed to read logo file: %w", err)
	}
	emailBody := fmt.Sprintf(util.EmailTemplate, Title, user.Username, emailSubject, verifyUrl, buttonText, additionalText, base64Image)
	to := []string{user.Email}

	err = processor.mailer.SendEmail(subject, emailBody, to, nil, nil, nil, config.SmtpAuthAddress, config.SmtpServerAddress)
	if err != nil {
		return fmt.Errorf("failed to send email to verify: %w", err)
	}
	// TODO: send email to user
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str(" email ", user.Email).Msg("processed task")
	return nil
}
