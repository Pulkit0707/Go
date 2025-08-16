package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main(){
	conn,err := amqp.Dial("amqp://username:password@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMq")
	defer conn.Close()
	ch,err:=conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()
	q,err:=ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to decalre a queue")
	ctx,cancel:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "Hello World!"
	err=ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)
	failOnError(err, "failed to publish a message")
	log.Printf("[x] Sent %s\n", body)
}