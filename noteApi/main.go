package main

import (
	. "./model"
	"flag"
	"./utils"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"time"
)

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB

var (
	logpath = flag.String("logpath", "./logger/gin_" + time.Now().Format("2006-01-02_15-04-05") + ".log", "Log Path")
)

func init() {
	//tao ket noi database
	var err error
	db, err = gorm.Open("mysql", "root:Doxuanlinh1994@/noteDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("fail")
	}

	//Migrate database
	db.AutoMigrate(&NoteModel{})
}

func main() {
	router := gin.Default()

	// tao log
	flag.Parse()
	utils.NewLog(*logpath)

	// tao validate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidateUser", ValidateHeaderUser)
		v.RegisterValidation("ValidatePassword", ValidateHeaderPassword)
	}

	v1 := router.Group("/api/v1/note")
	{
		v1.POST("/", createNote)
		v1.GET("/", getAllNote)
		v1.GET("/:id", getSingleNote)
		v1.PUT("/:id", updateNote)
		v1.DELETE("/:id", deleteNote)
	}
	router.Run()
}

func createNote(c *gin.Context) {
	header :=Header{User:c.GetHeader("user"),Password: c.GetHeader("password")}
	if err := c.ShouldBindWith(&header, binding.Query); err == nil {
		baseNote := BaseNote{Title: c.PostForm("title")}
		note := NoteModel{BaseNote: baseNote, Content: c.PostForm("content")}
		db.Save(&note)
		utils.Log.Println("Note created successfully! resourceId:", note.ID)
		c.SecureJSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Note created successfully!", "resourceId": note.ID})
	} else {
		utils.Log.Println(err)
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "authorization are invalid!"})
	}
}

func getAllNote(c *gin.Context) {
	header :=Header{User:c.GetHeader("user"),Password: c.GetHeader("password")}
	if err := c.ShouldBindWith(&header, binding.Query); err == nil {
		var notes []NoteModel
		var _notes []TransformedNote
		db.Find(&notes)
		if len(notes) <= 0 {
			c.SecureJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Note found!"})
			return
		}

		for _, item := range notes {
			_notes = append(_notes, TransformedNote{ID: item.ID, Title: item.Title, Content: item.Content})
		}

		c.SecureJSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _notes})
	} else {
		utils.Log.Println(err)
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "authorization are invalid!"})
	}
}

func getSingleNote(c *gin.Context) {
	header :=Header{User:c.GetHeader("user"),Password: c.GetHeader("password")}
	if err := c.ShouldBindWith(&header, binding.Query); err == nil {
		var note NoteModel
		noteID := c.Param("id")
		db.First(&note, noteID)
		if note.ID == 0 {
			c.SecureJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No note found!"})
			return
		}
		_note := TransformedNote{ID: note.ID, Title: note.Title, Content: note.Content}
		c.SecureJSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _note})
	} else {
		utils.Log.Println(err)
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "authorization are invalid!"})
	}
}

func updateNote(c *gin.Context) {
	header :=Header{User:c.GetHeader("user"),Password: c.GetHeader("password")}
	if err := c.ShouldBindWith(&header, binding.Query); err != nil {
		var note NoteModel
		noteID := c.Param("id")
		db.First(&note, noteID)
		if note.ID == 0 {
			c.SecureJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No noteApi found!"})
			return
		}
		db.Model(&note).Update("title", c.PostForm("title"))
		db.Model(&note).Update("content", c.PostForm("content"))
		utils.Log.Println("Note updated successfully! resourceId:", noteID)
		c.SecureJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "note updated successfully!"})
	} else {
		utils.Log.Println(err)
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "authorization are invalid!"})
	}
}

func deleteNote(c *gin.Context) {
	header :=Header{User:c.GetHeader("user"),Password: c.GetHeader("password")}
	if err := c.ShouldBindWith(&header, binding.Query); err != nil {
		var note NoteModel
		noteID := c.Param("id")
		db.First(&note, noteID)
		if note.ID == 0 {
			c.SecureJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No noteApi found!"})
			return
		}
		db.Delete(&note)
		c.SecureJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "note deleted successfully!"})
	} else {
		utils.Log.Println(err)
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "authorization are invalid!"})
	}
}