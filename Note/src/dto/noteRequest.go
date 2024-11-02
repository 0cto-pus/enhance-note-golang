package dto

type NoteCreate struct {
	UserId  string `json:"userId"`
	Content string `json:"Content"`
}