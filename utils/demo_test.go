package utils

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {

	bodyText := fmt.Sprintf("试用申请:\n  商户: %s, 手机号: %s, 邮箱: %s,币种名称: %s",
		"test", "1111111111", "123@qq.com", "BTC")
	em := EmailConfig{
		IamUserName:  "new-project@hoo.com",
		Recipient:    "473022457@qq.com",
		SmtpUsername: "AKIAZXPISQLLFWOJQDQ4",
		SmtpPassword: "BG/CV8RZPkXrtIde9ZONEO142PqfV+2lP+Fcsgq0pOMQ",
		Host:         "email-smtp.ap-northeast-1.amazonaws.com",
	}
	em.SendEmail(bodyText)
}
