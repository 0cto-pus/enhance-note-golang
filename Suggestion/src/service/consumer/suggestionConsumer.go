package consumer

import (
	"encoding/json"
	"enhance-notes-suggestion/src/dto"
	"enhance-notes-suggestion/src/helper"
	"enhance-notes-suggestion/src/service"
	"log"

	"github.com/streadway/amqp"
)


func StartConsuming(service service.ISuggestionService) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"note_queue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return
	}


		for d := range msgs {
			var noteMsg dto.ConsumerNoteMessage
			if err := json.Unmarshal(d.Body, &noteMsg); err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			suggestion, err := helper.GetEnhancedNote(noteMsg.Content)
			if err != nil {
				log.Printf("Error getting suggestion from GPT: %v", err)
				continue
			}
 		
			createdSuggestion, err := service.CreateSuggestion(dto.SuggestioneCreate{
				UserID:     noteMsg.UserID,
				NoteID:     noteMsg.NoteID,
				Suggestion: suggestion,
			})
			if err != nil {
				log.Printf("Error saving suggestion to DB: %v", err)
			} else {
				log.Printf("Suggestion saved for note %v", createdSuggestion)
			}
		}


	log.Println("Consumer started and waiting for messages.")
}