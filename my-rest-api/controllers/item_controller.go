package controllers

import (
	"log"
	"my-rest-api/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AddBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// always exported thing should be in capital
func GetBooks(c *gin.Context){
	// Respond with a JSON array of all books
	c.JSON(http.StatusOK, database.Books)
}

func AddBooks(c *gin.Context) {
	// Bind the incoming JSON to the AddBookRequest struct
	var req AddBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Respond with 400 if the request is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the book into the database using Prisma
	book, err := database.PrismaClient.Book.CreateOne(
		database.PrismaClient.Book.title.Set(req.Title),
		database.PrismaClient.Book.author.Set(req.Author),
	).Exec(context.Background())
	if err != nil {
		log.Printf("Error adding book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	// Respond with the created book
	c.JSON(http.StatusCreated, gin.H{"book": book})
}