package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/scanners"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

//////// SLA /////////////////////////////////////////////////////////

// Get SLA by ID from database
func GetSla(c *gin.Context, slaID int) (*structs.Sla, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := slaID
	var sla structs.Sla
	err := db.QueryRow("SELECT id, sla_name, priority_id, satisfaction_id, policy_id FROM sla WHERE id = ?", id).
		Scan(&sla.SlaID, &sla.SlaName, &sla.PriorityID, &sla.SatisfactionID, &sla.PolicyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SLA not found"})
		return nil, fmt.Errorf("SLA not found")
	}
	c.JSON(http.StatusOK, sla)
	return &sla, nil
}

// Update SLA by ID
func UpdateSla(c *gin.Context, slaID int, slaName string, priorityID int, satisfactionID int, policyID int) (*structs.Sla, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update SLA handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := slaID

	var sla structs.Sla
	if err := c.ShouldBindJSON(&sla); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}
	_, err := db.Exec("UPDATE sla SET sla_name = ?, priority_id = ?, satisfaction_id = ?, policy_id = ? WHERE id = ?", slaName, priorityID, satisfactionID, policyID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update SLA")
	}

	// Retrieve the updated SLA from the database
	updatedSla, err := GetSla(c, slaID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated SLA")
	}

	c.JSON(http.StatusOK, "SLA updated successfully")
	return updatedSla, nil
}

// Delete SLA by ID
func DeleteSla(c *gin.Context, slaID int) (*string, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update SLA handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := slaID
	var status string
	_, err := db.Exec("DELETE FROM sla WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "SLA deletion failed"
		return &status, err
	}
	status = "SLA deleted successfully"
	return &status, nil
}

// Create SLA
func CreateSla(c *gin.Context) (*structs.Sla, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var sla structs.Sla
	if err := c.ShouldBindJSON(&sla); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}
	result, err := db.Exec("INSERT INTO sla (sla_name, priority_id, satisfaction_id, policy_id) VALUES (?, ?, ?, ?)", sla.SlaName, sla.PriorityID, sla.SatisfactionID, sla.PolicyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create SLA")
	}

	lastInsertID, _ := result.LastInsertId()
	sla.SlaID = int(lastInsertID)
	c.JSON(http.StatusCreated, sla)

	c.JSON(http.StatusOK, "SLA created successfully")
	return &sla, nil
}

// List all SLAs
func ListSla(c *gin.Context) ([]structs.Sla, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, sla_name, priority_id, satisfaction_id, policy_id FROM sla")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch SLAs")
	}
	defer rows.Close()

	var slas []structs.Sla
	for rows.Next() {
		var sla structs.Sla
		if err := rows.Scan(&sla.SlaID, &sla.SlaName, &sla.PriorityID, &sla.SatisfactionID, &sla.PolicyID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan SLAs")
		}
		slas = append(slas, sla)
	}

	c.JSON(http.StatusOK, slas)
	return slas, nil
}

//////////////////////////////////////////////////////////////////////////

//////// PRIORITY ////////

// Get a priority by ID from the database
func GetPriority(c *gin.Context, priorityID int) (*structs.Priority, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var priority structs.Priority
	err := db.QueryRow("SELECT id, priority_name, first_response, colour FROM priority WHERE id = ?", priorityID).
		Scan(&priority.PriorityID, &priority.Name, &priority.FirstResponse, &priority.Colour)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Priority not found"})
		return nil, fmt.Errorf("priority not found")
	}
	c.JSON(http.StatusOK, priority)
	return &priority, nil
}

// Update a priority by ID
func UpdatePriority(c *gin.Context, priorityID int, priorityName string, firstResponse int, colour string) (*structs.Priority, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update priority handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := priorityID
	pName := priorityName
	fResponse := firstResponse
	col := colour

	var priority structs.Priority
	if err := c.ShouldBindJSON(&priority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}

	_, err := db.Exec("UPDATE priority SET priority_name = ?, first_response = ?, colour = ? WHERE id = ?", pName, fResponse, col, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update priority")
	}

	// Retrieve the updated priority from the database
	updatedPriority, err := GetPriority(c, priorityID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated priority")
	}

	c.JSON(http.StatusOK, "Priority updated successfully")
	return updatedPriority, nil
}

// Delete a priority by ID
func DeletePriority(c *gin.Context, priorityID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update priority handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := priorityID
	var status string
	_, err := db.Exec("DELETE FROM priority WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Priority deletion failed"
		return &status, err
	}
	status = "Priority deleted successfully"
	return &status, nil
}

