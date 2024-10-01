package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func RabbitMQApp() *amqp.Connection {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	if err != nil {
		defer connection.Close()
		panic(err)
	}
	fmt.Println("Successfully connected to RabbitMQ instance")

	return connection
}

func Channel() *amqp.Channel {
	connection := RabbitMQApp()
	channel, err := connection.Channel()
	if err != nil {
		defer channel.Close()
		panic(err)
	}
	return channel
}

func PublishQueue(deviceId string, message string) {
	// rabbitmq queue channel
	channel := Channel()

	err := channel.ExchangeDeclare(
		"notification.done", // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		log.Panicf("Failed to declare an exchange %v", err)
	}

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := struct {
		Identifier string `json:"identifier"`
		Type       string `json:"type"`
		DeviceID   string `json:"deviceId"`
		Text       string `json:"text"`
	}{
		Identifier: "fcm-msg-a1beff5ac",
		Type:       "device",
		DeviceID:   deviceId,
		Text:       message,
	}

	// publishing a message
	err = channel.Publish(
		"notification.done", // exchange
		"notification.fcm",  // key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(fmt.Sprintf("%v", payload)),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully published message")
}

type Payload struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	DeviceID   string `json:"deviceId"`
	Text       string `json:"text"`
}

func ReceiveQueue() {

	channel := Channel()

	err := channel.ExchangeDeclare(
		"notification.done", // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		log.Panicf("Failed to declare exchange %v", err)
	}

	q, err := channel.QueueDeclare(
		"notification.fcm", // name
		false,              // durable
		false,              // delete when unused
		true,               // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		log.Panicf("Failed to declare queue %v", err)
	}

	err = channel.QueueBind(
		q.Name,              // queue name
		"info",              // routing key
		"notification.done", // exchange
		false,
		nil)

	if err != nil {
		log.Panicf("Failed to bind queue %v", err)
	}

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Panicf("Failed to consume queue %v", err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)

			// IF MESSAGE RECEIVED THEN DO SOMETHING ?

			// data := Payload{}
			// err := json.Unmarshal([]byte(msg.Body), &data)
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }

			// var response string
			// if response, err = fb.SendFirebaseMessage(data.Text); err != nil {
			// 	return
			// }

			// // if response from FCM not empty then we immediately save to database
			// if response != "" {
			// 	// module.db.fcmJobModel.InsertFcmJob()
			// }
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
