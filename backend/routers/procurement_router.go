package routers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupProcurementRouter sets up the procurement router
func SetupProcurementRouter(procurement *gin.RouterGroup, db *sql.DB) {
	procurement.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	procurement.GET("/procurement/work/:id", handlers.ProcurementHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	procurement.POST("/procurement/work", handlers.ProcurementHandler)

	// Example route for PUT request
	procurement.PUT("/procurement/work/:id", func(c *gin.Context) {
		status := false
		id := c.Param("id")

		// Your logic here to handle the update of a procurement item with the specified ID
		if !status {
			id = ""
			fmt.Errorf("update operation failed")
		}
		fmt.Println(id)

	})

	// Example route for DELETE request
	procurement.DELETE("/procurement/work/:id", func(c *gin.Context) {
		id := c.Param("id")
		status := false
		// Your logic here to handle the update of a procurement item with the specified ID
		if !status {
			id = ""
			fmt.Errorf("update operation failed")
		}
		fmt.Println(id)
	})

	// Include other routes and handlers as needed
}
