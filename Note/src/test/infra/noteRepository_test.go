package infra

import (
	"context"
	"enhance-notes-note-service/src/domain"
	"enhance-notes-note-service/src/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var noteRepository repository.INoteRepository
var db *gorm.DB
var ctx context.Context

func TestMain(m *testing.M){
	
 

	var dsn string = "host=127.0.0.1 user=postgres password=root dbname=postgres port=6433 sslmode=disable"
	gormOpen, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	db = gormOpen;
	if err != nil {
		panic( err)
	}

	err = db.AutoMigrate(&domain.Note{})

	if err != nil {
		panic( err)
	}
	noteRepository = repository.NewNoteRepository(db)
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

	newNote := domain.Note{
		ID: 1,
		UserID: 1,
		Content: "Deneme Note",
	 }
	 t.Run("CreateNote", func(t *testing.T) {
		noteRepository.CreateNote(newNote)
		foundNote, _ := noteRepository.FindNoteById(newNote.ID)
		
		assert.Equal(t, domain.Note{
			ID: 1,
			UserID: 1,
			Content: "Deneme Note",
			CreatedAt: foundNote.CreatedAt,
			UpdatedAt: foundNote.UpdatedAt,
		 }, foundNote)
	})
	clear(ctx, db)
}

 
  func TestGetAllNotesByUserId(t *testing.T){
	notes := []domain.Note{
		{UserID: 1, Content: "Deneme Note1"},
		{UserID: 1, Content: "Deneme Note1"},
		{UserID: 2, Content: "Deneme Note1"},
	}

	expectedNotesByUserIdOne := notes[:2] // Kullanıcı ID 1 için beklenen notlar
	expectedNotesByUserIdTwo := notes[2:]  // Kullanıcı ID 2 için beklenen notlar

	 t.Run("GetAllNotesByUserId", func(t *testing.T) {

		for _, note := range notes {
			_, err := noteRepository.CreateNote(note)
			if err != nil {
				t.Fatalf("Failed to create note: %v", err)
			}
		}

		foundNotesByUserIdOne, err1 := noteRepository.GetAllNotesByUserId(1)
		foundNotesByUserIdTwo, err2 := noteRepository.GetAllNotesByUserId(2)

		assert.NoError(t, err1, "Error should be nil for user ID 1")
		assert.NoError(t, err2, "Error should be nil for user ID 2")

		assert.Equal(t, expectedNotesByUserIdOne[0].UserID, foundNotesByUserIdOne[0].UserID)
		assert.Equal(t, expectedNotesByUserIdOne[0].Content, foundNotesByUserIdOne[0].Content)

		assert.Equal(t, expectedNotesByUserIdTwo[0].UserID, foundNotesByUserIdTwo[0].UserID)
		assert.Equal(t, expectedNotesByUserIdTwo[0].Content, foundNotesByUserIdTwo[0].Content)
	})
	clear(ctx, db)
}  

