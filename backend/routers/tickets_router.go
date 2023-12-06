package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupTicketsRouter sets up the tickets router
func SetupTicketsRouter(tickets *gin.RouterGroup, db *sql.DB) {
	tickets.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	tickets.GET("/tickets/:id", handlers.GetTicketHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	tickets.POST("/tickets", handlers.CreateTicketHandler)

	// Example route for PUT request
	tickets.PUT("/tickets/:id", handlers.UpdateTicketHandler)

	// Example route for DELETE request
	tickets.DELETE("/tickets/:id", handlers.DeleteTicketHandler)

	// Routes for Status operations
	tickets.POST("/status", handlers.CreateStatusHandler)
	tickets.GET("/status/:id", handlers.GetStatusHandler)
	tickets.PUT("/status/:id", handlers.UpdateStatusHandler)
	tickets.DELETE("/status/:id", handlers.DeleteStatusHandler)

	// Routes for SLA operations
	tickets.POST("/sla", handlers.CreateSLAHandler)
	tickets.GET("/sla/:id", handlers.GetSLAHandler)
	tickets.PUT("/sla/:id", handlers.UpdateSLAHandler)
	tickets.DELETE("/sla/:id", handlers.DeleteSLAHandler)

	// Routes for Priority operations
	tickets.POST("/priority", handlers.CreatePriorityHandler)
	tickets.GET("/priority/:id", handlers.GetPriorityHandler)
	tickets.PUT("/priority/:id", handlers.UpdatePriorityHandler)
	tickets.DELETE("/priority/:id", handlers.DeletePriorityHandler)

	// Routes for Category operations
	tickets.POST("/category", handlers.CreateCategoryHandler)
	tickets.GET("/category/:id", handlers.GetCategoryHandler)
	tickets.PUT("/category/:id", handlers.UpdateCategoryHandler)
	tickets.DELETE("/category/:id", handlers.DeleteCategoryHandler)

	// Routes for Sub-Category operations
	tickets.POST("/subcategory", handlers.CreateSubCategoryHandler)
	tickets.GET("/subcategory/:id", handlers.GetSubCategoryHandler)
	tickets.PUT("/subcategory/:id", handlers.UpdateSubCategoryHandler)
	tickets.DELETE("/subcategory/:id", handlers.DeleteSubCategoryHandler)

	// Routes for Satisfaction operations
	tickets.POST("/satisfaction", handlers.CreateSatisfactionHandler)
	tickets.GET("/satisfaction/:id", handlers.GetSatisfactionHandler)
	tickets.PUT("/satisfaction/:id", handlers.UpdateSatisfactionHandler)
	tickets.DELETE("/satisfaction/:id", handlers.DeleteSatisfactionHandler)

	// Routes for Policies operations
	tickets.POST("/policies", handlers.CreatePolicyHandler)
	tickets.GET("/policies/:id", handlers.GetPolicyHandler)
	tickets.PUT("/policies/:id", handlers.UpdatePolicyHandler)
	tickets.DELETE("/policies/:id", handlers.DeletePolicyHandler)

}
