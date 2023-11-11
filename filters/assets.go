package filters

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/handlers"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

// Get a Asset by Staff
func GetAssetByStaff(c *gin.Context, staffaid int) (int, []structs.Asset) {
	var assets []structs.Asset
	staffid := staffaid
	exists := 0
	status := 0

	var assignments []structs.AssetAssignment

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, assets
	}
	rows, err := db.Query("SELECT id, asset_id, staff_id, assigned_by, assignment_type FROM asset_assignment WHERE staff_id = ?", staffid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, assets
	}
	defer rows.Close()

	for rows.Next() {
		var a structs.AssetAssignment
		if err := rows.Scan(&a.AssignmentID, &a.AssetID, &a.StaffID, &a.AssignedBy, &a.AssignmentType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, assets
		}
		assignments = append(assignments, a)
	}

	for _, eachAsset := range assignments {
		is, asset := handlers.GetAsset2(c, eachAsset.AssetID)

		if is > 0 {
			assets = append(assets, asset)
			exists = 1
		}
	}

	if exists > 0 {
		status = 1
	}

	c.JSON(http.StatusOK, assets)
	return status, assets
}

// Get a ticket by Agents
func GetAssetByAgents(c *gin.Context, aid int) (int, []structs.Asset) {
	var assets []structs.Asset
	agentid := aid
	exists := 0
	status := 0

	var assignments []structs.AssetAssignment

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, assets
	}
	rows, err := db.Query("SELECT id, asset_id, staff_id, assigned_by, assignment_type FROM asset_assignment WHERE assigned_by = ?", agentid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, assets
	}
	defer rows.Close()

	for rows.Next() {
		var a structs.AssetAssignment
		if err := rows.Scan(&a.AssignmentID, &a.AssetID, &a.StaffID, &a.AssignedBy, &a.AssignmentType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, assets
		}
		assignments = append(assignments, a)
	}

	for _, eachAsset := range assignments {
		is, asset := handlers.GetAsset2(c, eachAsset.AssetID)

		if is > 0 {
			assets = append(assets, asset)
			exists = 1
		}
	}

	if exists > 0 {
		status = 1
	}

	c.JSON(http.StatusOK, assets)
	return status, assets
}

// Get assets by Category
func GetAssetByCategory(c *gin.Context, categoryid int) (int, []structs.Asset) {
	var assets []structs.Asset
	var assetIDs []int
	categoryID := categoryid
	exists := 0
	status := 0

	stat, tickets := GetTicketByCategory(c, categoryID)
	if stat > 0 {
		for _, eachTicket := range tickets {
			if eachTicket.AssetID == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "No Asset Related To This Ticket"})
				continue
			}
			assetIDs = append(assetIDs, eachTicket.ID)

			if len(assetIDs) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "No Tickets Related to This Asset, under this category"})
				continue
			}
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Ticket under this category"})
		return 0, assets
	}
	for _, eachAssetID := range assetIDs {
		is, asset := handlers.GetAsset2(c, eachAssetID)

		if is > 0 {
			assets = append(assets, asset)
			exists = 1
		}
	}

	if exists > 0 {
		status = 1
	}

	c.JSON(http.StatusOK, assets)
	return status, assets
}

// Get a ticket by Sub-Category
func GetAssetBySubCategory(c *gin.Context, subcategoryid int) (int, []structs.Asset) {
	var assets []structs.Asset
	var assetIDs []int
	subCategoryID := subcategoryid
	exists := 0
	status := 0

	stat, tickets := GetTicketBySubCategory(c, subCategoryID)
	if stat > 0 {
		for _, eachTicket := range tickets {
			if eachTicket.AssetID == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "No Asset Related To This Ticket"})
				continue
			}
			assetIDs = append(assetIDs, eachTicket.ID)

			if len(assetIDs) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "No Tickets Related to This Asset, under this sub-category"})
				continue
			}
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Ticket under this sub-category"})
		return 0, assets
	}
	for _, eachAssetID := range assetIDs {
		is, asset := handlers.GetAsset2(c, eachAssetID)

		if is > 0 {
			assets = append(assets, asset)
			exists = 1
		}
	}

	if exists > 0 {
		status = 1
	}

	c.JSON(http.StatusOK, assets)
	return status, assets
}

// Get a ticket by Priotrity
func GetAssetByPriority(c *gin.Context, priorityid int) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetBySLA(c *gin.Context, slaid int) (int, []structs.Asset) {
	var tickets []structs.Asset
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
		var t structs.Asset
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
func GetAssetByStatus(c *gin.Context, statusid int) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetBySite(c *gin.Context, siteid int) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetByTags(c *gin.Context, tagid int) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a ticket by Ticket
func GetAssetByTicket(c *gin.Context, assetid int) (int, []structs.Asset) {
	var tickets []structs.Asset
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
		var t structs.Asset
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
func GetAssetByCreatedDate(c *gin.Context, createddate time.Time) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetByDueDate(c *gin.Context, duedate time.Time) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetByUpdatedDate(c *gin.Context, updateddate time.Time) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, tickets
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return 1, tickets
}

// Get a Assets by Related Assets
func GetAssetByRelatedAssets(c *gin.Context, relatedticketid int) (int, []structs.Asset) {
	var assets []structs.Asset
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
		var t structs.Asset
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
func GetAssetsByDepartment(c *gin.Context, dept int) (int, []structs.Asset) {
	var assets []structs.Asset
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
			var t structs.Asset
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
