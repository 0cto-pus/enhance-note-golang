package dto

type SuggestioneCreate struct {
	UserID     uint64 `json:"user_id"`
	NoteID     uint64 `json:"note_id"`
	Suggestion string `json:"suggestion"`
}

type ConsumerNoteMessage struct {
	UserID  uint64 `json:"user_id"`
	NoteID  uint64 `json:"note_id"`
	Content string `json:"content"`
}