package repository

import (
	"enhance-notes-note-service/src/domain"
	"errors"

	"gorm.io/gorm"
)

type INoteRepository interface {
	CreateNote(note domain.Note) (domain.Note, error)
	FindNoteById(userId uint64) (domain.Note, error)
	GetAllNotesByUserId(userId uint64) ([]domain.Note,error)
}

type NoteRepository struct{
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) INoteRepository{
	return &NoteRepository{
		db:db,
	}
}

func (noteRepository *NoteRepository) CreateNote(note domain.Note) (domain.Note, error) {
	if err := noteRepository.db.Create(&note).Error; err != nil {
		return domain.Note{}, err
	}
	return note, nil
}

func (noteRepository *NoteRepository) FindNoteById(userId uint64) (domain.Note, error) {
	var foundNote domain.Note
	if err := noteRepository.db.First(&foundNote, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Note{}, errors.New("note not found")
		}
		return domain.Note{}, err
	}
	return foundNote, nil
}

// GetAllNotesByUserId belirtilen kullanıcı ID'sine göre tüm notları döndürür
func (noteRepository *NoteRepository) GetAllNotesByUserId(userId uint64) ([]domain.Note, error) {
	var notes []domain.Note
	if err := noteRepository.db.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}