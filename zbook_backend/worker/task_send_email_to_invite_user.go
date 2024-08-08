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

const TaskInviteUser = "task:invite_user"

type PayloadInviteUser struct {
	Email string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskInviteUser(
	ctx context.Context,
	payload *PayloadInviteUser,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskInviteUser, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskInviteUser(ctx context.Context, task *asynq.Task) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}
	var payload PayloadInviteUser
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		return fmt.Errorf("use already exist for this email: %s", user.Email)
	}

	invitation, err := processor.store.CreateInvitation(ctx, db.CreateInvitationParams{
		Email:         payload.Email,
		InvitationUrl: util.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify code:%w", err)
	}

	subject := "Invitation to Join ZBook"

	title := "Register to ZBook"
	emailSubject := "Thank you for your interest in ZBook! Please complete your registration by clicking the button below:"
	verifyUrl := fmt.Sprintf("%s/auth/register?invitation_url=%s", config.HOMEADDRESS, invitation.InvitationUrl)
	buttonText := "Complete Registration"
	additionalText := "If you did not request an invitation, please ignore this email or contact support if you have any questions."
	base64Image, err := util.ReadImageBytesToBase64("icons/logo.png")
	if err != nil {
		return fmt.Errorf("failed to read logo file: %w", err)
	}
	emailBody := fmt.Sprintf(util.EmailTemplate, title, user.Username, emailSubject, verifyUrl, buttonText, additionalText, base64Image)
	to := []string{payload.Email}

	err = processor.mailer.SendEmail(subject, emailBody, to, nil, nil, nil, config.SmtpAuthAddress, config.SmtpServerAddress)
	if err != nil {
		return fmt.Errorf("failed to send email to verify: %w", err)
	}
	// TODO: send email to user
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str(" email ", user.Email).Msg("processed task")
	return nil
}
