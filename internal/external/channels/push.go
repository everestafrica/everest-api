package channels

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"log"
)

var firebaseApp *firebase.App

func Init() {

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	setApp(app)
}

func setApp(app *firebase.App) {
	firebaseApp = app
}

func SendPushNotification(data map[string]string, firebaseToken string, notification *messaging.Notification) {

	ctx := context.Background()
	client, err := firebaseApp.Messaging(ctx)

	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
	}
	androidConfig := &messaging.AndroidConfig{Priority: string(rune(messaging.PriorityDefault))}
	appleConfig := &messaging.APNSConfig{}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data:         data,
		Token:        firebaseToken,
		Android:      androidConfig,
		APNS:         appleConfig,
		Notification: notification,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)

	if err != nil {
		log.Println(err)
	}

	log.Printf("firebase message sent: %v", response)
}
