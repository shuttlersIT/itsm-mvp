package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
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

func CreateAgentHandler(c *gin.Context) {
	var a structs.Agent
	a.FirstName = c.PostForm("firstName")
	a.LastName = c.PostForm("lastName")
	a.AgentEmail = c.PostForm("agentEmail")
	a.Username, _ = strconv.Atoi(c.PostForm("username"))
	a.Phone = c.PostForm("phone")
	a.RoleID, _ = strconv.Atoi(c.PostForm("roleID"))
	a.Unit, _ = strconv.Atoi(c.PostForm("unit"))
	a.SupervisorID, _ = strconv.Atoi(c.PostForm("supervisorID"))

	createdAgent, _, err := CreateAgent(c, a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Agent created successfully",
		"createdAgent": createdAgent,
	})
}

// UpdateAgentHandler updates an agent
func UpdateAgentHandler(c *gin.Context) {
	var a structs.Agent
	// Get agent details from the request JSON or other sources
	a.AgentID, _ = strconv.Atoi(c.PostForm("agentID"))
	a.FirstName = c.PostForm("firstName")
	a.LastName = c.PostForm("lastName")
	a.AgentEmail = c.PostForm("agentEmail")
	a.Username, _ = strconv.Atoi(c.PostForm("username"))
	a.Phone = c.PostForm("phone")
	a.RoleID, _ = strconv.Atoi(c.PostForm("roleID"))
	a.Unit, _ = strconv.Atoi(c.PostForm("unit"))
	a.SupervisorID, _ = strconv.Atoi(c.PostForm("supervisorID"))

	// Call the UpdateAgent function
	id, err := UpdateAgent(c, a)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedAgent, _ := GetAgent(c, id)
	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Agent updated successfully",
		"UpdatedAgent": updatedAgent,
	})
}

// GetAgentHandler retrieves details of a specific agent by ID
func GetAgentHandler(c *gin.Context) {
	// Get agent ID from the URL parameter
	agentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	// Call the GetAgent function to retrieve the agent details
	agent, err := GetAgent(c, agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if agent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}

	// Return the agent details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Agent retrieved successfully",
		"agent":   agent,
	})
}

// ListAgentsHandler retrieves a list of all agents
func ListAgentsHandler(c *gin.Context) {
	// Call the ListAgents function to retrieve the list of agents
	agents, err := ListAgents(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of agents in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of agents retrieved successfully",
		"agents":  agents,
	})
}

// DeleteUserHandler deletes a specific staff by ID
func DeleteAgentHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	staffID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Call the DeleteUser function to delete the asset
	s, err := DeleteAgentOperation(c, staffID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Agent deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!------------------------------------------------------------STAFF------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateStaffHandler(c *gin.Context) {
	var s structs.Staff
	s.FirstName = c.PostForm("firstName")
	s.LastName = c.PostForm("lastName")
	s.StaffEmail = c.PostForm("staffEmail")
	s.Username, _ = strconv.Atoi(c.PostForm("username"))
	s.Phone = c.PostForm("phone")
	s.PositionID, _ = strconv.Atoi(c.PostForm("positionID"))
	s.DepartmentID, _ = strconv.Atoi(c.PostForm("departmentID"))

	createdStaff, _, err := CreateUser2(c, s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Staff created successfully",
		"createdStaff": createdStaff,
	})
}

// UpdateStaffHandler updates a staff member
func UpdateStaffHandler(c *gin.Context) {
	var s structs.Staff
	// Get staff details from the request JSON or other sources
	s.StaffID, _ = strconv.Atoi(c.PostForm("staffID"))
	s.FirstName = c.PostForm("firstName")
	s.LastName = c.PostForm("lastName")
	s.StaffEmail = c.PostForm("staffEmail")
	s.Username, _ = strconv.Atoi(c.PostForm("username"))
	s.Phone = c.PostForm("phone")
	s.PositionID, _ = strconv.Atoi(c.PostForm("positionID"))
	s.DepartmentID, _ = strconv.Atoi(c.PostForm("departmentID"))

	// Call the UpdateStaff function
	id, err := UpdateUser(c, s)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedStaff, e := GetUserByID(c, id)
	if e != nil {
		// Return a success response
		c.JSON(http.StatusOK, gin.H{
			"message": "Staff update failed",
		})
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Staff updated successfully",
		"createdStaff": updatedStaff,
	})
}

