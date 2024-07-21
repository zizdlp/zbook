package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailTemplate(t *testing.T) {
	Title := "Reset ZBook Password"
	recipientName := "John Doe"
	verificationLink := "https://example.com/verify"
	emailSubject := "Verify Your Email Address"
	buttonText := "Verify Email"
	additionalText := "Please verify your email address by clicking the button below."

	// Use fmt.Sprintf to generate the email body
	emailBody := fmt.Sprintf(EmailTemplate, Title, recipientName, emailSubject, verificationLink, buttonText, additionalText)

	// Check if the generated email body contains the expected text
	require.Contains(t, emailBody, emailSubject)
	require.Contains(t, emailBody, recipientName)
	require.Contains(t, emailBody, additionalText)
	require.Contains(t, emailBody, verificationLink)
	require.Contains(t, emailBody, buttonText)
}
