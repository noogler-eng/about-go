package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"my-rest-api/database"
)

// always exported thing should be in capital
func GetBooks(c *gin.Context){
	// Respond with a JSON array of all books
	c.JSON(http.StatusOK, database.Books)
}