// GetStaffHandler retrieves details of a specific staff by ID
func GetStaffHandler(c *gin.Context) {
	// Get staff ID from the URL parameter
	staffID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Call the GetStaff function to retrieve the staff details
	staff, err := GetUserByID(c, staffID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if staff == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}

	// Return the staff details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Staff retrieved successfully",
		"staff":   staff,
	})
}

// ListStaffHandler retrieves a list of all staff
func ListStaffHandler(c *gin.Context) {
	// Call the ListStaff function to retrieve the list of staff
	staffList, err := ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of staff in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of staff retrieved successfully",
		"staff":   staffList,
	})
}

// DeleteUserHandler deletes a specific staff by ID
func DeleteStaffHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	staffID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Call the DeleteUser function to delete the asset
	s, err := DeleteUserOperation(c, staffID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Staff deleted successfully",
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
	t.Tag = c.PostForm("ticketTag")
	t.Site = c.PostForm("ticketSite")
	t.Status = c.PostForm("ticketStatus")
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

// UpdateTicketHandler updates a ticket
func UpdateTicketHandler(c *gin.Context) {
	var t structs.Ticket
	// Get ticket details from the request JSON or other sources
	t.ID, _ = strconv.Atoi(c.PostForm("ticketID"))
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
	t.Tag = c.PostForm("ticketTag")
	t.Site = c.PostForm("ticketSite")
	t.Status = c.PostForm("ticketStatus")
	t.AttachmentID, _ = strconv.Atoi(c.PostForm("ticketAttachmentID"))

	// Call the UpdateTicket function
	updatedTicket, err := UpdateTicketOperation(c, t)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":       "Ticket updated successfully",
		"updatedTicket": updatedTicket,
	})
}

// GetTicketHandler retrieves details of a specific ticket by ID
func GetTicketHandler(c *gin.Context) {
	// Get ticket ID from the URL parameter
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Call the GetTicket function to retrieve the ticket details
	ticket, err := GetTicket(c, ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the ticket details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket retrieved successfully",
		"ticket":  ticket,
	})
}

// ListTicketHandler retrieves a list of all tickets
func ListTicketHandler(c *gin.Context) {
	// Call the ListTicket function to retrieve the list of tickets
	ticketList, err := ListTickets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of tickets in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of tickets retrieved successfully",
		"tickets": ticketList,
	})
}

// ListTicketHandler retrieves a paginated list of tickets
func ListTicketsPageHandler(c *gin.Context) {
	// Define default values for pagination
	page := 1
	perPage := 20

	// Parse query parameters for pagination
	if pageStr := c.Query("page"); pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if perPageStr := c.Query("per_page"); perPageStr != "" {
		perPage, _ = strconv.Atoi(perPageStr)
	}

	// Calculate offset based on page and perPage values
	offset := (page - 1) * perPage

	// Call the ListTickets function to retrieve the paginated list of tickets
	ticketList, err := ListPageTickets(c, offset, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the paginated list of tickets in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Paginated list of tickets retrieved successfully",
		"tickets": ticketList,
	})
}

// DeleteAssetHandler deletes a specific asset by ID
func DeleteTicketHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	assetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the asset
	s, err := DeleteTicketOperation(c, assetID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket deleted successfully",
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

// UpdateStatusHandler updates status details
func UpdateStatusHandler(c *gin.Context) {
	var status structs.Status
	//var e error
	// Get status details from the request JSON or other sources
	status.StatusID, _ = strconv.Atoi(c.PostForm("statusID"))
	status.StatusName = c.PostForm("statusName")
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"badRequest": fmt.Errorf("bad request"),
		})
	}

	if status.StatusID < 1 && status.StatusName != "" {
		s, e := getStatusByName(c, status.StatusName)
		if e != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":  fmt.Sprintf("unable to find status %v", status.StatusName),
				"status": nil,
			})
		}
		st, er := UpdateStatus(c, *s)
		if er != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":  fmt.Sprintf("unable to find status %v", status.StatusName),
				"status": nil,
			})
		} else {
			sta, erro := GetStatus(c, st.StatusID)
			if erro != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":  fmt.Sprintf("unable to update status %v", status.StatusName),
					"status": nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message":            "Status updated successfully",
					"updatedSubCategory": sta,
				})
			}
		}

	} else if status.StatusID < 1 && status.StatusName == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  fmt.Sprintf("unable to find status to update %v", status.StatusName),
			"status": nil,
		})

	} else if status.StatusID > 0 {
		// Call the UpdateStatus function
		updatedStatus, err := UpdateStatus(c, status)
		if err != nil {
			// Handle the error, e.g., return an error response
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":            "Status updated successfully",
			"updatedSubCategory": updatedStatus,
		})
	}
}