// Create Priority
func CreatePriority(c *gin.Context) (*structs.Priority, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var priority structs.Priority
	if err := c.ShouldBindJSON(&priority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}
	result, err := db.Exec("INSERT INTO priority (priority_name, first_response, colour) VALUES (?, ?, ?)", priority.Name, priority.FirstResponse, priority.Colour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create priority")
	}

	lastInsertID, _ := result.LastInsertId()
	priority.PriorityID = int(lastInsertID)
	c.JSON(http.StatusCreated, priority)

	c.JSON(http.StatusOK, "Priority created successfully")
	return &priority, nil
}

// List all Priority Ranks
func ListPriorities(c *gin.Context) ([]structs.Priority, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, priority_name, first_response, colour FROM priority")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch priorities")
	}
	defer rows.Close()

	var priorities []structs.Priority
	for rows.Next() {
		var priority structs.Priority
		if err := rows.Scan(&priority.PriorityID, &priority.Name, &priority.FirstResponse, &priority.Colour); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan priorities")
		}
		priorities = append(priorities, priority)
	}

	c.JSON(http.StatusOK, priorities)
	return priorities, nil
}

//////// SATISFACTION ////////

//////////////////////////////////////////////////////////////////////////

// Get Satisfaction rank from database
func GetSatisfaction(c *gin.Context, satisfactionID int) (*structs.Satisfaction, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var satisfaction structs.Satisfaction
	err := db.QueryRow("SELECT id, satisfaction_name, rank, emoji FROM satisfaction WHERE id = ?", satisfactionID).
		Scan(&satisfaction.SatisfactionID, &satisfaction.Name, &satisfaction.Rank, &satisfaction.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satisfaction Rank not found"})
		return nil, fmt.Errorf("satisfaction rank not found")
	}
	c.JSON(http.StatusOK, satisfaction)
	return &satisfaction, nil
}

// Delete a Satisfaction Rank by ID
func DeleteSatisfaction(c *gin.Context, satisfactionID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := satisfactionID
	var status string
	_, err := db.Exec("DELETE FROM satisfaction WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Satisfaction Rank deletion failed"
		return &status, err
	}
	status = "Satisfaction Rank deleted successfully"
	return &status, nil
}

// Update Satisfaction Rank by ID
func UpdateSatisfaction(c *gin.Context, satisfactionID int, satisfactionName string, rank int, emoji string) (*structs.Satisfaction, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update satisfaction handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := satisfactionID
	sName := satisfactionName
	sRank := rank
	sEmoji := emoji

	var satisfaction structs.Satisfaction
	if err := c.ShouldBindJSON(&satisfaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}

	_, err := db.Exec("UPDATE satisfaction SET satisfaction_name = ?, rank = ?, emoji = ? WHERE id = ?", sName, sRank, sEmoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update satisfaction")
	}

	// Retrieve the updated satisfaction from the database
	updatedSatisfaction, err := GetSatisfaction(c, satisfactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated satisfaction")
	}

	c.JSON(http.StatusOK, "Satisfaction updated successfully")
	return updatedSatisfaction, nil
}

// Create Satisfaction Rank
func CreateSatisfaction(c *gin.Context, satisfactionName string, rank int, emoji string) (*structs.Satisfaction, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var satisfaction structs.Satisfaction
	if err := c.ShouldBindJSON(&satisfaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}
	result, err := db.Exec("INSERT INTO satisfaction (satisfaction_name, rank, emoji) VALUES (?, ?, ?)", satisfactionName, rank, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create satisfaction")
	}

	lastInsertID, _ := result.LastInsertId()
	satisfaction.SatisfactionID = int(lastInsertID)
	c.JSON(http.StatusCreated, satisfaction)

	c.JSON(http.StatusOK, "Satisfaction created successfully")

	return &satisfaction, nil
}

// List all Satisfaction Ranks
func ListSatisfaction(c *gin.Context) ([]structs.Satisfaction, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, satisfaction_name, rank, emoji FROM satisfaction")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch satisfaction ranks")
	}
	defer rows.Close()

	var satisfactions []structs.Satisfaction
	for rows.Next() {
		var satisfaction structs.Satisfaction
		if err := rows.Scan(&satisfaction.SatisfactionID, &satisfaction.Name, &satisfaction.Rank, &satisfaction.Emoji); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan satisfaction ranks")
		}
		satisfactions = append(satisfactions, satisfaction)
	}

	c.JSON(http.StatusOK, satisfactions)
	return satisfactions, nil
}

