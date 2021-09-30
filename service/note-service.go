package service

import (
	"example.com/freshers-bootcamp/model"
	"example.com/freshers-bootcamp/repository"
)

type NoteService interface {
	CreateNote(model.Note) (uint64, error)
	GetAllNotes() ([]model.Note, error)
	GetSingleNote(uint64) (model.Note, error)
	UpdateSingleNote(model.Note) error
	DeleteSingleNote(uint64) error
	DeleteAllNotes() error
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func New(repository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: repository,
	}
}

// func (service *noteService) isNotePresent(id uint64) bool {
// 	for _, val := range service.noteList {
// 		if val.NoteName == name {
// 			return true
// 		}
// 	}
// 	return false
// }
func (service *noteService) CreateNote(note model.Note) (uint64, error) {

	if id, err := service.noteRepository.Create(note); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func (service *noteService) GetAllNotes() ([]model.Note, error) {
	if notes, err := service.noteRepository.GetAll(); err != nil {
		return []model.Note{}, err
	} else {
		return notes, nil
	}

}

func (service *noteService) GetSingleNote(id uint64) (model.Note, error) {
	if notes, err := service.noteRepository.GetOne(id); err != nil {
		return model.Note{}, err
	} else {
		return notes, nil
	}
}

func (service *noteService) UpdateSingleNote(newNote model.Note) error {
	if err := service.noteRepository.Update(newNote); err != nil {
		return err
	}
	return nil
}

func (service *noteService) DeleteSingleNote(id uint64) error {
	if err := service.noteRepository.DeleteOne(id); err != nil {
		return err
	}
	return nil
}

func (service *noteService) DeleteAllNotes() error {
	if err := service.noteRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}