// GetStatusHandler retrieves details of a specific status by ID
func GetStatusHandler(c *gin.Context) {
	// Get status ID from the URL parameter
	statusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	// Call the GetStatus function to retrieve the status details
	status, err := GetStatus(c, statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if status == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}

	// Return the status details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Status retrieved successfully",
		"status":  status,
	})
}

// ListStatusHandler retrieves a list of all statuses
func ListStatusHandler(c *gin.Context) {
	// Call the ListStatus function to retrieve the list of statuses
	statusList, err := ListStatus(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of statuses in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "List of statuses retrieved successfully",
		"statuses": statusList,
	})
}

// DeleteAssetHandler deletes a specific asset by ID
func DeleteStatusHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	assetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the asset
	s, err := DeleteStatus(c, assetID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket deleted successfully",
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
	a.AssetID = c.PostForm("assetID")
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

// UpdateAssetsHandler updates an asset
func UpdateAssetsHandler(c *gin.Context) {
	//var a frontendstructs.FrontendAsset
	var a structs.Asset
	var err error
	// Get asset details from the request JSON or other sources
	a.AssetID = c.PostForm("assetID")
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

	// Call the UpdateAssets function
	asset, err := UpdateAsset(c, a)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Asset updated successfully",
		"updatedAsset": asset,
	})
}

// UpdateAssetTypeHandler updates an asset type
func UpdateAssetTypeHandler(c *gin.Context) string {
	var assetType structs.AssetType
	// Get asset type details from the request JSON or other sources
	assetType.AssetTypeID, _ = strconv.Atoi(c.PostForm("assetTypeID"))
	assetType.AssetType = c.PostForm("typeName")

	// Call the UpdateAssetType function
	status, err := updateAssetType(c, assetType)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return status
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Asset type updated successfully",
	})
	return status
}

// GetAssetHandler retrieves details of a specific asset by ID
func GetAssetHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	assetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the GetAsset function to retrieve the asset details
	asset, err := GetAsset(c, assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if asset == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	// Return the asset details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Asset retrieved successfully",
		"asset":   asset,
	})
}

// ListAssetHandler retrieves a list of all assets
func ListAssetHandler(c *gin.Context) {
	// Call the ListAsset function to retrieve the list of assets
	assetList, status, err := ListAssets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of assets in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of assets retrieved successfully",
		"assets":  assetList,
	})
}

// DeleteAssetHandler deletes a specific asset by ID
func DeleteAssetHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	assetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the asset
	err = DeleteAsset(c, assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Asset deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------DEPARTMENTS--------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateDepartmentHandler(c *gin.Context) {
	var dep structs.Department
	dep.DepartmentName = c.PostForm("departmentName")
	dep.Emoji = c.PostForm("emoji")

	createdDepartment, err := CreateDepartment(c, dep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Department created successfully",
		"createdDepartment": createdDepartment,
	})
}

// UpdateDepartmentHandler updates department details
func UpdateDepartmentHandler(c *gin.Context) {
	var department structs.Department
	// Get department details from the request JSON or other sources
	department.DepartmentID, _ = strconv.Atoi(c.PostForm("departmentID"))
	department.DepartmentName = c.PostForm("departmentName")
	department.Emoji = c.PostForm("emoji")

	// Call the UpdateDepartment function
	updatedDepartment, err := UpdateDepartment(c, department.DepartmentID, department.DepartmentName, department.Emoji)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":           "Department updated successfully",
		"updatedDepartment": updatedDepartment,
	})
}

// GetDepartmentHandler retrieves details of a specific department by ID
func GetDepartmentHandler(c *gin.Context) {
	// Get department ID from the URL parameter
	departmentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	// Call the GetDepartment function to retrieve the department details
	department, err := GetDepartment(c, departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if department == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	// Return the department details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":    "Department retrieved successfully",
		"department": department,
	})
}

