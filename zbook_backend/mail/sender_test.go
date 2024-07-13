package mail

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestSendEmailWithGmail is ignored *****")
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email from @zizdlp.com"
	// 模拟用户数据
	user := struct {
		Username string
	}{
		Username: "admin",
	}
	verifyUrl := "http://localhost:3000/verify_email?verification_id=66ca2c9313264f449648a6e2aa6f8cf0"

	// 读取本地 HTML 文件内容
	htmlFilePath := "../email_verify_template.html"
	content, err := os.ReadFile(htmlFilePath)
	require.NoError(t, err)

	// 使用 fmt.Sprintf 插入变量
	finalContent := fmt.Sprintf(string(content), user.Username, verifyUrl, verifyUrl, verifyUrl)

	to := []string{"zizdlp@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, finalContent, to, nil, nil, attachFiles, config.SmtpAuthAddress, config.SmtpServerAddress)
	require.NoError(t, err)
}
