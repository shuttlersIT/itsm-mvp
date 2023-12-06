package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupAgentRouter sets up the staff router
func SetupAdminRouter(admin *gin.RouterGroup, db *sql.DB) {
	admin.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	admin.GET("/agent/:id", handlers.GetAgentHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	admin.POST("/agent", handlers.CreateAgentHandler)

	// Example route for PUT request
	admin.PUT("/agent/:id", handlers.UpdateAgentHandler)

	// Example route for DELETE request
	admin.DELETE("/agent/:id", handlers.DeleteAgentHandler)

	// Include other routes and handlers as needed
}
