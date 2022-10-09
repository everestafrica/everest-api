package channels

import (
	"context"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

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
	yourDomain := config.GetConf().EmailDomainUrl
	privateAPIKey := config.GetConf().EmailSecretKey
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(email.Sender, email.Subject, email.Body, email.Recipient)
	//body := `
	//<html>
	//<body>
	//	<h1>Sending HTML emails with Mailgun</h1>
	//	<p style="color:blue; font-size:30px;">Hello world</p>
	//	<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
	//</body>
	//</html>
	//`
	//
	//	message.SetHtml(body)
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
