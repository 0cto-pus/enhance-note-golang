package service

/* import (
	"enhance-notes-note-service/src/domain"
	"enhance-notes-note-service/src/repository"
	"fmt"
)


type MockUserRepository struct{
	notes  []domain.Note
}

func NewMockNoteRepository(mockNotes []domain.Note) repository.INoteRepository{
	return &MockUserRepository{
		notes: mockNotes,
	}
}

func (noteRepository *MockUserRepository) CreateNote(note domain.Note) (domain.Note, error) {
	noteRepository.notes = append(noteRepository.notes, domain.Note{
		ID: uint64(len(noteRepository.notes))+1,
		UserID: note.UserID,
		Content: note.Content,
	})
	return domain.Note{ID: uint64(len(noteRepository.notes))+1,
		UserID: note.UserID,
		Content: note.Content,}, nil
}

func (noteRepository *MockUserRepository) FindNoteById(noteId uint64) (domain.Note, error) {
	for _, note := range noteRepository.notes {
		if note.ID == noteId{
			return note,nil
		}
	}

	return domain.Note{}, fmt.Errorf("note with ID %v not found", noteId)
}

func (noteRepository *MockUserRepository) GetAllNotesByUserId(userId uint64) ([]domain.Note, error) {
	var foundNotes []domain.Note
	for _, note := range noteRepository.notes{
		if note.UserID == userId{
			foundNotes = append(foundNotes, note)
		}
	}

	return foundNotes,nil
}

func (noteRepository *MockUserRepository) FindSelectedNotes(noteIds []uint64) ([]domain.Note, error) {
	var foundNotes []domain.Note
	for i , note := range noteRepository.notes{
		if note.ID == noteIds[i]{
			foundNotes = append(foundNotes, note)
		}
	}
	return foundNotes,nil
}
*/