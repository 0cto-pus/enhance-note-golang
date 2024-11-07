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
    CreateSuggestion(suggestion dto.SuggestioneCreate)(domain.Suggestion, error)
    GetUserSuggestions(userId uint64) ([]domain.Suggestion,error)
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


func(suggestionService *SuggestionService) CreateSuggestion(userInput dto.SuggestioneCreate)(domain.Suggestion, error){

  suggestion := domain.Suggestion{
        UserID: userInput.UserID,
        NoteID: userInput.NoteID,
        Suggestion: userInput.Suggestion,
    }


    createdSuggestion, err := suggestionService.suggestionRepository.CreateSuggestion(suggestion)
    if err != nil {
            return domain.Suggestion{}, errors.New("failed to create note")
    }

    return createdSuggestion, nil
}

func(suggestionService *SuggestionService) GetUserSuggestions(userId uint64) ([]domain.Suggestion,error){
    suggestions, err := suggestionService.suggestionRepository.GetAllSuggestionsByUserId(userId)
    if err != nil {
        return nil, errors.New("failed to retrieve user notes")
    }

    return suggestions, nil
}  