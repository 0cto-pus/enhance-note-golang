package service

import (
	"enhance-notes-suggestion/config"
	"enhance-notes-suggestion/src/domain"
	"enhance-notes-suggestion/src/dto"
	"enhance-notes-suggestion/src/helper"
	"enhance-notes-suggestion/src/repository"
	"errors"
)

type ISuggestionService interface {
    CreateNote(note dto.SuggestioneCreate, userId uint64)(domain.Suggestion, error)
    GetUserNotes(userId uint64) ([]domain.Suggestion,error)
}

type SuggestionService struct {
	suggestionRepository repository.ISuggestionRepository
    Auth helper.Auth
    Config config.AppConfig 
}



func NewSuggestionService(suggestionRepository repository.ISuggestionRepository,  auth helper.Auth, config config.AppConfig ) ISuggestionService{
    return &SuggestionService{
        suggestionRepository: suggestionRepository,
         Auth: auth,
        Config: config, 
    }
}


func(SuggestionService *SuggestionService) CreateSuggestion(userInput dto.SuggestioneCreate, userId uint64, noteId uint64)(domain.Suggestion, error){

  suggestion := domain.Suggestion{
        NoteID: noteId,
        UserID:  userId,
        Suggestion: userInput.Suggestion,
    }


    createdSuggestion, err := SuggestionService.suggestionRepository.CreateSuggestion(suggestion)
    if err != nil {
            return domain.Suggestion{}, errors.New("failed to create note")
    }

    return createdSuggestion, nil
}

func(SuggestionService *SuggestionService) GetUserNotes(userId uint64) ([]domain.Suggestion,error){
    suggestions, err := SuggestionService.suggestionRepository.GetAllSuggestionsByUserId(userId)
    if err != nil {
        return nil, errors.New("failed to retrieve user notes")
    }

    return suggestions, nil
}  