package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

// Get an agent id from database
func GetAgentHandler(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	adminSession := sessions.Default(c)
	email := adminSession.Get("user-email")
	var a structs.Agent
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, username, role_id, unit, supervisor_id FROM agents WHERE email = ?", email).
		Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	adminSession.Set("id", a.AgentID)
	agent := a
	c.JSON(http.StatusOK, agent)
}

// Update an agent by ID
func UpdateAgentHandlers(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	adminSession := sessions.Default(c)
	id := adminSession.Get("user-id")
	var t structs.Agent
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE agents SET first_name = ?, last_name = ?, agent_email = ?, username = ?, role_id = ?, unit = ?, supervisor_id = ?, WHERE id = ?", t.FirstName, t.LastName, t.AgentEmail, t.Username, t.RoleID, t.Unit, t.SupervisorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Agent updated successfully")
}

// Delete an agent by ID
func DeleteAgentHandlers(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	adminSession := sessions.Default(c)
	id := adminSession.Get("user-id")
	_, err := db.Exec("DELETE FROM agents WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Agent deleted successfully")
}

// List all Agents
func ListAgentsHandler(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, first_name, last_name, agent_email, username, role_id, unit, supervisor_id FROM agents")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var agents []structs.Agent
	for rows.Next() {
		var a structs.Agent
		if err := rows.Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		agents = append(agents, a)
	}

	c.JSON(http.StatusOK, agents)
}
