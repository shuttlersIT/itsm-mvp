package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

// CreateExample function Usage
func CreateExamplerHandler(c *gin.Context) {
	// Get the status name from the request JSON or other sources
	statusName := c.PostForm("statusName")

	// Call the CreateStatus function
	createdStatus, err := CreateStatus(c, statusName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created status
	c.JSON(http.StatusOK, gin.H{
		"message":       "Status created successfully",
		"createdStatus": createdStatus,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------AGENTS------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Example route handler for creating agent
func CreateAgentHandler(c *gin.Context) {
	// Get agent details from the request JSON or other sources
	// ...

	// Call the CreateAgent function
	createdAgent, _, err := CreateAgent(c /* pass agent details */)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created agent
	c.JSON(http.StatusOK, gin.H{
		"message":      "Agent created successfully",
		"createdAgent": createdAgent,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!------------------------------------------------------------STAFF------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateStaffHandler(c *gin.Context) {
	// Get staff details from the request JSON or other sources
	// ...

	// Call the CreateStaff function
	createdStaff, _, err := CreateUser(c /* pass staff details */)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created staff
	c.JSON(http.StatusOK, gin.H{
		"message":      "Staff created successfully",
		"createdStaff": createdStaff,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------TICKETS-----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateTicket function Usage
func CreateTicketHandler(c *gin.Context) {
	var t structs.Ticket
	// Get the ticket details name from the request JSON or other sources
	t.Subject = c.PostForm("ticketSubject")
	t.Description = c.PostForm("ticketDescription")
	t.Category, _ = strconv.Atoi(c.PostForm("ticketCategory"))
	t.SubCategory, _ = strconv.Atoi(c.PostForm("ticketSubSubject"))
	t.Priority, _ = strconv.Atoi(c.PostForm("ticketPriority"))
	t.SLA, _ = strconv.Atoi(c.PostForm("ticketSla"))
	t.StaffID, _ = strconv.Atoi(c.PostForm("ticketStaffID"))
	t.AgentID, _ = strconv.Atoi(c.PostForm("ticketAgentID"))
	t.DueAt, _ = time.Parse("02-01-2006", c.PostForm("ticketDueDate"))
	t.AssetID, _ = strconv.Atoi(c.PostForm("ticketAssetID"))
	t.RelatedTicketID, _ = strconv.Atoi(c.PostForm("ticketRelatedTicket"))
	t.Tag, _ = strconv.Atoi(c.PostForm("ticketTag"))
	t.Site, _ = strconv.Atoi(c.PostForm("ticketSite"))
	t.Status, _ = strconv.Atoi(c.PostForm("ticketStatus"))
	t.AttachmentID, _ = strconv.Atoi(c.PostForm("ticketAttachmentID"))

	// Call the CreateStatus function
	createdTicket, err := CreateTicket(c, t)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created status
	c.JSON(http.StatusOK, gin.H{
		"message":       "Ticket created successfully",
		"createdStatus": createdTicket,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------STATUS------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateStatus function Usage
func CreateStatusHandler(c *gin.Context) {
	// Get the status name from the request JSON or other sources
	statusName := c.PostForm("statusName")

	// Call the CreateStatus function
	createdStatus, err := CreateStatus(c, statusName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created status
	c.JSON(http.StatusOK, gin.H{
		"message":       "Status created successfully",
		"createdStatus": createdStatus,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------ASSETS------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Create Asset
func CreateAssetHandler(c *gin.Context) {
	var a structs.Asset
	// Get asset details from the request JSON or other sources
	a.AssetID, _ = strconv.Atoi(c.PostForm("assetID"))
	a.AssetType = c.PostForm("assetType")
	a.AssetName = c.PostForm("assetName")
	a.Description = c.PostForm("assetDescription")
	a.Manufacturer = c.PostForm("assetManufacturer")
	a.Model = c.PostForm("assetModel")
	a.SerialNumber = c.PostForm("assetSerialNumber")
	a.PurchaseDate = c.PostForm("assetPurchaseDate")
	a.PurchasePrice = c.PostForm("assetPurchasePrice")
	a.Vendor = c.PostForm("assetVendor")
	a.Site = c.PostForm("assetSite")
	a.Status = c.PostForm("assetStatus")
	//session.Get("agent-id") = c.PostForm("assetName")
	// ...

	// Call the CreateAsset function
	createdAsset, _, err := CreateAsset(c, a /* pass asset details */)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created asset
	c.JSON(http.StatusOK, gin.H{
		"message":      "Asset created successfully",
		"createdAsset": createdAsset,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------DEPARTMENTS--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateDepartment function Usage
func CreateDepartmentHandler(c *gin.Context) {
	var d structs.Department
	// Get the status name from the request JSON or other sources
	d.DepartmentName = c.PostForm("newDepartmentName")
	d.Emoji = c.PostForm("newDepartmentEmoji")

	// Call the CreateStatus function
	createdDepartment, err := CreateDepartment(c, d.DepartmentName, d.Emoji)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created status
	c.JSON(http.StatusOK, gin.H{
		"message":       "Department created successfully",
		"createdStatus": createdDepartment,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------PRIORITY----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreatePriority function Usage
func CreatePriorityHandler(c *gin.Context) {
	var p structs.Priority

	// Get the priority name from the request JSON or other sources
	p.Name = c.PostForm("priorityName")
	p.Colour = c.PostForm("priorityColour")
	p.FirstResponse, _ = strconv.Atoi(c.PostForm("priorityFirstResponse"))

	// Call the Createpriority function
	createdPriority, err := CreatePriority(c, p)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created priority
	c.JSON(http.StatusOK, gin.H{
		"message":       "Priority created successfully",
		"createdStatus": createdPriority,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------POSITIONS----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreatePosition function Usage
func CreatePositionHandler(c *gin.Context) {
	var p structs.Position

	// Get the position name from the request JSON or other sources
	p.PositionName = c.PostForm("positionName")
	p.CadreName = c.PostForm("positionName")

	// Call the CreatePosition function
	createdPosition, err := CreatePosition(c, p.PositionName, p.CadreName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the created position
	c.JSON(http.StatusOK, gin.H{
		"message":       "Priority created successfully",
		"createdStatus": createdPosition,
	})
}