//////////////////////////////////////////////////////////////////////////

//////// POLICY ////////

// Get a policy from the database
func GetPolicy(c *gin.Context, policyID int) (*structs.Policies, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get policy handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := policyID
	var policy structs.Policies
	err := db.QueryRow("SELECT id, policy_name, embedded_link, policy_url FROM policy WHERE id = ?", id).
		Scan(&policy.PolicyID, &policy.PolicyName, &policy.EmbeddedLink, &policy.PolicyUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return nil, fmt.Errorf("policy not found")
	}
	c.JSON(http.StatusOK, policy)
	return &policy, nil
}

// Update Policy by ID
func UpdatePolicy(c *gin.Context, policyID int, policyName string, embeddedLink string, policyUrl string) (*structs.Policies, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from update policy handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := policyID
	var policy structs.Policies
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request data")
	}

	_, err := db.Exec("UPDATE policy SET policy_name = ?, embedded_link = ?, policy_url = ? WHERE id = ?", policyName, embeddedLink, policyUrl, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update policy")
	}

	// Retrieve the updated policy from the database
	updatedPolicy, err := GetPolicy(c, policyID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated policy")
	}

	c.JSON(http.StatusOK, "Policy updated successfully")
	return updatedPolicy, nil
}

// Delete a policy by ID
func DeletePolicy(c *gin.Context, policyID int) (*string, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from delete policy handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := policyID
	var status string
	_, err := db.Exec("DELETE FROM policy WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Policy deletion failed"
		return &status, err
	}
	status = "Policy deleted successfully"
	return &status, nil
}

// Create policy
func CreatePolicy(c *gin.Context, policyName string, embeddedLink string, policyUrl string) (*structs.Policies, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from create policy handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var policy structs.Policies
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	result, err := db.Exec("INSERT INTO policy (policy_name, embedded_link, policy_url) VALUES (?, ?, ?)", policyName, embeddedLink, policyUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create policy")
	}

	lastInsertID, _ := result.LastInsertId()
	policy.PolicyID = int(lastInsertID)
	c.JSON(http.StatusCreated, policy)

	c.JSON(http.StatusOK, "Policy created successfully")
	return &policy, nil
}

// List all policies
func ListPolicies(c *gin.Context) ([]structs.Policies, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from list policies handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, policy_name, embedded_link, policy_url FROM policy")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch policies")
	}
	defer rows.Close()

	var policies []structs.Policies
	for rows.Next() {
		var policy structs.Policies
		if err := rows.Scan(&policy.PolicyID, &policy.PolicyName, &policy.EmbeddedLink, &policy.PolicyUrl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan policies")
		}
		policies = append(policies, policy)
	}

	c.JSON(http.StatusOK, policies)
	return policies, nil
}

//////////////////////////////////////////////////////////////////////////

//////// POSITIONS ////////

// Get a position from the database
func GetPosition(c *gin.Context, positionID int) (*structs.Position, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get position handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var position structs.Position
	err := db.QueryRow("SELECT id, position_name, cadre_name FROM positions WHERE id = ?", positionID).
		Scan(&position.PositionID, &position.PositionName, &position.CadreName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return nil, fmt.Errorf("position not found")
	}
	c.JSON(http.StatusOK, position)
	return &position, nil
}

// Update a position by ID
func UpdatePosition(c *gin.Context, positionID int, positionName string, cadreName string) (*structs.Position, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from update position handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := positionID
	var position structs.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE positions SET position_name = ?, cadre_name = ? WHERE id = ?", positionName, cadreName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update position")
	}
	c.JSON(http.StatusOK, "Position updated successfully")
	return &position, nil
}

// Delete a position by ID
func DeletePosition(c *gin.Context, positionID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from delete position handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := positionID
	var status string
	_, err := db.Exec("DELETE FROM positions WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Position deletion failed"
		return &status, err
	}
	status = "Position deleted successfully"
	return &status, nil
}

// Create position
func CreatePosition(c *gin.Context, positionName string, cadreName string) (*structs.Position, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from create position handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var position structs.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request data")
	}
	result, err := db.Exec("INSERT INTO positions (position_name, cadre_name) VALUES (?, ?)", positionName, cadreName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create position")
	}

	lastInsertID, _ := result.LastInsertId()
	position.PositionID = int(lastInsertID)
	c.JSON(http.StatusCreated, position)

	c.JSON(http.StatusOK, "Position created successfully")

	return &position, nil
}

