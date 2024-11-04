package repository

import (
	"enhance-notes-suggestion/src/domain"
	"errors"

	"gorm.io/gorm"
)

type ISuggestionRepository interface {
	CreateSuggestion(suggestion domain.Suggestion) (domain.Suggestion, error)
	FindSuggestionById(noteId uint64) (domain.Suggestion, error)
	GetAllSuggestionsByUserId(userId uint64) ([]domain.Suggestion,error)
}

type SuggestionRepository struct{
	db *gorm.DB
}

func NewSuggestionRepository(db *gorm.DB) ISuggestionRepository{
	return &SuggestionRepository{
		db:db,
	}
}

func (SuggestionRepository *SuggestionRepository) CreateSuggestion(suggestion domain.Suggestion) (domain.Suggestion, error) {
	if err := SuggestionRepository.db.Create(&suggestion).Error; err != nil {
		return domain.Suggestion{}, err
	}
	return suggestion, nil
}

func (SuggestionRepository *SuggestionRepository) FindSuggestionById(suggestionId uint64) (domain.Suggestion, error) {
	var foundSuggestion domain.Suggestion
	if err := SuggestionRepository.db.Where("id = ?", suggestionId).First(&foundSuggestion).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Suggestion{}, errors.New("suggestion not found")
		}
		return domain.Suggestion{}, err
	}
	return foundSuggestion, nil
}

func (SuggestionRepository *SuggestionRepository) GetAllSuggestionsByUserId(userId uint64) ([]domain.Suggestion, error) {
	var suggestions []domain.Suggestion
	if err := SuggestionRepository.db.Where("user_id = ?", userId).Find(&suggestions).Error; err != nil {
		return nil, err
	}
	return suggestions, nil
}

func (SuggestionRepository *SuggestionRepository) FindSelectedSuggestions(suggestionIds []uint64) ([]domain.Suggestion, error) {
	var suggestions []domain.Suggestion

	if err := SuggestionRepository.db.Where("id IN ?", suggestionIds).Find(&suggestions).Error; err != nil {
		return nil, errors.New("unable to find suggestions")
	}

	return suggestions, nil
}
