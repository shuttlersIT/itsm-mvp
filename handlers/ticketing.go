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
// List all tickets
func ListTicketsOperation(c *gin.Context) ([]*structs.Ticket, error) {
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
func CreateTicketOperation(c *gin.Context, t structs.Ticket) (*structs.Ticket, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}

	result, err := db.Exec("INSERT INTO tickets (subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		t.Subject, t.Description, t.Category, t.SubCategory, t.Priority, t.SLA, t.StaffID, t.AgentID, t.DueAt, t.AssetID, t.RelatedTicketID, t.Tag, t.Site, t.Status, t.AttachmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create ticket")
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
	return &t, nil
}

// Get a ticket by ID
func GetTicketOperation(c *gin.Context, tid int) (*structs.Ticket, error) {
	id := tid

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to read DB")
	}
	rows, err := db.Query("SELECT * FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return nil, fmt.Errorf("ticket not found")
	}

	for rows.Next() {
		return scanners.ScanIntoTicket(rows)
	}

	return nil, fmt.Errorf("ticket %d not found", id)
}

// Update a ticket by ID
func UpdateTicketOperation(c *gin.Context, t structs.Ticket) (*structs.Ticket, error) {
	id := t.ID

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}

	_, err := db.Exec("UPDATE tickets SET subject = ?, description = ?, category_id = ?, sub_category_id = ?, priority_id = ?, sla_id = ?, staff_id = ?, agent_id = ?, due_at = ?, asset_id = ?, related_ticket_id = ?, tag = ?, site = ?, status = ?, attachment_id = ? WHERE id = ?",
		t.Subject, t.Description, t.Category, t.SubCategory, t.Priority, t.SLA, t.StaffID, t.AgentID, t.DueAt, t.AssetID, t.RelatedTicketID, t.Tag, t.Site, t.Status, t.AttachmentID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to update ticket")
	}
	c.JSON(http.StatusOK, "Ticket updated successfully")
	return &t, nil
}

// Delete a ticket by ID
func DeleteTicketOperation(c *gin.Context, tid int) (*string, error) {
	id := tid
	var status string

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach DB")
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		status = "Ticket deletion failed"
		return &status, err
	}
	status = "Ticket deleted successfully"
	return &status, nil
}

/*----------------------------------------------------------------------------------------------------------------------------------------*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*----------------------------------------------------------------------------------------------------------------------------------------*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*----------------------------------------------------------------------------------------------------------------------------------------*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*----------------------------------------------------------------------------------------------------------------------------------------*/

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
func CreateTicket(c *gin.Context, t structs.Ticket) (*structs.Ticket, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("bad request")
	}

	result, err := db.Exec("INSERT INTO tickets (subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		t.Subject, t.Description, t.Category, t.SubCategory, t.Priority, t.SLA, t.StaffID, t.AgentID, t.DueAt, t.AssetID, t.RelatedTicketID, t.Tag, t.Site, t.Status, t.AttachmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to create ticket")
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
	return &t, nil
}

// Get a ticket by ID
func GetTicket2(c *gin.Context) {
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
func GetTicket(c *gin.Context, tid int) (int, structs.Ticket) {
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