// List all positions
func ListPositions(c *gin.Context) ([]structs.Position, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from list positions handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, position_name, cadre_name FROM positions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch positions")
	}
	defer rows.Close()

	var positions []structs.Position
	for rows.Next() {
		var position structs.Position
		if err := rows.Scan(&position.PositionID, &position.PositionName, &position.CadreName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan positions")
		}
		positions = append(positions, position)
	}

	c.JSON(http.StatusOK, positions)
	return positions, nil
}

//////////////////////////////////////////////////////////////////////////

//////// DEPARTMENTS ////////

// Get a department from the database
func GetDepartment(c *gin.Context, departmentID int) (*structs.Department, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get department handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var department structs.Department
	err := db.QueryRow("SELECT id, department_name, emoji FROM departments WHERE id = ?", departmentID).
		Scan(&department.DepartmentID, &department.DepartmentName, &department.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return nil, fmt.Errorf("department not found")
	}
	c.JSON(http.StatusOK, department)
	return &department, nil
}

// Update a department by ID
func UpdateDepartment(c *gin.Context, departmentID int, departmentName string, emoji string) (*structs.Department, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from update department handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := departmentID
	var department structs.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE departments SET department_name = ?, emoji = ? WHERE id = ?", departmentName, emoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update department")
	}
	c.JSON(http.StatusOK, "Department updated successfully")
	return &department, nil
}

// Delete a department by ID
func DeleteDepartment(c *gin.Context, departmentID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from delete department handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := departmentID
	var status string
	_, err := db.Exec("DELETE FROM departments WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Department deletion failed"
		return &status, err
	}
	status = "Department deleted successfully"
	return &status, nil
}

// Create department
func CreateDepartment(c *gin.Context, departmentName string, emoji string) (*structs.Department, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from create department handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var department structs.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	result, err := db.Exec("INSERT INTO departments (department_name, emoji) VALUES (?, ?)", departmentName, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create department")
	}

	lastInsertID, _ := result.LastInsertId()
	department.DepartmentID = int(lastInsertID)
	c.JSON(http.StatusCreated, department)

	c.JSON(http.StatusOK, "Department created successfully")
	return &department, nil
}

// List all departments
func ListDepartments(c *gin.Context) ([]structs.Department, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from list departments handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, department_name, emoji FROM departments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch departments")
	}
	defer rows.Close()

	var departments []structs.Department
	for rows.Next() {
		var t structs.Department
		if err := rows.Scan(&t.DepartmentID, &t.DepartmentName, &t.Emoji); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan departments")
		}
		departments = append(departments, t)
	}

	c.JSON(http.StatusOK, departments)
	return departments, nil
}

//////////////////////////////////////////////////////////////////////////

//////// UNITS ////////

// Get a unit from the database
func GetUnit(c *gin.Context, unitID int) (*structs.Unit, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get unit handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var unit structs.Unit
	err := db.QueryRow("SELECT id, unit_name, emoji FROM units WHERE id = ?", unitID).
		Scan(&unit.UnitID, &unit.UnitName, &unit.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return nil, fmt.Errorf("unit not found")
	}
	c.JSON(http.StatusOK, unit)
	return &unit, nil
}

// Update a unit by ID
func UpdateUnit(c *gin.Context, unitID int, unitName string, emoji string) (*structs.Unit, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from update unit handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := unitID
	var unit structs.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE units SET unit_name = ?, emoji = ? WHERE id = ?", unitName, emoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update unit")
	}
	c.JSON(http.StatusOK, "Unit updated successfully")
	return &unit, nil
}

// Delete a unit by ID
func DeleteUnit(c *gin.Context, unitID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from delete unit handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := unitID
	var status string
	_, err := db.Exec("DELETE FROM units WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Unit deletion failed"
		return &status, err
	}
	status = "Unit deleted successfully"
	return &status, nil
}

// Create unit
func CreateUnit(c *gin.Context, unitName string, emoji string) (*structs.Unit, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from create unit handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var unit structs.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request data")
	}
	result, err := db.Exec("INSERT INTO units (unit_name, emoji) VALUES (?, ?)", unitName, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create unit")
	}

	lastInsertID, _ := result.LastInsertId()
	unit.UnitID = int(lastInsertID)
	c.JSON(http.StatusCreated, unit)

	c.JSON(http.StatusOK, "Unit created successfully")

	return &unit, nil
}

