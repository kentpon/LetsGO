package routes

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kentpon/LetsGO/routes/models"
	"github.com/kentpon/LetsGO/utils"
)

var welcomeString = "Hello World"
var DB = utils.DB

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/", Root)

	NoteAPI := router.Group("/notes")
	{
		NoteAPI.GET("", GetNoteList)
		NoteAPI.POST("", AddNote)
		NoteAPI.GET("/:id", GetNote)
		NoteAPI.PATCH("/:id", ModifyNote)
		NoteAPI.DELETE("/:id", DeleteNote)
	}

	return router
}

func Root(c *gin.Context) {
	c.String(http.StatusOK, welcomeString)
}

func GetNoteList(c *gin.Context) {
	var notes []models.Note
	if err := DB.Find(&notes).Error; err != nil {
		fmt.Println("GetNoteList DB.Find error", err)
		c.String(http.StatusInternalServerError, "GetNoteList db error")
		return
	}

	c.JSON(http.StatusOK, notes)
	return
}

func AddNote(c *gin.Context) {
	noteReq := models.Note{}
	if err := c.ShouldBindJSON(&noteReq); err != nil {
		fmt.Println("AddNote bind note request fail", err)
		c.String(http.StatusBadRequest, "AddNote Invalid Request Body")
		return
	}

	if err := DB.Create(&noteReq).Error; err != nil {
		fmt.Println("AddNote DB.Create error", err)
		c.String(http.StatusInternalServerError, "AddNote db error")
		return
	}

	c.String(http.StatusCreated, "note created")
	return
}

func GetNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := DB.Where("id = ?", id).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.String(http.StatusNotFound, "note not found")
			return
		}
		fmt.Println("GetNote DB.First", err)
		c.String(http.StatusInternalServerError, "GetNote db error")
		return
	}

	c.JSON(http.StatusOK, note)
	return
}

func ModifyNote(c *gin.Context) {
	id := c.Param("id")
	noteReq := models.Note{}
	if err := c.ShouldBindJSON(&noteReq); err != nil {
		fmt.Println("ModifyNote bind note request fail", err)
		c.String(http.StatusBadRequest, "ModifyNote Invalid Request Body")
		return
	}

	if err := DB.Model(&noteReq).Where("id = ?", id).Updates(noteReq).Error; err != nil {
		fmt.Println("ModifyNote DB.Updates", err)
		c.String(http.StatusInternalServerError, "ModifyNote db error")
		return
	}

	c.String(http.StatusOK, "note modified")
	return
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(models.Note{}, "id = ?", id).Error; err != nil {
		fmt.Println("DeleteNote DB.Delete", err)
		c.String(http.StatusInternalServerError, "DeleteNote db error")
		return
	}

	c.String(http.StatusOK, "note deleted")
	return
}
