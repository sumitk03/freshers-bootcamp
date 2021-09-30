package controller

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"example.com/freshers-bootcamp/model"
	"example.com/freshers-bootcamp/service"
	"github.com/gin-gonic/gin"
)

type NoteController interface {
	CreateNote(*gin.Context) (uint64, error)
	GetAllNotes() ([]model.Note, error)
	GetSingleNote(*gin.Context) (model.Note, error)
	DeleteAllNotes() error
	DeleteSingleNote(*gin.Context) error
	UpdateSingleNote(*gin.Context) error
}

type noteContrroller struct {
	noteService service.NoteService
}

func New(s service.NoteService) NoteController {
	return &noteContrroller{noteService: s}
}

func (controller *noteContrroller) CreateNote(ctx *gin.Context) (uint64, error) {
	var newNote model.Note
	if err := ctx.ShouldBindJSON(&newNote); err != nil {
		return 0, err
	}
	if v := len(newNote.NoteTitle); !(v > 0) {
		return 0, errors.New("note title is empty")
	}
	if v := len(newNote.NoteDetail); !(v > 0) {
		return 0, errors.New("note detail is empty")
	}
	curr_time := time.Now()
	newNote.NoteCreatedAt = curr_time
	newNote.NoteUpdatedAt = curr_time
	fmt.Println(newNote)
	if id, err := controller.noteService.CreateNote(newNote); err != nil {
		return 0, err
	} else {
		return id, nil
	}

}

func (controller *noteContrroller) GetAllNotes() ([]model.Note, error) {
	if notes, err := controller.noteService.GetAllNotes(); err != nil {
		return []model.Note{}, err
	} else {
		return notes, nil
	}
}

func (controller *noteContrroller) GetSingleNote(ctx *gin.Context) (model.Note, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return model.Note{}, err
	}
	if note, err := controller.noteService.GetSingleNote(id); err != nil {
		return model.Note{}, err
	} else {
		return note, nil
	}
}

func (controller *noteContrroller) DeleteAllNotes() error {
	if err := controller.noteService.DeleteAllNotes(); err != nil {
		return err
	}
	return nil
}

func (controller *noteContrroller) DeleteSingleNote(ctx *gin.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	if err := controller.noteService.DeleteSingleNote(id); err != nil {
		return err
	}
	return nil
}

func (controller *noteContrroller) UpdateSingleNote(ctx *gin.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	var newNote model.Note
	if err := ctx.ShouldBindJSON(&newNote); err != nil {
		return err
	}
	if v1, v2 := len(newNote.NoteTitle), len(newNote.NoteDetail); !(v1 > 0) && !(v2 > 0) {
		return errors.New("both field can not be empty")
	}
	newNote.NoteUpdatedAt = time.Now()
	newNote.ID = id
	if err := controller.noteService.UpdateSingleNote(newNote); err != nil {
		return err
	}
	return nil
}