// List all units
func ListUnit(c *gin.Context) ([]structs.Unit, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from list unit handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, unit_name, emoji FROM units")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch units")
	}
	defer rows.Close()

	var units []structs.Unit
	for rows.Next() {
		var unit structs.Unit
		if err := rows.Scan(&unit.UnitID, &unit.UnitName, &unit.Emoji); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan units")
		}
		units = append(units, unit)
	}

	c.JSON(http.StatusOK, units)
	return units, nil
}

//////////////////////////////////////////////////////////////////////////
//////////////// ROLES //////////////////////////////////////////////////////////

// Get a role from the database
func GetRole(c *gin.Context, roleID int) (*structs.Role, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get role handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var role structs.Role
	err := db.QueryRow("SELECT id, role_name FROM roles WHERE id = ?", roleID).
		Scan(&role.RoleID, &role.RoleName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return nil, fmt.Errorf("role not found")
	}
	c.JSON(http.StatusOK, role)
	return &role, nil
}

// Update a role by ID
func UpdateRole(c *gin.Context, roleID int, roleName string) (*structs.Role, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from update role handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := roleID
	var role structs.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE roles SET role_name = ? WHERE id = ?", roleName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update role")
	}
	c.JSON(http.StatusOK, "Role updated successfully")
	return &role, nil
}

// Delete a role by ID
func DeleteRole(c *gin.Context, roleID int) (*string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from delete role handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := roleID
	var status string
	_, err := db.Exec("DELETE FROM roles WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Role deletion failed"
		return &status, err
	}
	status = "Role deleted successfully"
	return &status, nil
}

// Create Role
func CreateRole(c *gin.Context, roleName string) (*structs.Role, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from create role handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var role structs.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	result, err := db.Exec("INSERT INTO roles (role_name) VALUES (?)", roleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create role")
	}

	lastInsertID, _ := result.LastInsertId()
	role.RoleID = int(lastInsertID)
	c.JSON(http.StatusCreated, role)

	c.JSON(http.StatusOK, "Role created successfully")
	return &role, nil
}

// List all roles
func ListRoles(c *gin.Context) ([]structs.Role, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from list roles handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, role_name FROM roles")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch roles")
	}
	defer rows.Close()

	var roles []structs.Role
	for rows.Next() {
		var role structs.Role
		if err := rows.Scan(&role.RoleID, &role.RoleName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan roles")
		}
		roles = append(roles, role)
	}

	c.JSON(http.StatusOK, roles)
	return roles, nil
}

//////////////////////////////////////////////////////////////////////////

//////// CATEGORIES ////////

// Get a Category from database
func GetCategory(c *gin.Context, categoryID int) (*structs.Category, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var category structs.Category
	err := db.QueryRow("SELECT id, category_name FROM category WHERE id = ?", categoryID).
		Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return nil, fmt.Errorf("category not found")
	}
	c.JSON(http.StatusOK, category)
	return &category, nil
}

// Update a category by ID
func UpdateCategory(c *gin.Context, cid int) (*structs.Category, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := cid
	var s structs.Category
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE category SET category_name = ? WHERE id = ?", s.CategoryName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update category")
	}
	c.JSON(http.StatusOK, "Category updated successfully")
	return &s, nil
}

// Delete a category by ID
func DeleteCategory(c *gin.Context, cid int) (*string, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := cid
	var status string
	_, err := db.Exec("DELETE FROM category WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Category deletion failed"
		return &status, err
	}
	status = "Category deleted successfully"
	return &status, nil
}

// Create category
func CreateCategory(c *gin.Context) (*structs.Category, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var category structs.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("invalid request")
	}
	result, err := db.Exec("INSERT INTO category (category_name) VALUES (?)", category.CategoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create category")
	}

	lastInsertID, _ := result.LastInsertId()
	category.CategoryID = int(lastInsertID)
	c.JSON(http.StatusCreated, category)

	c.JSON(http.StatusOK, "Category created successfully")
	return &category, nil
}

