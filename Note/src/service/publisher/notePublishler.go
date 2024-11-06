package publisher

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type NoteMessage struct {
	NoteID  uint64 `json:"note_id"`
	Content string `json:"content"`
}

func PublishNoteMessage(noteID uint64, content string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"note_queue", // queue name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	noteMsg := NoteMessage{
		NoteID:  noteID,
		Content: content,
	}

	body, err := json.Marshal(noteMsg)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",          // exchange
		q.Name,      // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf("Message sent: %s", body)
	return nil
}