// ListDepartmentsHandler retrieves a list of all departments
func ListDepartmentsHandler(c *gin.Context) {
	// Call the ListDepartments function to retrieve the list of departments
	departments, err := ListDepartments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of departments in the response
	c.JSON(http.StatusOK, gin.H{
		"message":     "List of departments retrieved successfully",
		"departments": departments,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeleteDepartmentHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	deptID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteDepartment(c, deptID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Dept deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------PRIORITY----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreatePriorityHandler(c *gin.Context) {
	var p structs.Priority
	p.Name = c.PostForm("priorityName")
	p.FirstResponse, _ = strconv.Atoi(c.PostForm("firstResponse"))
	p.Colour = c.PostForm("colour")

	createdPriority, err := CreatePriority(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Priority created successfully",
		"createdPriority": createdPriority,
	})
}

// UpdatePriorityHandler updates a priority
func UpdatePriorityHandler(c *gin.Context) {
	var priority structs.Priority
	// Get priority details from the request JSON or other sources
	priority.PriorityID, _ = strconv.Atoi(c.PostForm("priorityID"))
	priority.Name = c.PostForm("priorityName")
	priority.FirstResponse, _ = strconv.Atoi(c.PostForm("firstResponse"))
	priority.Colour = c.PostForm("colour")

	// Call the UpdatePriority function
	updatePriority, err := UpdatePriority(c, priority.PriorityID, priority.Name, priority.FirstResponse, priority.Colour)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":        "Priority updated successfully",
		"updatePriority": updatePriority,
	})
}

// GetPriorityHandler retrieves details of a specific priority by ID
func GetPriorityHandler(c *gin.Context) {
	// Get priority ID from the URL parameter
	priorityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority ID"})
		return
	}

	// Call the GetPriority function to retrieve the priority details
	priority, err := GetPriority(c, priorityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if priority == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Priority not found"})
		return
	}

	// Return the priority details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Priority retrieved successfully",
		"priority": priority,
	})
}

// ListPrioritiesHandler retrieves a list of all priorities
func ListPrioritiesHandler(c *gin.Context) {
	// Call the ListPriorities function to retrieve the list of priorities
	priorities, err := ListPriorities(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of priorities in the response
	c.JSON(http.StatusOK, gin.H{
		"message":    "List of priorities retrieved successfully",
		"priorities": priorities,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeletePriorityHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeletePriority(c, pID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "priority deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------POSITIONS----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreatePositionHandler(c *gin.Context) {
	var pos structs.Position
	pos.PositionName = c.PostForm("positionName")
	pos.CadreName = c.PostForm("cadreName")

	createdPosition, err := CreatePosition(c, pos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Position created successfully",
		"createdPosition": createdPosition,
	})
}

// UpdatePositionHandler updates position details
func UpdatePositionHandler(c *gin.Context) {
	var position structs.Position
	// Get position details from the request JSON or other sources
	position.PositionID, _ = strconv.Atoi(c.PostForm("positionID"))
	position.PositionName = c.PostForm("positionName")
	position.CadreName = c.PostForm("cadreName")

	// Call the UpdatePosition function
	updatedPolicy, err := UpdatePosition(c, position.PositionID, position.PositionName, position.CadreName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":       "Position updated successfully",
		"updatedPolicy": updatedPolicy,
	})
}

// GetPositionHandler retrieves details of a specific position by ID
func GetPositionHandler(c *gin.Context) {
	// Get position ID from the URL parameter
	positionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}

	// Call the GetPosition function to retrieve the position details
	position, err := GetPosition(c, positionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if position == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}

	// Return the position details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Position retrieved successfully",
		"position": position,
	})
}

// ListPositionsHandler retrieves a list of all positions
func ListPositionsHandler(c *gin.Context) {
	// Call the ListPositions function to retrieve the list of positions
	positions, err := ListPositions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of positions in the response
	c.JSON(http.StatusOK, gin.H{
		"message":   "List of positions retrieved successfully",
		"positions": positions,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeletePositionHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteDepartment(c, pID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Position deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------SLA---------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateSLAHandler(c *gin.Context) {
	var s structs.Sla
	s.SlaName = c.PostForm("slaName")
	s.PriorityID, _ = strconv.Atoi(c.PostForm("priorityID"))
	s.SatisfactionID, _ = strconv.Atoi(c.PostForm("satisfactionID"))
	s.PolicyID, _ = strconv.Atoi(c.PostForm("policyID"))

	createdSla, e := CreateSla(c, s)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SLA"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "SLA created successfully",
		"createdSlaID": createdSla.SlaID,
	})
}

// UpdateSlaHandler updates an SLA
func UpdateSLAHandler(c *gin.Context) {
	var sla structs.Sla
	// Get SLA details from the request JSON or other sources
	sla.SlaID, _ = strconv.Atoi(c.PostForm("slaID"))
	sla.SlaName = c.PostForm("slaName")
	sla.PriorityID, _ = strconv.Atoi(c.PostForm("priorityID"))
	sla.SatisfactionID, _ = strconv.Atoi(c.PostForm("satisfactionID"))
	sla.PolicyID, _ = strconv.Atoi(c.PostForm("policyID"))

	// Call the UpdateSla function
	updatedSla, err := UpdateSla(c, sla.SlaID, sla.SlaName, sla.PriorityID, sla.SatisfactionID, sla.PolicyID)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":    "SLA updated successfully",
		"updatedSla": updatedSla,
	})
}

