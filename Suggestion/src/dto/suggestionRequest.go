package dto

type SuggestioneCreate struct {
	Suggestion string `json:"suggestion"`
	NoteID     uint64 `json:"noteId"`
}