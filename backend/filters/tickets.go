package filters

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get a ticket by Staff
func GetTicketByStaff(c *gin.Context, stafftid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := stafftid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE staff_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Agents
func GetTicketByAgents(c *gin.Context, agentid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := agentid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE agent_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a tickets by Category
func GetTicketByCategory(c *gin.Context, categoryid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := categoryid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE category_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Sub-Category
func GetTicketBySubCategory(c *gin.Context, subcategoryid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := subcategoryid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE sub_category_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Priotrity
func GetTicketByPriority(c *gin.Context, priorityid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := priorityid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE priority_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by SLA
func GetTicketBySLA(c *gin.Context, slaid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := slaid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE sla_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Status
func GetTicketByStatus(c *gin.Context, statusid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := statusid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE status = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Site
func GetTicketBySite(c *gin.Context, siteid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := siteid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE site = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Tag
func GetTicketByTags(c *gin.Context, tagid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := tagid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE tag = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Asset
func GetTicketByAsset(c *gin.Context, assetid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := assetid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE asset_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a Ticket by Created Date/Time
func GetTicketByCreatedDate(c *gin.Context, createddate time.Time) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	date := createddate

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE created_at = ?", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a Ticket by Due Date/Time
func GetTicketByDueDate(c *gin.Context, duedate time.Time) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	date := duedate

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE due_at = ?", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a Ticket by Updated Date/Time
func GetTicketByUpdatedDate(c *gin.Context, updateddate time.Time) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	date := updateddate

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE updated_at = ?", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Related Tickets
func GetTicketByRelatedTickets(c *gin.Context, relatedticketid int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	id := relatedticketid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}
	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE related_ticket_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, tickets
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
			//continue // break here
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Department
func GetTicketsByDepartment(c *gin.Context, dept int) (int, []structs.Ticket) {
	var tickets []structs.Ticket
	var exist int
	department := dept
	status, staff := GetStaffByDepartment(c, department)

	if status == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Staff in this department"})
		return 0, tickets
	}

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, tickets
	}

	for _, eachStaff := range staff {
		id := eachStaff.StaffID
		rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets WHERE staff_id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue // break here
		}
		defer rows.Close()

		for rows.Next() {
			var t structs.Ticket
			if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				continue // break here
			}
			tickets = append(tickets, t)
		}
	}

	if len(tickets) > 0 {
		c.JSON(http.StatusOK, tickets)
		exist = 1
	} else {
		exist = 0
	}

	return exist, tickets
}