// GetSLAHandler retrieves details of a specific SLA by ID
func GetSLAHandler(c *gin.Context) {
	// Get SLA ID from the URL parameter
	slaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SLA ID"})
		return
	}

	// Call the GetSLA function to retrieve the SLA details
	sla, err := GetSla(c, slaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sla == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SLA not found"})
		return
	}

	// Return the SLA details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "SLA retrieved successfully",
		"sla":     sla,
	})
}

// ListSLAsHandler retrieves a list of all SLAs
func ListSLAHandler(c *gin.Context) {
	// Call the ListSLAs function to retrieve the list of SLAs
	slas, err := ListSla(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of SLAs in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of SLAs retrieved successfully",
		"slas":    slas,
	})
}

// DeleteSLAHandler deletes a specific dept by ID
func DeleteSLAHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	slaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteSla(c, slaID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Sla deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!--------------------------------------------------------SATISFACTION---------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateSatisfactionHandler(c *gin.Context) {
	var sat structs.Satisfaction
	sat.Name = c.PostForm("satisfactionName")
	sat.Rank, _ = strconv.Atoi(c.PostForm("rank"))
	sat.Emoji = c.PostForm("emoji")

	createdSatisfaction, err := CreateSatisfaction(c, sat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":             "Satisfaction created successfully",
		"createdSatisfaction": createdSatisfaction,
	})
}

// UpdateSatisfactionHandler updates satisfaction details
func UpdateSatisfactionHandler(c *gin.Context) {
	var satisfaction structs.Satisfaction
	// Get satisfaction details from the request JSON or other sources
	satisfaction.SatisfactionID, _ = strconv.Atoi(c.PostForm("satisfactionID"))
	satisfaction.Name = c.PostForm("satisfactionName")
	satisfaction.Rank, _ = strconv.Atoi(c.PostForm("rank"))
	satisfaction.Emoji = c.PostForm("emoji")

	// Call the UpdateSatisfaction function
	updatedSatisfaction, err := UpdateSatisfaction(c, satisfaction.SatisfactionID, satisfaction.Name, satisfaction.Rank, satisfaction.Emoji)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":             "Satisfaction updated successfully",
		"updatedSatisfaction": updatedSatisfaction,
	})
}

// GetSatisfactionHandler retrieves details of a specific satisfaction by ID
func GetSatisfactionHandler(c *gin.Context) {
	// Get satisfaction ID from the URL parameter
	satisfactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid satisfaction ID"})
		return
	}

	// Call the GetSatisfaction function to retrieve the satisfaction details
	satisfaction, err := GetSatisfaction(c, satisfactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if satisfaction == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satisfaction not found"})
		return
	}

	// Return the satisfaction details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Satisfaction retrieved successfully",
		"satisfaction": satisfaction,
	})
}

