package main

import (
	"log"
	"log/slog"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("amqp.Dial: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("conn.Channel: %v", err)
	}

	msgs, err := ch.Consume("events", "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("ch.Consume: %v", err)
	}

	seen := make(map[string]bool)
	for message := range msgs {
		if seen[message.MessageId] {
			slog.Info("Duplicate detected, skipping message",
				slog.String("id", message.MessageId),
			)
			message.Ack(false)
			continue
		}

		seen[message.MessageId] = true

		slog.Info("Processing message",
			slog.String("id", message.MessageId),
			slog.String("body", string(message.Body)),
		)

		message.Ack(true)
	}
}
