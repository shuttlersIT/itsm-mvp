package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/scanners"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get an agent id from database
func GetAgentHandler2(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	adminSession := sessions.Default(c)
	email := adminSession.Get("user-email")
	var a structs.Agent
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, usernam_id, role_id, unit, supervisor_id FROM agents WHERE email = ?", email).
		Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	adminSession.Set("id", a.AgentID)
	agent := a
	c.JSON(http.StatusOK, agent)
}

// Get an agent from database
func GetAgent(c *gin.Context, agentID int) (*structs.Agent, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("db unreacheable")
	}

	rows, err := db.Query("SELECT * FROM agents WHERE id = ?", agentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return nil, err
	}
	for rows.Next() {
		agent, e := scanners.ScanIntoAgent(rows)
		if e != nil {
			c.JSON(http.StatusNotFound, "an error occured when getting agent from db")
			return nil, e
		} else {
			c.JSON(http.StatusOK, "agent retrieval successfull")
			return agent, nil
		}
	}
	return nil, fmt.Errorf("agent ID %d not found", agentID)
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
	_, err := db.Exec("UPDATE agents SET first_name = ?, last_name = ?, agent_email = ?, username_id = ?, role_id = ?, unit = ?, supervisor_id = ?, WHERE id = ?", t.FirstName, t.LastName, t.AgentEmail, t.Username, t.RoleID, t.Unit, t.SupervisorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Agent updated successfully")
}

// Delete an agent by ID
// Delete a ticket by ID
func DeleteAgentOperation(c *gin.Context, tid int) (bool, error) {
	id := tid

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get agent handler"})
		return false, fmt.Errorf("unable to reach DB")
	}
	_, err := db.Exec("DELETE FROM agents WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false, err
	}
	return true, nil
}

// List all Agents
func ListAgents(c *gin.Context) ([]*structs.Agent, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	rows, err := db.Query("SELECT id, first_name, last_name, agent_email, username_id, role_id, unit, supervisor_id FROM agents")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("unable to retrieve agents at query")
	}
	defer rows.Close()

	var agents []*structs.Agent
	for rows.Next() {
		//var a structs.Agent
		a, err := scanners.ScanIntoAgent(rows)
		//if err := rows.Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID); err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return nil, fmt.Errorf("unable to retrieve agents")
		//}
		if err != nil {
			return nil, fmt.Errorf("unable to add agent to array")
		}
		agents = append(agents, a)
	}
	c.JSON(http.StatusOK, "Agents Listed successfully")
	return agents, nil
}

// Create staff
func CreateAgent(c *gin.Context, agent structs.Agent) (*structs.Agent, int, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, 0, fmt.Errorf("db unreacheable")
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("invalid request")
	}

	result, err := db.Exec("INSERT INTO agents (first_name, last_name, staff_email, phone, username_id, role_id, unit_id, supervisor_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", agent.FirstName, agent.LastName, agent.AgentEmail, agent.Phone, agent.Username, agent.RoleID, agent.Unit, agent.SupervisorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("failed to create Agent")
	}

	lastInsertID, _ := result.LastInsertId()
	ag, e := GetAgent(c, int(lastInsertID))
	if e != nil {
		c.JSON(http.StatusNotFound, "Agent creation failed")
		return nil, 0, e
	}

	c.JSON(http.StatusCreated, ag)
	c.JSON(http.StatusOK, "Agent created successfully")
	return ag, ag.AgentID, nil

	/*
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("Add Staff failed not found")
		}

		//lastInsertID, _ := result.LastInsertId()
		s = int(lastInsertID)
		c.JSON(http.StatusCreated, s)

		c.JSON(http.StatusOK, "User created successfully")

		return s.StaffID
	*/
}

// Update a update by Struct
func UpdateAgent(c *gin.Context, a structs.Agent) (int, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0, fmt.Errorf("unable to reach DB")
	}

	// session := sessions.Default(c)
	// id := session.Get("id")
	// var s structs.Staff
	id := a.AgentID
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, fmt.Errorf("unable to bind json")
	}
	_, err := db.Exec("UPDATE agents SET first_name = ?, last_name = ?, agent_email = ?, phone = ?, username_id = ?, role_id = ?, unit_id = ?, supervisor_id = ? WHERE id = ?", &a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.Phone, &a.RoleID, &a.Unit, &a.SupervisorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, fmt.Errorf("user not found")
	}

	if err != nil {
		c.JSON(http.StatusOK, "Agent update failed")
		return 0, fmt.Errorf("update failed")
	}
	c.JSON(http.StatusOK, "User updated successfully")
	return id, fmt.Errorf("successful")
}
