package fb

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/kalleriakronos24/imbee-backend/rabbitmq"
	"google.golang.org/api/option"
)

func FirebaseApp() (fb *firebase.App, err error) {
	opt := option.WithCredentialsFile("./imbee-backend-service-account.json")
	fb, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase: %v", err)
	}
	log.Print("Success initializing Firebase Connection")
	return fb, nil
}

func SendFirebaseMessage(message string) (res string, err error) {

	ctx := context.Background()
	fb, _ := FirebaseApp()

	client, err := fb.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	cl, err := fb.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := cl.CustomToken(ctx, "some-uid")
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := token // "YOUR_DEVICE_TOKEN_HERE not from CustomToken func"

	// See documentation on defining a message payload.
	msg := &messaging.Message{
		Data: map[string]string{
			"message": message,
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}

	// send also to the rbmq topic
	rabbitmq.PublishQueue(registrationToken, message)
	fmt.Println("Successfully sent message:", response)

	return response, err
}
