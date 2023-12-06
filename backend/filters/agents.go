package filters

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get a Agent by First Name
func GetAgentByFirstName(c *gin.Context, fname string) (int, structs.Agent) {
	var s structs.Agent
	first_name := fname

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, username_id, role_id, unit_id, supervisor_id FROM agents WHERE first_name = ?", first_name).
		Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get Agent by Last Name
func GetAgentByLastName(c *gin.Context, lname string) (int, structs.Agent) {
	var s structs.Agent
	last_name := lname

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, role_id, unit_id, supervisor_id FROM staff WHERE last_name = ?", last_name).
		Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get an Agent by Email
func GetAgentByEmail(c *gin.Context, e string) (int, structs.Agent) {
	var s structs.Agent
	email := e

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, role_id, unit_id, supervisor_id FROM staff WHERE email = ?", email).
		Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get an Agent by Username
func GetAgentByUsername(c *gin.Context, u int) (int, structs.Agent) {
	var s structs.Agent
	username := u

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, role_id, unit_id, supervisor_id FROM staff WHERE username = ?", username).
		Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get an Agent by Role
func GetAgentByRole(c *gin.Context, r int) (int, []structs.Agent) {
	var staff []structs.Agent
	role := r

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, staff
	}
	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username_id, role_id, unit_id, supervisor_id FROM staff WHERE position_id = ?", role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, staff
	}
	defer rows.Close()

	for rows.Next() {
		var s structs.Agent
		if err := rows.Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, staff
		}
		staff = append(staff, s)
	}

	c.JSON(http.StatusOK, staff)
	return 1, staff
}

// Get an Agent by Unit
func GetAgentByUnit(c *gin.Context, u int) (int, []structs.Agent) {
	var staff []structs.Agent
	unit := u

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, staff
	}
	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username, role_id, unit_id, supervisor_id FROM staff WHERE department_id = ?", unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, staff
	}
	defer rows.Close()

	for rows.Next() {
		var s structs.Agent
		if err := rows.Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, staff
		}
		staff = append(staff, s)
	}

	c.JSON(http.StatusOK, staff)
	return 1, staff
}

// Get an Agent by Supervisor
func GetAgentBySupervisor(c *gin.Context, supo int) (int, []structs.Agent) {
	var staff []structs.Agent
	supervisor := supo

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, staff
	}
	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username, role_id, unit_id, supervisor_id FROM staff WHERE department_id = ?", supervisor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, staff
	}
	defer rows.Close()

	for rows.Next() {
		var s structs.Agent
		if err := rows.Scan(&s.AgentID, &s.FirstName, &s.LastName, &s.AgentEmail, &s.Username, &s.RoleID, &s.Unit, &s.SupervisorID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, staff
		}
		staff = append(staff, s)
	}

	c.JSON(http.StatusOK, staff)
	return 1, staff
}
