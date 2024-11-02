package repository

import (
	"enhance-notes-note-service/src/domain"
	"errors"

	"gorm.io/gorm"
)

type INoteRepository interface {
	CreateNote(note domain.Note) (domain.Note, error)
	FindNoteById(noteId uint64) (domain.Note, error)
	GetAllNotesByUserId(userId uint64) ([]domain.Note,error)
	FindSelectedNotes(noteIds []uint64) ([]domain.Note, error)
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

func (noteRepository *NoteRepository) FindNoteById(noteId uint64) (domain.Note, error) {
	var foundNote domain.Note
	if err := noteRepository.db.Where("id = ?", noteId).First(&foundNote).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Note{}, errors.New("note not found")
		}
		return domain.Note{}, err
	}
	return foundNote, nil
}

func (noteRepository *NoteRepository) GetAllNotesByUserId(userId uint64) ([]domain.Note, error) {
	var notes []domain.Note
	if err := noteRepository.db.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (noteRepository *NoteRepository) FindSelectedNotes(noteIds []uint64) ([]domain.Note, error) {
	var notes []domain.Note

	if err := noteRepository.db.Where("id IN ?", noteIds).Find(&notes).Error; err != nil {
		return nil, errors.New("unable to find notes")
	}

	return notes, nil
}
