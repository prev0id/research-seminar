package main

import (
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Event struct {
	ID   string
	Body string
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("amqp.Dial: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("conn.Channel: %v", err)
	}

	queue, err := ch.QueueDeclare("events", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("ch.QueueDeclare: %v", err)
	}

	dupID := uuid.NewString()
	events := []Event{
		{
			ID:   dupID,
			Body: "event1",
		},
		{
			ID:   uuid.NewString(),
			Body: "event2",
		},
		{
			ID:   dupID,
			Body: "event1",
		},
	}

	for _, event := range events {
		err := ch.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event.Body),
			MessageId:   event.ID,
		})

		if err != nil {
			slog.Error("Publish error", slog.String("id", event.ID), slog.String("error", err.Error()))
		} else {
			slog.Info("Published", slog.String("id", event.ID), slog.String("body", event.Body))
		}
	}
}