// List all categories
func ListCategories(c *gin.Context) ([]structs.Category, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, category_name FROM category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to fetch categories")
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan categories")
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
	return categories, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
/*-------------------------------------------------------------------------------------------------*/
//////// SUB-CATEGORIES /////////////////////////////////////////////////////////////////////////////

// Get a Subcategory from database
func GetSubCategory(c *gin.Context, subCategoryID int) (*structs.SubCategory, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var subCategory structs.SubCategory
	err := db.QueryRow("SELECT id, sub_category_name, category_id FROM sub_category WHERE id = ?", subCategoryID).
		Scan(&subCategory.SubCategoryID, &subCategory.SubCategoryName, &subCategory.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub_Category not found"})
		return nil, fmt.Errorf("sub-category not found")
	}
	c.JSON(http.StatusOK, subCategory)
	return &subCategory, nil
}

// Update a Subcategory by ID
func UpdateSubCategory(c *gin.Context, scid int) (*structs.SubCategory, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := scid
	var s structs.SubCategory
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE sub_category SET sub_category_name = ?, category_id = ? WHERE id = ?", s.SubCategoryName, s.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update sub-category")
	}
	c.JSON(http.StatusOK, "Sub-Category updated successfully")
	return &s, nil
}

// Delete a Subcategory by ID
func DeleteSubCategory(c *gin.Context, scid int) (*string, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := scid
	var status string
	_, err := db.Exec("DELETE FROM sub_category WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Sub-Category deletion failed"
		return &status, err
	}
	status = "Sub-Category deleted successfully"
	return &status, nil
}

// Create Subcategory
func CreateSubCategory(c *gin.Context, subCategoryName string, categoryID int) (*structs.SubCategory, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var subCategory structs.SubCategory
	subCategory.SubCategoryName = subCategoryName
	subCategory.CategoryID = categoryID

	result, err := db.Exec("INSERT INTO sub_category (sub_category_name, category_id) VALUES (?, ?)", subCategory.SubCategoryName, subCategory.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create sub-category")
	}

	lastInsertID, _ := result.LastInsertId()
	subCategory.SubCategoryID = int(lastInsertID)
	c.JSON(http.StatusCreated, subCategory)

	c.JSON(http.StatusOK, "Sub-Category created successfully")
	return &subCategory, nil
}

// List all SubCategories
func ListSubCategories(c *gin.Context) ([]structs.SubCategory, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, sub_category_name, category_id FROM sub_category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to query sub-categories")
	}
	defer rows.Close()

	var subCategories []structs.SubCategory
	for rows.Next() {
		var subCategory structs.SubCategory
		if err := rows.Scan(&subCategory.SubCategoryID, &subCategory.SubCategoryName, &subCategory.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan sub-category")
		}
		subCategories = append(subCategories, subCategory)
	}

	c.JSON(http.StatusOK, subCategories)
	return subCategories, nil
}

/*//////////////////////////////////////////////////////////////////////////*/

//////// STATUS ////////

// Get a Status by ID
func GetStatus2(c *gin.Context, sid int) (*structs.Status, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get status handler"})
		fmt.Println("Unable to reach DB from get status handler")
	}

	id := sid
	rows, err := db.Query("SELECT * FROM status WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return nil, err
	}
	for rows.Next() {
		return scanners.ScanIntoStatus(rows)
	}
	return nil, fmt.Errorf("status %d not found", id)
}

// Update a Status by ID
func UpdateStatus(c *gin.Context, sid int) (*structs.Status, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := sid
	var s structs.Status
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}
	_, err := db.Exec("UPDATE status SET status_name = ? WHERE id = ?", s.StatusName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update status")
	}
	c.JSON(http.StatusOK, "Status updated successfully")
	return &s, nil
}

// Delete a Status by ID
func DeleteStatus(c *gin.Context, sid int) (*string, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	id := sid
	var status string
	_, err := db.Exec("DELETE FROM status WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Status deletion failed"
		return &status, err
	}
	status = "Status deleted successfully"
	return &status, nil
}

// Create Status
func CreateStatus(c *gin.Context, statusName string) (*structs.Status, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	var status structs.Status
	status.StatusName = statusName

	result, err := db.Exec("INSERT INTO status (status_name) VALUES (?)", status.StatusName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create status")
	}

	lastInsertID, _ := result.LastInsertId()
	status.StatusID = int(lastInsertID)
	c.JSON(http.StatusCreated, status)

	c.JSON(http.StatusOK, "Status created successfully")
	return &status, nil
}

// List all Status
func ListStatus(c *gin.Context) ([]structs.Status, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, status_name FROM status")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to query status")
	}
	defer rows.Close()

	var statusList []structs.Status
	for rows.Next() {
		var status structs.Status
		if err := rows.Scan(&status.StatusID, &status.StatusName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("failed to scan status")
		}
		statusList = append(statusList, status)
	}

	c.JSON(http.StatusOK, statusList)
	return statusList, nil
}
