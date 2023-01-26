package channels

import (
	"context"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/mailgun/mailgun-go/v4"
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

type Email struct {
	Type      EmailType
	Subject   string
	Body      string
	Recipient string
	Data      interface{}
}
type EmailResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
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

func GetEmailBody(emailType EmailType, data interface{}) string {
	var body string
	switch emailType {
	case Auth:
		body = `
	<html>
	<body>
		<h1>Authentication</h1>
		<p style="color:blue; font-size:30px;">Hello world</p>
		<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		<ul> 
 			{{range .Data}}
				<li>{{.Name}}</li>
				{{end}}
			{{end}}
		</ul>	
</body>
	</html>
	`
	case Support:
		body = `
	<html>
	<body>
		<h1>Sending HTML emails with Mailgun</h1>
		<p style="color:blue; font-size:30px;">Hello world</p>
		<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		<ul> 
 			{{range .Data}}
				<li>{{.Name}}</li>
				{{end}}
			{{end}}
		</ul>	
</body>
	</html>
	`
	case Subscription:
		body = `
	<html>
	<body>
		<h1>Sending HTML emails with Mailgun</h1>
		<p style="color:blue; font-size:30px;">Hello world</p>
		<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		<ul> 
 			{{range .Data}}
				<li>{{.Name}}</li>
				{{end}}
			{{end}}
		</ul>	
</body>
	</html>
	`
	case Budget:
		body = `
	<html>
	<body>
		<h1>Sending HTML emails with Mailgun</h1>
		<p style="color:blue; font-size:30px;">Hello world</p>
		<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		<ul> 
 			{{range .Data}}
				<li>{{.Name}}</li>
				{{end}}
			{{end}}
		</ul>	
</body>
	</html>
	`
	case Others:
		body = `
	<html>
	<body>
		<h1>Sending HTML emails with Mailgun</h1>
		<p style="color:blue; font-size:30px;">Hello world</p>
		<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		<ul> 
 			{{range .Data}}
				<li>{{.Name}}</li>
				{{end}}
			{{end}}
		</ul>	
</body>
	</html>
	`
	}
	return body
}
