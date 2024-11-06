package service

import (
	"enhance-notes-suggestion/src/domain"
	"enhance-notes-suggestion/src/repository"
	"fmt"
)


type MockUserRepository struct{
	suggestions  []domain.Suggestion
}

func NewMockSuggestionRepository(mockSuggestions []domain.Suggestion) repository.ISuggestionRepository{
	return &MockUserRepository{
		suggestions: mockSuggestions,
	}
}

func (suggestionRepository *MockUserRepository) CreateSuggestion(suggestion domain.Suggestion) (domain.Suggestion, error) {
	suggestionRepository.suggestions = append(suggestionRepository.suggestions, domain.Suggestion{
		ID: uint64(len(suggestionRepository.suggestions))+1,
		UserID: suggestion.UserID,
		NoteID: suggestion.NoteID,
		Suggestion: suggestion.Suggestion,
	})
	return domain.Suggestion{
		ID: uint64(len(suggestionRepository.suggestions))+1,
		UserID: suggestion.UserID,
		NoteID: suggestion.NoteID,
		Suggestion: suggestion.Suggestion,
	}, nil
}

func (suggestionRepository *MockUserRepository) FindSuggestionById(noteId uint64) (domain.Suggestion, error) {
	
	for _, note := range suggestionRepository.suggestions {
		if note.ID == noteId{
			return note,nil
		}
	}
	
	return domain.Suggestion{}, fmt.Errorf("note with ID %v not found", noteId)
}

func (suggestionRepository *MockUserRepository) GetAllSuggestionsByUserId(userId uint64) ([]domain.Suggestion, error) {
	var foundSuggestions []domain.Suggestion
	for _, suggestion := range suggestionRepository.suggestions{
		if suggestion.UserID == userId{
			foundSuggestions = append(foundSuggestions, suggestion)
		}
	}

	return foundSuggestions,nil
}


func (suggestionRepository *MockUserRepository) FindSelectedSuggestions(suggestionIds []uint64) ([]domain.Suggestion, error) {
	var foundSuggestions []domain.Suggestion
	for i , suggestion := range suggestionRepository.suggestions{
		if suggestion.ID == suggestionIds[i]{
			foundSuggestions = append(foundSuggestions, suggestion)
		}
	}
	return foundSuggestions,nil
}
