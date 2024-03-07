package main

import (
	"book-inventory-go/app"
	"book-inventory-go/auth"
	"book-inventory-go/db"
	"book-inventory-go/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	// home
	router.GET("/", auth.HomeHandler)

	// get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)

	// detail book
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// add book
	router.GET("/add-book", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// update book
	router.GET("/update-book/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/update-book/:id", middleware.AuthValid, handler.PutBook)

	// delete book
	router.POST("/delete-book/:id", middleware.AuthValid, handler.DeleteBook)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	router.Run()
}
