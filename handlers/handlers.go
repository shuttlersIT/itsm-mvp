package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// CreateStatus function Usage
func CreateStatusHandler(c *gin.Context) {
	// Get the status name from the request JSON or other sources
	statusName := c.PostForm("statusName")

	// Call the CreateStatus function
	createdStatus, err := CreateStatus(c, statusName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created status
	c.JSON(http.StatusOK, gin.H{
		"message":       "Status created successfully",
		"createdStatus": createdStatus,
	})
}
