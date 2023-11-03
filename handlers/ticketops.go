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

//////// SLA /////////////////////////////////////////////////////////

// Get SLA from database
func GetSla(c *gin.Context, slaid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := slaid
	var s structs.Sla
	err := db.QueryRow("SELECT id, sla_name, priority_id, satisfaction_id, policy_id FROM sla WHERE id = ?", id).
		Scan(&s.SlaID, &s.SlaName, &s.PriorityID, &s.SatisfactionID, &s.PolicyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SLA not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update SLA
func UpdateSla(c *gin.Context, slaid int, sn string, prid int, satid int, polid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := slaid
	slaName := sn
	priorityID := prid
	satisfactionID := satid
	policyID := polid

	var s structs.Sla
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE sla SET sla_name = ?, priority_id = ?, satisfaction_id = ?, policy_id = ? WHERE id = ?", slaName, priorityID, satisfactionID, policyID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "SLA updated successfully")
}

// Delete SLA
func DeleteSla(c *gin.Context, slaid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := slaid
	_, err := db.Exec("DELETE FROM sla WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "SLA deleted successfully")
}

// Create SLA
func CreateSla(c *gin.Context, slaid int, sn string, prid int, satid int, polid int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	slaName := sn
	priorityID := prid
	satisfactionID := satid
	policyID := polid

	var s structs.Sla
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO sla (sla_name, priority_id, satisfaction_id, policy_id) VALUES (?, ?, ?, ?)", slaName, priorityID, satisfactionID, policyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.SlaID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "SLA created successfully")

	return s.SlaID
}

// List all SLA's
func ListSla(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, sla_name, priority_id, satisfaction_id, policy_id FROM sla")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var slas []structs.Sla
	for rows.Next() {
		var t structs.Sla
		if err := rows.Scan(&t.SlaID, &t.PriorityID, &t.SatisfactionID, &t.PolicyID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		slas = append(slas, t)
	}

	c.JSON(http.StatusOK, slas)
}

//////////////////////////////////////////////////////////////////////////

//////// PRIORITY ////////

// Get a user ID from database
func GetPriority(c *gin.Context, pid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	var s structs.Priority
	err := db.QueryRow("SELECT id, priority_name, first_response, colour FROM priority WHERE id = ?", id).
		Scan(&s.PriorityID, &s.Name, &s.FirstResponse, &s.Colour)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Priority not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a priority by ID
func UpdatePriority(c *gin.Context, pid int, pn string, fr int, co string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	priority_name := pn
	first_response := fr
	colour := co

	var s structs.Priority
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE priority SET priority_name = ?, first_response = ?, colour = ? WHERE id = ?", priority_name, first_response, colour, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Priority updated successfully")
}

// Delete a priority by ID
func DeletePriority(c *gin.Context, pid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	_, err := db.Exec("DELETE FROM priority WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Priority deleted successfully")
}

// Create Priority
func CreatePriority(c *gin.Context, pn string, fr int, co string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	priority_name := pn
	first_response := fr
	colour := co

	var s structs.Priority
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO priority (priority_name, first_response, colour) VALUES (?, ?, ?)", priority_name, first_response, colour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.PriorityID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Priority created successfully")

	return s.PriorityID
}

// List all Priority Ranks
func ListPriorities(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, priority_name, first_response, colour FROM priority")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var prioritys []structs.Priority
	for rows.Next() {
		var t structs.Priority
		if err := rows.Scan(&t.PriorityID, &t.Name, &t.FirstResponse, &t.Colour); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		prioritys = append(prioritys, t)
	}

	c.JSON(http.StatusOK, prioritys)
}

//////// SATISFACTION ////////

//////////////////////////////////////////////////////////////////////////

// Get Satisfaction rank from database
func GetSatisfaction(c *gin.Context, sid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	var s structs.Satisfaction
	err := db.QueryRow("SELECT id, satisfaction_name, rank, emoji FROM satisfaction WHERE id = ?", id).
		Scan(&s.SatisfactionID, &s.Rank, &s.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satisfaction Rank not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a Satisfaction Rank by ID
func UpdateSatisfaction(c *gin.Context, sid int, sn string, srank int, em string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	satisfactionName := sn
	rank := srank
	emoji := em
	var s structs.Satisfaction
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE satisfaction SET satisfaction_name = ?, rank = ?, emoji WHERE id = ?", satisfactionName, rank, emoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Satisfaction Rank updated successfully")
}

// Delete a Satisfaction Rank by ID
func DeleteSatisfaction(c *gin.Context, sid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	_, err := db.Exec("DELETE FROM satisfaction WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Satisfaction Rank deleted successfully")
}

// Create Satisfaction Rank
func CreateSatisfaction(c *gin.Context, sn string, srank int, em string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	satisfactionName := sn
	rank := srank
	emoji := em

	var s structs.Satisfaction
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO satisfaction (satisfaction_name, rank, emoji) VALUES (?, ?, ?)", satisfactionName, rank, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.SatisfactionID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Satisfaction created successfully")

	return s.SatisfactionID
}

// List all Satisfaction Ranks
func ListSatisfaction(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, satisfaction_name, rank, emoji FROM satisfaction")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var satisfactions []structs.Satisfaction
	for rows.Next() {
		var t structs.Satisfaction
		if err := rows.Scan(&t.SatisfactionID, &t.Name, &t.Emoji); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		satisfactions = append(satisfactions, t)
	}

	c.JSON(http.StatusOK, satisfactions)
}

//////////////////////////////////////////////////////////////////////////

//////// POLICY ////////

// Get a policy from database
func GetPolicy(c *gin.Context, poid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := poid
	var s structs.Policies
	err := db.QueryRow("SELECT id, policy_name, embedded_link, policy_url FROM policy WHERE id = ?", id).
		Scan(&s.PolicyID, &s.PolicyName, &s.EmbeddedLink, &s.PolicyUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a policy by ID
func UpdatePolicy(c *gin.Context, poid int, pon string, pel string, purl string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := poid
	policyName := pon
	embeddedLink := pel
	policyUrl := purl
	var s structs.Policies
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE policy SET policy_name = ?, embedded_link = ?, policy_url = ? WHERE id = ?", policyName, embeddedLink, policyUrl, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Policy updated successfully")
}

// Delete a policy by ID
func DeletePolicy(c *gin.Context, poid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	session := sessions.Default(c)
	id := session.Get("user-id")
	_, err := db.Exec("DELETE FROM policy WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User deleted successfully")
}

// Create policy
func CreatePolicy(c *gin.Context, pon string, pel string, purl string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	policyName := pon
	embeddedLink := pel
	policyUrl := purl

	var s structs.Policies
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO policy (policy_name, embedded_link, policy_url) VALUES (?, ?, ?)", policyName, embeddedLink, policyUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.PolicyID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Policy created successfully")

	return s.PolicyID
}

// List all policies
func ListPolicies(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, policy_name, embedded_link, policy_url FROM policy")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var policies []structs.Policies
	for rows.Next() {
		var t structs.Policies
		if err := rows.Scan(&t.PolicyID, &t.PolicyName, &t.EmbeddedLink, &t.PolicyUrl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		policies = append(policies, t)
	}

	c.JSON(http.StatusOK, policies)
}

//////////////////////////////////////////////////////////////////////////

//////// POSITIONS ////////

// Get positions from database
func GetPosition(c *gin.Context, pid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	var s structs.Position
	err := db.QueryRow("SELECT id, position_name, emoji FROM positions WHERE id = ?", id).
		Scan(&s.PositionID, &s.PositionName, &s.CadreName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a position by ID
func UpdatePosition(c *gin.Context, pid int, pn string, cn string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	positionName := pn
	cadreName := cn
	var s structs.Position
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE positions SET position_name = ?, cadre_name = ? WHERE id = ?", positionName, cadreName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Position updated successfully")
}

// Delete a position by ID
func DeletePosition(c *gin.Context, pid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := pid
	_, err := db.Exec("DELETE FROM positions WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Position deleted successfully")
}

// Create position
func CreatePosition(c *gin.Context, pn string, pcn string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	position_name := pn
	cadre_name := pcn

	var s structs.Position
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO positions (position_name, cadre_name) VALUES (?, ?)", position_name, cadre_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.PositionID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Position created successfully")

	return s.PositionID
}

// List all positions
func Listpositions(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, positions_name, cadre_name FROM positions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var positions []structs.Position
	for rows.Next() {
		var t structs.Position
		if err := rows.Scan(&t.PositionID, &t.PositionName, &t.CadreName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		positions = append(positions, t)
	}

	c.JSON(http.StatusOK, positions)
}

//////////////////////////////////////////////////////////////////////////

//////// DEPARTMENTS ////////

// Get a department from database
func GetDepartment(c *gin.Context, did int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := did
	var s structs.Department
	err := db.QueryRow("SELECT id, department_name, emoji FROM departments WHERE id = ?", id).
		Scan(&s.DepartmentID, &s.DepartmentName, &s.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a department by ID
func UpdateDepartment(c *gin.Context, did int, dn string, em string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := did
	departmentName := dn
	emoji := em

	var s structs.Department
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE departments SET department_name = ?, emoji = ? WHERE id = ?", departmentName, emoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Department updated successfully")
}

// Delete a department by ID
func DeleteDepartment(c *gin.Context, did int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := did
	_, err := db.Exec("DELETE FROM departments WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Department deleted successfully")
}

// Create department
func CreateDepartment(c *gin.Context, dn string, de string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	department_name := dn
	emoji := de

	var s structs.Department
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO departments (department_name, emoji) VALUES (?, ?)", department_name, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.DepartmentID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Department created successfully")

	return s.DepartmentID
}

// List all departments
func ListDepartments(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, department_name, emoji FROM departments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var departments []structs.Department
	for rows.Next() {
		var t structs.Department
		if err := rows.Scan(&t.DepartmentID, &t.DepartmentName, &t.Emoji); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		departments = append(departments, t)
	}

	c.JSON(http.StatusOK, departments)
}

//////////////////////////////////////////////////////////////////////////

//////// UNITS ////////

// Get a unit from database
func GetUnit(c *gin.Context, uid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := uid
	var s structs.Unit
	err := db.QueryRow("SELECT id, unit_name, emoji FROM units WHERE id = ?", id).
		Scan(&s.UnitID, &s.UnitName, &s.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a unit by ID
func UpdateUnit(c *gin.Context, uid int, un string, ue string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := uid
	unitName := un
	emoji := ue

	var s structs.Unit
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE units SET unit_name = ?, emoji = ? WHERE id = ?", unitName, emoji, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Unit updated successfully")
}

// Delete a unit by ID
func DeleteUnit(c *gin.Context, uid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := uid
	_, err := db.Exec("DELETE FROM units WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Unit deleted successfully")
}

// Create unit
func CreateUnit(c *gin.Context, un string, ue string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	unit_name := un
	emoji := ue

	var s structs.Unit
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO units (unit_name, emoji) VALUES (?, ?)", unit_name, emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.UnitID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Unit created successfully")

	return s.UnitID
}

// List all units
func ListUnit(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, unit_name FROM units")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var units []structs.Unit
	for rows.Next() {
		var t structs.Unit
		if err := rows.Scan(&t.UnitID, &t.UnitName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		units = append(units, t)
	}

	c.JSON(http.StatusOK, units)
}

//////////////////////////////////////////////////////////////////////////

//////// ROLES ////////

// Get a role from database
func GetRole(c *gin.Context, rid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := rid
	var s structs.Role
	err := db.QueryRow("SELECT id, role_name FROM roles WHERE id = ?", id).
		Scan(&s.RoleID, &s.RoleName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a role by ID
func UpdateRole(c *gin.Context, rid int, rn string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := rid
	roleName := rn
	var s structs.Role
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE roles SET role_name = ? WHERE id = ?", roleName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Role updated successfully")
}

// Delete a role by ID
func DeleteRole(c *gin.Context, rid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := rid
	_, err := db.Exec("DELETE FROM roles WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Role deleted successfully")
}

// Create Role
func CreateRole(c *gin.Context, rn int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	role_name := rn

	var s structs.Role
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO roles (role_name) VALUES (?)", role_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.RoleID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Role created successfully")

	return s.RoleID
}

// List all roles
func ListRoles(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, role_name FROM roles")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var roles []structs.Role
	for rows.Next() {
		var t structs.Role
		if err := rows.Scan(&t.RoleID, &t.RoleName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		roles = append(roles, t)
	}

	c.JSON(http.StatusOK, roles)
}

//////////////////////////////////////////////////////////////////////////

//////// CATEGORIES ////////

// Get a Category from database
func GetCategory(c *gin.Context, cid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := cid
	var s structs.Categories
	err := db.QueryRow("SELECT id, category_name FROM category WHERE id = ?", id).
		Scan(&s.CategoryID, &s.CategoryName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a category by ID
func UpdateCategory(c *gin.Context, cid int, cn string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := cid
	categoryName := cn
	var s structs.Categories
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE category SET category_name = ? WHERE id = ?", categoryName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Category updated successfully")
}

// Delete a category by ID
func DeleteCategory(c *gin.Context, cid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := cid
	_, err := db.Exec("DELETE FROM category WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Category deleted successfully")
}

// Create category
func CreateCategories(c *gin.Context, cn string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	category_name := cn

	var s structs.Categories
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO category (category_name) VALUES (?)", category_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.CategoryID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Category created successfully")

	return s.CategoryID
}

// List all categories
func ListCategories(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, category_name FROM category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []structs.Categories
	for rows.Next() {
		var t structs.Categories
		if err := rows.Scan(&t.CategoryID, &t.CategoryName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, t)
	}

	c.JSON(http.StatusOK, categories)
}

//////////////////////////////////////////////////////////////////////////

//////// SUB-CATEGORIES ////////

// Get a user ID from database
func GetSubCategory(c *gin.Context, scid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := scid
	var s structs.SubCategories
	err := db.QueryRow("SELECT id, sub_category_name, category_id FROM sub_category WHERE id = ?", id).
		Scan(&s.SubCategoryID, &s.SubCategoryName, &s.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub_Category not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a update by ID
func UpdateSubCategories(c *gin.Context, scid int, scn string, cid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := scid
	sub_category_name := scn
	category_id := cid

	var s structs.SubCategories
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE sub_category SET sub_category_name = ?, category_id = ? WHERE id = ?", sub_category_name, category_id, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Sub_Category updated successfully")
}

// Delete a user by ID
func DeleteSubCategories(c *gin.Context, scid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := scid
	_, err := db.Exec("DELETE FROM sub_category WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Sub-Category deleted successfully")
}

// Create staff
func CreateSubCategories(c *gin.Context, scn string, cid int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	sub_category_name := scn
	category_id := cid

	var s structs.SubCategories
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO sub_category (sub_category_name, category_id) VALUES (?, ?)", sub_category_name, category_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.SubCategoryID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Sub-Category created successfully")

	return s.SubCategoryID
}

// List all tickets
func ListSubCategories(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, sub_category_name, category_id FROM sub_category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var SubCategories []structs.SubCategories
	for rows.Next() {
		var sc structs.SubCategories
		if err := rows.Scan(&sc.SubCategoryID, &sc.SubCategoryName, &sc.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		SubCategories = append(SubCategories, sc)
	}

	c.JSON(http.StatusOK, SubCategories)
}


//////////////////////////////////////////////////////////////////////////

//////// STATUS ////////

// Get a Status from database
func GetStatus(c *gin.Context, sid int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	var s structs.Status
	err := db.QueryRow("SELECT id, status_name FROM status WHERE id = ?", id).
		Scan(&s.StatusID, &s.StatusName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a category by ID
func UpdateStatus(c *gin.Context, sid int, sn string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	statusName := sn
	var s structs.Status
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE status SET status_name = ? WHERE id = ?", statusName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Status updated successfully")
}

// Delete a Status by ID
func DeleteStatus(c *gin.Context, sid int) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	//session := sessions.Default(c)
	id := sid
	_, err := db.Exec("DELETE FROM status WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Status deleted successfully")
}

// Create Status
func CreateStatus(c *gin.Context, sn string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	status_name := sn

	var s structs.Status
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO status (status_name) VALUES (?)", status_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.StatusID = int(lastInsertID)
	c.JSON(http.StatusCreated, s)

	c.JSON(http.StatusOK, "Status created successfully")

	return s.StatusID
}

// List all categories
func ListCategories(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, status_name FROM status")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var status []structs.Status
	for rows.Next() {
		var t structs.Status
		if err := rows.Scan(&t.StatusID, &t.StatusName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		status = append(status, t)
	}

	c.JSON(http.StatusOK, status)
}