package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get a username from database
func GetAgentCred2(c *gin.Context, username int) (*structs.AgentLoginCredentials, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to read db")
	}

	var s structs.AgentLoginCredentials
	err := db.QueryRow("SELECT id, username, password FROM agent_credentials WHERE id = ?", username).
		Scan(&s.CredentialID, s.Username, s.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return nil, fmt.Errorf("agent credentials not found")
	}
	agentCredentials := s
	c.JSON(http.StatusOK, agentCredentials)

	return &agentCredentials, nil
}
