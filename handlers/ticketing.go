package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/scanners"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

type ItTime time.Time

//Ticket Handlers
/*
// Ticketing Handlers
func ListTickets(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementportal.html", gin.H{"Username": userID})
}

func CreateTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementadmin.html", gin.H{"Username": userID})
}
func UpdateTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userID})
}
func DeleteTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userID})
}
*/

// List all tickets
func ListTicketsOperation(c *gin.Context) ([]*structs.Ticket, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	rows, err := db.Query("SELECT * FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("there are no tickets")
	}
	defer rows.Close()

	tickets := []*structs.Ticket{}
	for rows.Next() {
		ticket, err := scanners.ScanIntoTicket(rows)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// Create a new ticket
func CreateTicketOperation(c *gin.Context, ts string, td int, cat int, scat int, pr string, slaid int, staffid int, agentid int, due ItTime, assetid int, relatedTid int, tags string, siTe string, staTus string, attachment int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	subject := ts
	description := td
	category := cat
	subCategory := scat
	priority := pr
	sLA := slaid
	staff := staffid
	agent := agentid
	dueAt := due
	asset := assetid
	relatedT := relatedTid
	tag := tags
	site := siTe
	status := staTus
	attachmentID := attachment

	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	result, err := db.Exec("INSERT INTO tickets (subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", subject, description, category, subCategory, priority, sLA, staff, agent, dueAt, asset, relatedT, tag, site, status, attachmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
	return t.ID
}

// Get a ticket by ID
func GetTicketOperation(c *gin.Context, tid int) (*structs.Ticket, error) {
	id := tid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to read DB")
	}
	rows, err := db.Query("SELECT * FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return nil, err
	}

	for rows.Next() {
		return scanners.ScanIntoTicket(rows)
	}

	return nil, fmt.Errorf("ticket %d not found", id)
}

// Update a ticket by ID
func UpdateTicketOperation(c *gin.Context, tid int, ts string, td int, cat int, scat int, pr string, slaid int, staffid int, agentid int, due ItTime, assetid int, relatedTid int, tags string, siTe string, staTus string, attachment int) int {
	id := tid
	subject := ts
	description := td
	category := cat
	subCategory := scat
	priority := pr
	sLA := slaid
	staff := staffid
	agent := agentid
	dueAt := due
	asset := assetid
	relatedT := relatedTid
	tag := tags
	site := siTe
	status := staTus
	attachmentID := attachment

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}
	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	_, err := db.Exec("UPDATE tickets SET subject = ?, description = ?, category_id = ?, sub_category_id = ?, priority_id = ?, sla_id = ?, staff_id = ?, agent_id = ?, due_at = ?, asset_id = ?, related_ticket_id = ?, tag = ?, site = ?, status = ?, attachment_id = ? WHERE id = ?", subject, description, category, subCategory, priority, sLA, staff, agent, dueAt, asset, relatedT, tag, site, status, attachmentID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}
	c.JSON(http.StatusOK, "Ticket updated successfully")
	return t.ID
}

// Delete a ticket by ID
func DeleteTicketOperation(c *gin.Context, tid int) {
	id := tid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket deleted successfully")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*----------------------------------------------------------------------------------------------------------------------------------------*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// List all tickets
func ListTickets(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, title, description, status FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tickets []structs.Ticket
	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
}

// Create a new ticket
func CreateTicket(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tickets (title, description, status) VALUES (?, ?, ?)", t.Subject, t.Description, t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
}

// Get a ticket by ID
func GetTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	var t structs.Ticket
	err := db.QueryRow("SELECT id, subject, description, status FROM tickets WHERE id = ?", id).
		Scan(&t.ID, &t.Subject, &t.Description, &t.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Get a ticket by ID
func GetTicket2(c *gin.Context, tid int) (int, structs.Ticket) {
	id := tid
	var t structs.Ticket

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, t
	}
	err := db.QueryRow("SELECT id, subject, description, status FROM tickets WHERE id = ?", id).
		Scan(&t.ID, &t.Subject, &t.Description, &t.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return 0, t
	}
	c.JSON(http.StatusOK, t)
	return 1, t
}

// Update a ticket by ID
func UpdateTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket deleted successfully")
}

// Delete a ticket by ID
func DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket deleted successfully")
}
