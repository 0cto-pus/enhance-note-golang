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
		UserID: 1,
		Content: "Deneme Note",
	 }

	 t.Run("CreateNote", func(t *testing.T) {
		noteRepository.CreateNote(newNote)
		addedUser, _ := noteRepository.FindNoteByUserId(1)
		
		assert.Equal(t, domain.Note{
			ID: 1,
			UserID: 1,
			Content: "Deneme Note",
			CreatedAt: addedUser.CreatedAt,
			UpdatedAt: addedUser.UpdatedAt,}, addedUser)
	})
	clear(ctx, db)
}