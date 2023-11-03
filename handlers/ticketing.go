package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
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

// Update a ticket by ID
func UpdateTicketOperation(c *gin.Context, tid int, ts string, td int, cat int, scat int, pr string, slaid int, staffid int, agentid int, due ItTime, asset int, relatedTid int, tag string, site string, status string attachmentid int) int {
	id := tid
	subject := ts
	description := td
	category := cat
	subCategory := scat
	priority := pr
	sLA := slaid
	staffID := staffid
	agentID := agentid
	dueAt := due
	assetID := assetid
	relatedTicketID := relatedTid
	tag := tag
	site := site
	status := status
	attachmentID := attachmentid

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

	_, err := db.Exec("UPDATE tickets SET subject = ?, description = ?, category_id = ?, sub_category_id = ?, priority_id, sla_id, staff_id, agent_id, due_at, asset_id, related_ticket_id, tag, site, attachment_id WHERE id = ?", t.Subject, t.Description, t.Status, id)
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

// Delete a ticket by ID
func UpdateTicket(c *gin.Context) {
	//id := tid

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
	//id := tid

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
