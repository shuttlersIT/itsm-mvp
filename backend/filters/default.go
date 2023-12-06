package filters

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------TICKET LIST--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// List 10 Most Recent Tickets
func List10MostRecentTickets(c *gin.Context) ([]structs.Ticket, error) {
	return listRecentTickets(c, 10)
}

// List 20 Most Recent Tickets
func List20MostRecentTickets(c *gin.Context) ([]structs.Ticket, error) {
	return listRecentTickets(c, 20)
}

// List 30 Most Recent Tickets
func List30MostRecentTickets(c *gin.Context) ([]structs.Ticket, error) {
	return listRecentTickets(c, 30)
}

// List 40 Most Recent Tickets
func List40MostRecentTickets(c *gin.Context) ([]structs.Ticket, error) {
	return listRecentTickets(c, 40)
}

// List 10 Most Recent Tickets
func List50MostRecentTickets(c *gin.Context) ([]structs.Ticket, error) {
	return listRecentTickets(c, 50)
}

// Helper function to list recent tickets
func listRecentTickets(c *gin.Context, limit int) ([]structs.Ticket, error) {
	var tickets []structs.Ticket

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return tickets, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, subject, description, category_id, sub_category_id, priority_id, sla_id, staff_id, agent_id, created_at, updated_at, due_at, asset_id, related_ticket_id, tag, site, status, attachment_id FROM tickets ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return tickets, err
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Category, &t.SubCategory, &t.Priority, &t.SLA, &t.StaffID, &t.AgentID, &t.CreatedAt, &t.UpdatedAt, &t.DueAt, &t.AssetID, &t.RelatedTicketID, &t.Tag, &t.Site, &t.Status, &t.Status, &t.AttachmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return tickets, err
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
	return tickets, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------STAFF LIST--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// List 10 Most Recent Staff Members
func List10MostRecentStaff(c *gin.Context) ([]structs.Staff, error) {
	return listRecentStaff(c, 10)
}

// List 20 Most Recent Staff Members
func List20MostRecentStaff(c *gin.Context) ([]structs.Staff, error) {
	return listRecentStaff(c, 20)
}

// List 30 Most Recent Staff Members
func List30MostRecentStaff(c *gin.Context) ([]structs.Staff, error) {
	return listRecentStaff(c, 30)
}

// List 40 Most Recent Staff Members
func List40MostRecentStaff(c *gin.Context) ([]structs.Staff, error) {
	return listRecentStaff(c, 40)
}

// List 50 Most Recent Staff Members
func List50MostRecentStaff(c *gin.Context) ([]structs.Staff, error) {
	return listRecentStaff(c, 50)
}

// Helper function to list recent staff members
func listRecentStaff(c *gin.Context, limit int) ([]structs.Staff, error) {
	var staffList []structs.Staff

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get staff handler"})
		return staffList, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, name, email, phone, department_id, position_id, created_at, updated_at FROM staff ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return staffList, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff structs.Staff
		if err := rows.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName, &staff.StaffEmail, &staff.Phone, &staff.DepartmentID, &staff.PositionID, &staff.CreatedAt, &staff.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return staffList, err
		}
		staffList = append(staffList, staff)
	}

	c.JSON(http.StatusOK, staffList)
	return staffList, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------AGENTS LIST--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// List 10 Most Recent Agents
func List10MostRecentAgents(c *gin.Context) ([]structs.Agent, error) {
	return listRecentAgents(c, 10)
}

// List 20 Most Recent Agents
func List20MostRecentAgents(c *gin.Context) ([]structs.Agent, error) {
	return listRecentAgents(c, 20)
}

// List 30 Most Recent Agents
func List30MostRecentAgents(c *gin.Context) ([]structs.Agent, error) {
	return listRecentAgents(c, 30)
}

// List 40 Most Recent Agents
func List40MostRecentAgents(c *gin.Context) ([]structs.Agent, error) {
	return listRecentAgents(c, 40)
}

// List 50 Most Recent Agents
func List50MostRecentAgents(c *gin.Context) ([]structs.Agent, error) {
	return listRecentAgents(c, 50)
}

// Helper function to list recent agents
func listRecentAgents(c *gin.Context, limit int) ([]structs.Agent, error) {
	var agentList []structs.Agent

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get agent handler"})
		return agentList, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, name, email, phone, department_id, position_id, created_at, updated_at FROM agents ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return agentList, err
	}
	defer rows.Close()

	for rows.Next() {
		var agent structs.Agent
		if err := rows.Scan(&agent.AgentID, &agent.FirstName, &agent.LastName, &agent.AgentEmail, &agent.Phone, &agent.RoleID, &agent.SupervisorID, &agent.Unit, &agent.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return agentList, err
		}
		agentList = append(agentList, agent)
	}

	c.JSON(http.StatusOK, agentList)
	return agentList, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------ASSETS LIST--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// List 10 Most Recent Assets
func List10MostRecentAssets(c *gin.Context) ([]structs.Asset, error) {
	return listRecentAssets(c, 10)
}

// List 20 Most Recent Assets
func List20MostRecentAssets(c *gin.Context) ([]structs.Asset, error) {
	return listRecentAssets(c, 20)
}

// List 30 Most Recent Assets
func List30MostRecentAssets(c *gin.Context) ([]structs.Asset, error) {
	return listRecentAssets(c, 30)
}

// List 40 Most Recent Assets
func List40MostRecentAssets(c *gin.Context) ([]structs.Asset, error) {
	return listRecentAssets(c, 40)
}

// List 50 Most Recent Assets
func List50MostRecentAssets(c *gin.Context) ([]structs.Asset, error) {
	return listRecentAssets(c, 50)
}

// Helper function to list recent assets
func listRecentAssets(c *gin.Context, limit int) ([]structs.Asset, error) {
	var assetList []structs.Asset

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get asset handler"})
		return assetList, fmt.Errorf("unable to reach DB")
	}

	rows, err := db.Query("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by, created_at, updated_at FROM assets ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return assetList, err
	}
	defer rows.Close()

	for rows.Next() {
		var asset structs.Asset
		if err := rows.Scan(&asset.ID, &asset.AssetID, &asset.AssetType, &asset.AssetName, &asset.Description, &asset.Manufacturer, &asset.Model, &asset.SerialNumber, &asset.PurchaseDate, &asset.PurchasePrice, &asset.Vendor, &asset.Site, &asset.Status, &asset.CreatedBy, &asset.CreatedAt, &asset.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return assetList, err
		}
		assetList = append(assetList, asset)
	}

	c.JSON(http.StatusOK, assetList)
	return assetList, nil
}
