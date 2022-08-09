package mail

import (
	"context"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

var yourDomain string = config.GetConf().EmailDomainUrl
var privateAPIKey string = config.GetConf().EmailSecretKey

type Email struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}
type EmailResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

func SendMail(email *Email) (*EmailResponse, error) {

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(email.Sender, email.Subject, email.Body, email.Recipient)
	//message.SetHtml("<html><body><h1>Testing some Mailgun Awesomeness!!</h1></body></html>")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		return nil, err
	}

	return &EmailResponse{
		Message: resp,
		Id:      id,
	}, nil
}
