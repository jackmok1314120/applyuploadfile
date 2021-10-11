package utils

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {

	bodyText := fmt.Sprintf("试用申请:\n  商户: %s, 手机号: %s, 邮箱: %s,币种名称: %s",
		"test", "1111111111", "123@qq.com", "BTC")
	em := EmailConfig{
		IamUserName:  "moyunrz@163.com",
		Recipient:    "laoqiu@hoo.com",
		SmtpUsername: "AKIAZXPISQLLCJ3AFQWD",
		SmtpPassword: "BAhIdWFxhya3k7TG3cz4Gl08zMfwNDYW+WJug95iaINO",
	}
	em.SendEmail(bodyText)
}
