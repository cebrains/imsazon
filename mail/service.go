/*
	The mail service's purpose is to send emails.
	Some code courtesy of https://nathanleclaire.com/blog/2013/12/17/sending-email-from-gmail-using-golang/
 */
package mail

import (
	"errors"
	"gopkg.in/gomail.v2"
)

// EmailUser describes the login credentials for a mail server
type EmailUser struct {
	Username		string
	Password		string
	EmailServer		string
	Port			int
}

// add our own email user configuration
var imsazonMailConfig = &EmailUser{
	"office.imsazon@gmail.com",
	"QILzLpnLisnvFx2oHBEr",
	"smtp.gmail.com",
	587,
}

// SmtpTemplateData describes the template for simple email (one recipient per mail, no attachments)
type SmtpTemplateData struct {
	From			string
	To				string
	Subject			string
	Body			string
}

const emailTemplate = `From: &#123;&#123;.From&#125;&#125;
To: &#123;&#123;.To&#125;&#125;
Subject: &#123;&#123;.Subject&#125;&#125;

&#123;&#123;.Body&#125;&#125;

Sincerely,

&#123;&#123;.From&#125;&#125;
`

// ErrInvalidArgument is returned when on or more arguments are invalid
var ErrInvalidArgument = errors.New("Invalid argument")

// ErrParseTemplate is returned when the email template could not be parsed
var ErrParseTemplate = errors.New("Could not parse template")

// ErrExecuteTemplate is returned when the SMTP template could not be executed
var ErrExecuteTemplate = errors.New("Could not execute template")

// Service is the interface that provides the mail send method
type Service interface {
	Send(smtpTemplate SmtpTemplateData) error
}

type service struct {

}

func(s *service) Send(smtpTemplate SmtpTemplateData) error {
	/*if smtpTemplate == (SmtpTemplateData{}) {
		return ErrInvalidArgument
	}

	var err error
	var doc bytes.Buffer

	t := template.New("emailTemplate")
	t, err = t.Parse(emailTemplate)

	if err != nil {
		return ErrParseTemplate
	}

	err = t.Execute(&doc, smtpTemplate)

	if err != nil {
		return ErrExecuteTemplate
	}

	// email user auth information
	auth := smtp.PlainAuth(
		"",
		imsazonMailConfig.Username,
		imsazonMailConfig.Password,
		imsazonMailConfig.EmailServer,
	)

	err = smtp.SendMail(
		imsazonMailConfig.EmailServer + ":" + strconv.Itoa(imsazonMailConfig.Port),
		auth,
		imsazonMailConfig.Username,
		[]string{imsazonMailConfig.Username},
		doc.Bytes(),
	)

	if err != nil {
		return err
	}

	return nil*/

	m := gomail.NewMessage()
	m.SetHeader("From", "imsazon@micsti.at")
	m.SetHeader("To", "stifter.michael@gmx.net")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("w00e8823.kasserver.com", 587, "m042be71", "wennG7TFqE66cQA8")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func NewService() Service {
	return &service{

	}
}