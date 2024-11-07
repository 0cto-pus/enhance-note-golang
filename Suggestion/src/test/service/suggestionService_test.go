package service

import (
	"enhance-notes-suggestion/config"
	"enhance-notes-suggestion/src/domain"
	"enhance-notes-suggestion/src/dto"
	"enhance-notes-suggestion/src/helper"
	"enhance-notes-suggestion/src/service"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var suggestionService service.ISuggestionService
var auth helper.Auth

func TestMain(m *testing.M){
	 initalSuggestions := []domain.Suggestion{
		{
		ID: 1,
		UserID: 1,
		NoteID: 1,
		Suggestion: "GPT suggestion",
		},
		{
		ID: 2,
		UserID: 1,
		NoteID: 2,
		Suggestion: "GPT suggestion",
		},
		{
		ID: 3,
		UserID: 2,
		NoteID: 3,
		Suggestion: "GPT suggestion",
		},
	}

	mockRepository := NewMockSuggestionRepository(initalSuggestions)
	auth = helper.SetupAuth("enhance-notes-2024") // will change later with env
	cfg , _:= config.SetupEnv()
	suggestionService =service.NewSuggestionService(mockRepository,auth,cfg)
	exitCode:=m.Run()
	os.Exit(exitCode)
}


func Test_ShouldCreateNote(t *testing.T){
	t.Run("ShouldCreateSuggestion", func(t *testing.T) {
		suggestion, err:= suggestionService.CreateSuggestion(dto.SuggestioneCreate{Suggestion: "GPT tidy" , NoteID: 4, UserID: 1})
		assert.NoError(t, err)
		assert.Equal(t, "GPT tidy", suggestion.Suggestion)
		assert.Equal(t, uint64(1), suggestion.UserID)
	})
}

func Test_ShouldGetUserSuggestions(t *testing.T){
	t.Run("ShouldGetUserSuggestions", func(t *testing.T) {
		suggestions, err:= suggestionService.GetUserSuggestions( 1)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(suggestions))
	})
} 