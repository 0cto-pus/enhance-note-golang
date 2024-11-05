package dto

type SuggestioneCreate struct {
	NoteId     string `json:"noteId"`
	UserId     string `json:"userId"`
	Suggestion string `json:"suggestion"`
}