// ListSatisfactionsHandler retrieves a list of all satisfactions
func ListSatisfactionsHandler(c *gin.Context) {
	// Call the ListSatisfactions function to retrieve the list of satisfactions
	satisfactions, err := ListSatisfaction(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of satisfactions in the response
	c.JSON(http.StatusOK, gin.H{
		"message":       "List of satisfactions retrieved successfully",
		"satisfactions": satisfactions,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeleteSatisfactionHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	satID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid satisfaction ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteSatisfaction(c, satID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "satisfaction deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------POLICIES-----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreatePolicyHandler(c *gin.Context) {
	var pol structs.Policies
	pol.PolicyName = c.PostForm("policyName")
	pol.EmbeddedLink = c.PostForm("embeddedLink")
	pol.PolicyUrl = c.PostForm("policyUrl")

	createdPolicy, err := CreatePolicy(c, pol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Policy created successfully",
		"createdPolicies": createdPolicy,
	})
}

// UpdatePoliciesHandler updates policy details
func UpdatePolicyHandler(c *gin.Context) {
	var policies structs.Policies
	// Get policy details from the request JSON or other sources
	policies.PolicyID, _ = strconv.Atoi(c.PostForm("policyID"))
	policies.PolicyName = c.PostForm("policyName")
	policies.EmbeddedLink = c.PostForm("embeddedLink")
	policies.PolicyUrl = c.PostForm("policyUrl")

	// Call the UpdatePolicies function
	updatedPolicy, err := UpdatePolicy(c, policies.PolicyID, policies.PolicyName, policies.EmbeddedLink, policies.PolicyUrl)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":       "Policy updated successfully",
		"updatedPolicy": updatedPolicy,
	})
}

// GetPoliciesHandler retrieves details of a specific policy by ID
func GetPolicyHandler(c *gin.Context) {
	// Get policy ID from the URL parameter
	policyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	// Call the GetPolicies function to retrieve the policy details
	policy, err := GetPolicy(c, policyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if policy == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return
	}

	// Return the policy details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Policy retrieved successfully",
		"policy":  policy,
	})
}

// ListPoliciesHandler retrieves a list of all policies
func ListPoliciesHandler(c *gin.Context) {
	// Call the ListPolicies function to retrieve the list of policies
	policies, err := ListPolicies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of policies in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "List of policies retrieved successfully",
		"policies": policies,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeletePolicyHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeletePolicy(c, pID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Policy deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!------------------------------------------------------------UNIT-------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateUnitHandler(c *gin.Context) {
	var unit structs.Unit
	unit.UnitName = c.PostForm("unitName")
	unit.Emoji = c.PostForm("emoji")

	createdUnit, err := CreateUnit(c, unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Unit created successfully",
		"createdUnit": createdUnit,
	})
}

// UpdateUnitHandler updates unit details
func UpdateUnitHandler(c *gin.Context) {
	var unit structs.Unit
	// Get unit details from the request JSON or other sources
	unit.UnitID, _ = strconv.Atoi(c.PostForm("unitID"))
	unit.UnitName = c.PostForm("unitName")
	unit.Emoji = c.PostForm("emoji")

	// Call the UpdateUnit function
	updatedUnit, err := UpdateUnit(c, unit.UnitID, unit.UnitName, unit.Emoji)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Unit updated successfully",
		"updatedUnit": updatedUnit,
	})
}

// GetUnitHandler retrieves details of a specific unit by ID
func GetUnitHandler(c *gin.Context) {
	// Get unit ID from the URL parameter
	unitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	// Call the GetUnit function to retrieve the unit details
	unit, err := GetUnit(c, unitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if unit == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}

	// Return the unit details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Unit retrieved successfully",
		"unit":    unit,
	})
}

// ListUnitsHandler retrieves a list of all units
func ListUnitsHandler(c *gin.Context) {
	// Call the ListUnits function to retrieve the list of units
	units, err := ListUnit(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of units in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of units retrieved successfully",
		"units":   units,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeleteUnitHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	uID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteUnit(c, uID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Unit deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------ROLE--------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateRoleHandler(c *gin.Context) {
	var r structs.Role
	r.RoleName = c.PostForm("roleName")

	createdRole, err := CreateRole(c, r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Role created successfully",
		"createdRole": createdRole,
	})
}

// UpdateRoleHandler updates role details
func UpdateRoleHandler(c *gin.Context) {
	var role structs.Role
	// Get role details from the request JSON or other sources
	role.RoleID, _ = strconv.Atoi(c.PostForm("roleID"))
	role.RoleName = c.PostForm("roleName")

	// Call the UpdateRole function
	updatedRole, err := UpdateRole(c, role.RoleID, role.RoleName)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Role updated successfully",
		"updatedRole": updatedRole,
	})
}

// GetRoleHandler retrieves details of a specific role by ID
func GetRoleHandler(c *gin.Context) {
	// Get role ID from the URL parameter
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	// Call the GetRole function to retrieve the role details
	role, err := GetRole(c, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if role == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	// Return the role details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Role retrieved successfully",
		"role":    role,
	})
}

// ListRolesHandler retrieves a list of all roles
func ListRolesHandler(c *gin.Context) {
	// Call the ListRoles function to retrieve the list of roles
	roles, err := ListRoles(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of roles in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "List of roles retrieved successfully",
		"roles":   roles,
	})
}

