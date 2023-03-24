package mail

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/mailgun/mailgun-go/v4"
	"html/template"
	"strings"
	"time"
)

type EmailType string

const (
	Auth         EmailType = "auth"
	Support      EmailType = "support"
	Subscription EmailType = "subscription"
	Budget       EmailType = "budget"
	Others       EmailType = "others"
)

var templateDir embed.FS

const (
	templatePath = "templates/"
	fileSuffix   = ".html"
)

type Email struct {
	Type      EmailType
	Subject   string
	Body      string
	Recipient string
	// Params to replace the template variables.
	Params interface{} `json:"params,omitempty"`

	// Glob represents which template to use in building the email
	EmailType EmailType `json:"template_name,omitempty"`
}
type EmailResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

func (t EmailType) String() string {
	return string(t)
}

type Mail struct {
	templ *template.Template
	body  bytes.Buffer
}

// Build TODO: glob pattern must not match more than one template
func (e *Mail) Build(glob string, params interface{}) error {
	templ, err := e.templ.ParseFS(templateDir, e.buildGlob(glob))
	if err != nil {
		return fmt.Errorf("failed to parse template fs: %v", err)
	}
	e.templ = templ

	err = e.templ.Execute(&e.body, params)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func (e *Mail) buildGlob(glob string) string {
	var s strings.Builder

	s.WriteString(templatePath)
	s.WriteString(glob)

	if !strings.HasSuffix(glob, fileSuffix) {
		s.WriteString(fileSuffix)
	}

	return s.String()
}

func SendMail(email *Email) error {
	yourDomain := config.GetConf().EmailDomainUrl
	privateAPIKey := config.GetConf().EmailSecretKey
	testSender := config.GetConf().EmailFrom

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(testSender, email.Subject, email.Body, email.Recipient)
	//body := GetEmailBody(email.Type, email.Data)
	//t, err := template.New("email").Parse(body)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var tpl bytes.Buffer
	//if err = t.Execute(&tpl, Email{
	//	Data: email.Data,
	//}); err != nil {
	//	return nil, err
	//}
	//
	//result := tpl.String()
	//if err != nil {
	//	return nil, err
	//}
	//
	//message.SetHtml(result)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	return nil
}
