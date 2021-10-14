package utils

import (
	"fmt"
	"gopkg.in/gomail.v2" //go get gopkg.in/gomail.v2
)

const (

	// ConfigSet The name of the configuration set to use for this message.
	// If you comment out or remove this variable, you will also need to
	// comment out or remove the header below.
	ConfigSet = "ConfigSet"
	// Host If you're using Amazon SES in an AWS Region other than US West (Oregon),
	// replace email-smtp.us-west-2.amazonaws.com with the Amazon SES SMTP
	// endpoint in the appropriate region.
	Host = "email-smtp.ap-northeast-1.amazonaws.com"

	Port = 587

	// Subject The subject line for the email.
	Subject = "HOO 商户申请提交，测试，无需理会"

	// HtmlBody
	// The HTML body for the email.
	HtmlBody = "<html><head><title>HOO 商户申请提交</title></head><body>" +
		"<h1>HOO 商户申请提交</h1>" +
		"<p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using " +
		"the <a href='https://github.com/go-gomail/gomail/'>Gomail " +
		"package</a> for <a href='https://golang.org/'>Go</a>.</p>" +
		"</body></html>"

	// Tags
	// The tags to apply to this message. Separate multiple key-value pairs
	// with commas.
	// If you comment out or remove this variable, you will also need to
	// comment out or remove the header on line 80.
	Tags = "genre=test,genre2=test2"
	// The character encoding for the email.
	CharSet = "UTF-8"
)

type EmailConfig struct {
	IamUserName  string `json:"iam_user_name"`
	Recipient    string `json:"recipient"`
	SmtpUsername string `json:"smtp_username"`
	SmtpPassword string `json:"smtp_password"`
	Host         string `json:"host"`
}

func (emailConfig *EmailConfig) SendEmail(TextBody string) {

	// Create a new message.
	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	//m.SetBody("text/html", HtmlBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", TextBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(emailConfig.IamUserName, emailConfig.IamUserName)},
		"To":      {emailConfig.Recipient},
		"Subject": {Subject},
		// Comment or remove the next line if you are not using a configuration set
		//"X-SES-CONFIGURATION-SET": {ConfigSet},
		// Comment or remove the next line if you are not using custom tags
		"X-SES-MESSAGE-TAGS": {Tags},
	})

	// Send the email.
	d := gomail.NewPlainDialer(emailConfig.Host, Port, emailConfig.SmtpUsername, emailConfig.SmtpPassword)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent!")
	}
}
