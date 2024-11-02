package service

import (
	"enhance-notes-note-service/config"
	"enhance-notes-note-service/src/domain"
	"enhance-notes-note-service/src/dto"
	"enhance-notes-note-service/src/helper"
	"enhance-notes-note-service/src/service"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var noteService service.INoteService
var auth helper.Auth

func TestMain(m *testing.M){
	 initialUsers := []domain.Note{
		{
		ID: 1,
		UserID: 1,
		Content: "This is my note",
		},
		{
		ID: 2,
		UserID: 1,
		Content: "This is my note",
		},
		{
		ID: 3,
		UserID: 2,
		Content: "This is my note",
		},
	} 

	mockRepository := NewMockNoteRepository(initialUsers)
	auth = helper.SetupAuth("enhance-notes-2024")
	cfg , _:= config.SetupEnv()
	noteService =service.NewNoteService(mockRepository,auth,cfg)
	exitCode:=m.Run()
	os.Exit(exitCode)
}


func Test_ShouldCreateNote(t *testing.T){
	t.Run("ShouldCreateNote", func(t *testing.T) {
		note, err:= noteService.CreateNote(dto.NoteCreate{Content: "My new test note"}, 1)
		assert.NoError(t, err)
		assert.Equal(t, "My new test note", note.Content)
		assert.Equal(t, uint64(1), note.UserID)	
	})
}

func Test_ShouldGetUserNotes(t *testing.T){
	t.Run("ShouldGetUserNotes", func(t *testing.T) {
		notes, err:= noteService.GetUserNotes( 1)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(notes))
	})
}