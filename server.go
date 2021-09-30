package main

import (
	"net/http"
	"strconv"

	"example.com/freshers-bootcamp/controller"
	"example.com/freshers-bootcamp/middleware"
	"example.com/freshers-bootcamp/repository"
	"example.com/freshers-bootcamp/service"
	"github.com/gin-gonic/gin"
)

var (
	noteRepository  repository.NoteRepository = repository.NewConnection()
	noteService     service.NoteService       = service.New(noteRepository)
	noteContrroller controller.NoteController = controller.New(noteService)
)

func main() {
	defer noteRepository.CloseDB()
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), middleware.AuthouriseRoute())

	router.GET("/heart-beat", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	v1 := router.Group("/v1/notes")
	{
		v1.GET("", func(c *gin.Context) {
			if notes, err := noteContrroller.GetAllNotes(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to Fetch All records",
				})
			} else {
				c.JSON(http.StatusOK, notes)
			}
		})

		v1.GET("/:id", func(c *gin.Context) {
			if note, err := noteContrroller.GetSingleNote(c); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to fetch record",
				})
			} else {
				c.JSON(http.StatusOK, note)
			}
		})

		v1.POST("/create", func(c *gin.Context) {
			if id, err := noteContrroller.CreateNote(c); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Either Duplicate NoteName OR Try Again with other Note",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Note Successfully Created",
					"id":      strconv.FormatUint(id, 10),
				})
			}

		})

		v1.PATCH("/:id", func(c *gin.Context) {
			if err := noteContrroller.UpdateSingleNote(c); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Not Updated. Try Again!",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Note Successfully Updated",
				})
			}

		})

		v1.DELETE("", func(c *gin.Context) {
			if err := noteContrroller.DeleteAllNotes(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Not Deleted. Try Again!",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Successfully Deleted All notes",
				})
			}
		})

		v1.DELETE("/:id", func(c *gin.Context) {
			if err := noteContrroller.DeleteSingleNote(c); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Not Deleted. Try Again!",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Successfully Deleted the note",
				})
			}
		})
	}

	router.Run(":8080")
}
