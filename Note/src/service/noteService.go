package service

import (
	"enhance-notes-note-service/config"
	"enhance-notes-note-service/src/domain"
	"enhance-notes-note-service/src/dto"
	"enhance-notes-note-service/src/helper"
	"enhance-notes-note-service/src/repository"
	"enhance-notes-note-service/src/service/publisher"
	"errors"
	"log"
)

type INoteService interface {
    CreateNote(note dto.NoteCreate, userId uint64)(domain.Note, error)
    GetUserNotes(userId uint64) ([]domain.Note,error)
}

type NoteService struct {
	noteRepository repository.INoteRepository
    Auth helper.Auth
    Config config.AppConfig
}



func NewNoteService(noteRepository repository.INoteRepository, auth helper.Auth, config config.AppConfig ) INoteService{
    return &NoteService{
        noteRepository: noteRepository,
        Auth: auth,
        Config: config,
    }
}


func(noteService *NoteService) CreateNote(userInput dto.NoteCreate, userId uint64)(domain.Note, error){

  note := domain.Note{
        UserID:  userId,
        Content: userInput.Content,
    }


    createdNote, err := noteService.noteRepository.CreateNote(note)
    if err != nil {
            return domain.Note{}, errors.New("failed to create note")
    }

    err = publisher.PublishNoteMessage(createdNote.ID, userId, note.Content)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}

    return createdNote, nil
}

func(noteService *NoteService) GetUserNotes(userId uint64) ([]domain.Note,error){
    notes, err := noteService.noteRepository.GetAllNotesByUserId(userId)
    if err != nil {
        return nil, errors.New("failed to retrieve user notes")
    }

    return notes, nil
} 