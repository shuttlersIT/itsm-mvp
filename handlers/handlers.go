package handlers

import (
	"net/http"

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
	a.AssetName = c.PostForm("assetName")
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
/*!-----------------------------------------------------------PRIORITY----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreatePriority function Usage
func CreatePriorityHandler(c *gin.Context) {
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
/*!----------------------------------------------------------POSITIONS----------------------------------------------------------------!*/
/*!-----------------------------------------------------------------------------------------------------------------------------------!*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreatePosition function Usage
func CreatePositionHandler(c *gin.Context) {
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