// DeleteDeptHandler deletes a specific dept by ID
func DeleteRoleHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	rID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	// Call the DeleteAsset function to delete the dept
	s, err := DeleteRole(c, rID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------CATEGORY-----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateCategoryHandler(c *gin.Context) {
	var cat structs.Category
	cat.CategoryName = c.PostForm("categoryName")

	createdCategory, err := CreateCategory(c, cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Category created successfully",
		"createdCategory": createdCategory,
	})
}

// UpdateCategoryHandler updates category details
func UpdateCategoryHandler(c *gin.Context) {
	var category structs.Category
	// Get category details from the request JSON or other sources
	category.CategoryID, _ = strconv.Atoi(c.PostForm("categoryID"))
	category.CategoryName = c.PostForm("categoryName")

	// Call the UpdateCategory function
	updatedCategory, err := UpdateCategory(c, category)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":         "Category updated successfully",
		"updatedCategory": updatedCategory,
	})
}

// GetCategoryHandler retrieves details of a specific category by ID
func GetCategoryHandler(c *gin.Context) {
	// Get category ID from the URL parameter
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Call the GetCategory function to retrieve the category details
	category, err := GetCategory(c, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Return the category details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Category retrieved successfully",
		"category": category,
	})
}

// ListCategoriesHandler retrieves a list of all categories
func ListCategoriesHandler(c *gin.Context) {
	// Call the ListCategories function to retrieve the list of categories
	categories, err := ListCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of categories in the response
	c.JSON(http.StatusOK, gin.H{
		"message":    "List of categories retrieved successfully",
		"categories": categories,
	})
}

// DeleteCategoryHandler deletes a specific dept by ID
func DeleteCategoryHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	cID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Call the DeleteCategory function to delete the dept
	s, err := DeleteCategory(c, cID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!---------------------------------------------------------SUBCATEGORY---------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateSubCategoryHandler(c *gin.Context) {
	var subCat structs.SubCategory
	subCat.SubCategoryName = c.PostForm("subCategoryName")
	subCat.CategoryID, _ = strconv.Atoi(c.PostForm("categoryID"))

	createdSubCategory, err := CreateSubCategory(c, subCat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Sub-category created successfully",
		"createdSubCategory": createdSubCategory,
	})
}

// UpdateSubCategoryHandler updates sub-category details
func UpdateSubCategoryHandler(c *gin.Context) {
	var subCategory structs.SubCategory
	// Get sub-category details from the request JSON or other sources
	subCategory.SubCategoryID, _ = strconv.Atoi(c.PostForm("subCategoryID"))
	subCategory.SubCategoryName = c.PostForm("subCategoryName")
	subCategory.CategoryID, _ = strconv.Atoi(c.PostForm("categoryID"))

	// Call the UpdateSubCategory function
	updatedSubCategory, err := UpdateSubCategory(c, subCategory)
	if err != nil {
		// Handle the error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":            "Sub-category updated successfully",
		"updatedSubCategory": updatedSubCategory,
	})
}

// GetSubCategoryHandler retrieves details of a specific subcategory by ID
func GetSubCategoryHandler(c *gin.Context) {
	// Get subcategory ID from the URL parameter
	subcategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subcategory ID"})
		return
	}

	// Call the GetSubCategory function to retrieve the subcategory details
	subcategory, err := GetSubCategory(c, subcategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if subcategory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	// Return the subcategory details in the response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Subcategory retrieved successfully",
		"subcategory": subcategory,
	})
}

// ListSubCategoriesHandler retrieves a list of all subcategories
func ListSubCategoriesHandler(c *gin.Context) {
	// Call the ListSubCategories function to retrieve the list of subcategories
	subcategories, err := ListSubCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of subcategories in the response
	c.JSON(http.StatusOK, gin.H{
		"message":       "List of sub-categories retrieved successfully",
		"subcategories": subcategories,
	})
}

// DeleteCategoryHandler deletes a specific dept by ID
func DeleteSubCategoryHandler(c *gin.Context) {
	// Get asset ID from the URL parameter
	scID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sub-category ID"})
		return
	}

	// Call the DeleteSubCategory function to delete the dept
	s, err := DeleteSubCategory(c, scID)
	if err != nil || !s {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "SubCategory deleted successfully",
	})
}
