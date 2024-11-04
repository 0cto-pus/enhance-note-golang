package infra

import (
	"context"
	"enhance-notes-suggestion/src/domain"
	"enhance-notes-suggestion/src/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var suggestionRepository repository.ISuggestionRepository
var db *gorm.DB
var ctx context.Context

func TestMain(m *testing.M){



	var dsn string = "host=127.0.0.1 user=postgres password=root dbname=postgres port=6434 sslmode=disable"
	gormOpen, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	db = gormOpen;
	if err != nil {
		panic( err)
	}

	err = db.AutoMigrate(&domain.Suggestion{})

	if err != nil {
		panic( err)
	}
	suggestionRepository = repository.NewSuggestionRepository(db)
	exitCode := m.Run()
	os.Exit(exitCode)

}
func setup(ctx context.Context, db *gorm.DB) {
	TestDataInitialize(ctx, db)
}
func clear(ctx context.Context, db *gorm.DB) {
	TruncateTestData(ctx, db)
}

func TestCreateNote(t *testing.T){

	newNote := domain.Suggestion{
		ID: 1,
		UserID: 1,
		NoteID: 1,
		Suggestion: "Deneme deneme", 
	 }
	 t.Run("CreateNote", func(t *testing.T) {
		suggestionRepository.CreateSuggestion(newNote)
		foundNote, _ := suggestionRepository.FindSuggestionById(newNote.ID)

		assert.Equal(t, domain.Suggestion{
			ID: 1,
			UserID: 1,
			NoteID: 1,
			Suggestion: "Deneme deneme", 
			CreatedAt: foundNote.CreatedAt,
		 }, foundNote)
	})
	clear(ctx, db)
}


  func TestGetAllNotesByUserId(t *testing.T){


	suggestions := []domain.Suggestion{
		{ID: 1,
			UserID: 1,
			NoteID: 1,
			Suggestion: "Deneme deneme", },
		{ID: 2,
			UserID: 1,
			NoteID: 2,
			Suggestion: "Deneme deneme", },
		{ID: 3,
			UserID: 2,
			NoteID: 3,
			Suggestion: "Deneme deneme", },
	}

	expectedSuggestionsByUserIdOne := suggestions[:2] // Kullanıcı ID 1 için beklenen notlar
	expectedSuggestionsByUserIdTwo := suggestions[2:]  // Kullanıcı ID 2 için beklenen notlar

	 t.Run("GetAllNotesByUserId", func(t *testing.T) {

		for _, suggestion := range suggestions {
			_, err := suggestionRepository.CreateSuggestion(suggestion)
			if err != nil {
				t.Fatalf("Failed to create note: %v", err)
			}
		}

		foundSuggestionsByUserIdOne, err1 := suggestionRepository.GetAllSuggestionsByUserId(1)
		foundSuggestionsByUserIdTwo, err2 := suggestionRepository.GetAllSuggestionsByUserId(2)

		assert.NoError(t, err1, "Error should be nil for user ID 1")
		assert.NoError(t, err2, "Error should be nil for user ID 2")

		assert.Equal(t, expectedSuggestionsByUserIdOne[0].UserID, foundSuggestionsByUserIdOne[0].UserID)
		assert.Equal(t, expectedSuggestionsByUserIdOne[0].Suggestion, foundSuggestionsByUserIdOne[0].Suggestion)

		assert.Equal(t, expectedSuggestionsByUserIdTwo[0].UserID, foundSuggestionsByUserIdTwo[0].UserID)
		assert.Equal(t, expectedSuggestionsByUserIdTwo[0].Suggestion, foundSuggestionsByUserIdTwo[0].Suggestion)
	})
	clear(ctx, db)
}