package service

/* import (
	"enhance-notes-note-service/config"
	"enhance-notes-note-service/src/dto"
	"enhance-notes-note-service/src/helper"
	"enhance-notes-note-service/src/repository"
)

type INoteService interface {
    Login(userInput dto.UserLogin)(string, error)
    SignUp(userInput dto.UserSignUp) (string,error)
}

type NoteService struct {
	noteRepository repository.INoteRepository
    Auth helper.Auth
    Config config.AppConfig
}



func NewNoteService(noteRepository repository.INoteRepository, auth helper.Auth, config config.AppConfig ) INoteService{
    return &NoteService{
        noteRepository: noteRepository,
        Auth: auth,
        Config: config,
    }
}


func(noteService *NoteService) Login(userInput dto.UserLogin)(string, error){
  return "", nil
}

func(noteService *NoteService) SignUp(userInput dto.UserSignUp) (string, error){
 return "", nil
} */