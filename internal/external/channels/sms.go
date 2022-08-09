package channels

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	messagebird "github.com/messagebird/go-rest-api/v9"
	"github.com/messagebird/go-rest-api/v9/balance"
)

// Access keys can be managed through our dashboard.
var accessKey = config.GetConf().TestSmsApiKey

func SendSMS() {

	// Create a client.
	client := messagebird.New(accessKey)
	//sms.Create(client, "",)
	// Request the balance information, returned as a balance.Balance object.
	balances, err := balance.Read(client)
	if err != nil {
		// Handle error.
		return
	}

	// Display the results.
	fmt.Println("Payment: ", balances.Payment)
	fmt.Println("Type:", balances.Type)
	fmt.Println("Amount:", balances.Amount)

}
