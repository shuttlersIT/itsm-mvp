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

func CreateAgentHandler(c *gin.Context) {
	var a structs.Agent
	a.FirstName = c.PostForm("firstName")
	a.LastName = c.PostForm("lastName")
	a.AgentEmail = c.PostForm("agentEmail")
	a.Username, _ = strconv.Atoi(c.PostForm("username"))
	a.Phone, _ = strconv.Atoi(c.PostForm("phone"))
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
	s.Phone, _ = strconv.Atoi(c.PostForm("phone"))
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

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!-----------------------------------------------------------SLA---------------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateSlaHandler(c *gin.Context) {
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
		"createdSlaID": createdSla,
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

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*!___________________________________________________________________________________________________________________________________!*/
/*!----------------------------------------------------------POLICIES-----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreatePoliciesHandler(c *gin.Context) {
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
