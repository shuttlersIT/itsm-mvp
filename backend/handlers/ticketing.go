package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/scanners"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
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
func DeleteTicketOperation(c *gin.Context, tid int) (bool, error) {
	id := tid

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return false, fmt.Errorf("unable to reach DB")
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false, err
	}
	return true, nil
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
func ListTickets(c *gin.Context) ([]*structs.Ticket, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	rows, err := db.Query("SELECT id, title, description, status FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("unable to find tickets")
	}
	defer rows.Close()

	var tickets []*structs.Ticket
	for rows.Next() {
		//var t *structs.Ticket
		t, err := scanners.ScanIntoTicket(rows)

		if err != nil {
			return nil, fmt.Errorf("unable to add tickets to array")
		}
		tickets = append(tickets, t)
	}
	c.JSON(http.StatusOK, "Tickets Listed successfully")
	return tickets, nil
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
func GetTicket(c *gin.Context, tid int) (*structs.Ticket, error) {
	id := tid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to get ticket")
	}
	rows, err := db.Query("SELECT * FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return nil, fmt.Errorf("unable to get ticket")
	}
	for rows.Next() {
		t, e := scanners.ScanIntoTicket(rows)
		if e != nil {
			c.JSON(http.StatusNotFound, "an error occured when getting ticket from db")
			return nil, e
		} else {
			c.JSON(http.StatusOK, "ticket retrieval successfull")
			return t, nil
		}
	}
	return nil, fmt.Errorf("ticket %d not found", id)
}

// Update a ticket by ID
func DeleteTicket2(c *gin.Context) {
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

// ListTickets retrieves a paginated list of tickets from the database
func ListPageTickets(c *gin.Context, offset, perPage int) ([]*structs.Ticket, error) {
	var tickets []*structs.Ticket

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("unable to reach DB")
	}

	// Use OFFSET and LIMIT in the SQL query for pagination
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets ORDER BY created_at DESC OFFSET $1 LIMIT $2", offset, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t *structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.AttachmentID); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}
