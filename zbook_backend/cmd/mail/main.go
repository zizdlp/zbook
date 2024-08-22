package main

import (
	"fmt"
	"log"

	"github.com/zizdlp/zbook/mail"
	"github.com/zizdlp/zbook/util"
)

func main() {
	// 加载配置
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sender := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email from zbook"
	// 模拟用户数据
	user := struct {
		Username string
	}{
		Username: "admin",
	}
	verifyUrl := "http://localhost:3000/verify_email?verification_id=66ca2c9313264f449648a6e2aa6f8cf0"

	Title := "Verify Your Email Address"
	emailSubject := "Thank you for registering with us! Please verify your email address by clicking the button below:"
	buttonText := "Verify Email"
	additionalText := "If you did not register for an account, please ignore this email or contact support if you have any questions."
	base64Image, err := util.ReadImageBytesToBase64("./icons/logo.png")
	if err != nil {
		log.Fatalf("Failed to read image file: %v", err)
	}

	emailBody := fmt.Sprintf(util.EmailTemplate, Title, user.Username, emailSubject, verifyUrl, buttonText, additionalText, base64Image)

	to := []string{"zizdlp@gmail.com"}

	err = sender.SendEmail(subject, emailBody, to, nil, nil, nil, config.SmtpAuthAddress, config.SmtpServerAddress)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
}
