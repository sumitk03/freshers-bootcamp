package repository

import (
	"errors"
	"fmt"

	"example.com/freshers-bootcamp/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(model.Note) (uint64, error)
	Update(model.Note) error
	DeleteAll() error
	DeleteOne(uint64) error
	GetAll() ([]model.Note, error)
	GetOne(uint64) (model.Note, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewConnection() NoteRepository {
	db, err := gorm.Open(sqlite.Open("usernote.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Note{})

	return &database{
		connection: db,
	}
}

func (db *database) Create(note model.Note) (uint64, error) {
	result := db.connection.Create(&note)
	if result.Error != nil {
		return 0, result.Error
	}
	return note.ID, nil
}

func (db *database) Update(newNote model.Note) error {
	result := db.connection.Model(&model.Note{}).Where("id = ?", newNote.ID).Updates(newNote)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not exist")
	}
	return nil
}

func (db *database) DeleteAll() error {
	if result := db.connection.Where("id Like ?", "_%").Delete(&model.Note{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *database) DeleteOne(id uint64) error {
	result := db.connection.Delete(&model.Note{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not exist")
	}
	return nil
}

func (db *database) GetAll() ([]model.Note, error) {
	note := []model.Note{}
	if result := db.connection.Find(&note); result.Error != nil {
		return []model.Note{}, result.Error
	}
	return note, nil
}

func (db *database) GetOne(id uint64) (model.Note, error) {
	fmt.Println(id)
	note := model.Note{}
	if result := db.connection.Where("id = ?", id).First(&note); result.Error != nil {
		fmt.Println(result.Error)
		return model.Note{}, result.Error
	}
	return note, nil
}

func (db *database) CloseDB() {
	if v, err := db.connection.DB(); err != nil {
		panic(err)
	} else {
		v.Close()
	}
}
