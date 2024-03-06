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

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	router.Run()
}
