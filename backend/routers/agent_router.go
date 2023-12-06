package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupAgentRouter sets up the staff router
func SetupAgentRouter(agent *gin.RouterGroup, db *sql.DB) {
	agent.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	agent.GET("/agent/:id", handlers.GetAgentHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	agent.POST("/agent", handlers.CreateAgentHandler)

	// Example route for PUT request
	agent.PUT("/agent/:id", handlers.UpdateAgentHandler)

	// Example route for DELETE request
	agent.DELETE("/agent/:id", handlers.DeleteAgentHandler)

	// Include other routes and handlers as needed

	// Example route with dynamic parameter
	agent.GET("/agent/:id", handlers.GetAgentHandler)

	// Routes for Username operations
	//agent.POST("/agent/username", handlers.RegisterAgentHandler)
	//agent.GET("/agent/username/:username", handlers.GetAgentUsernameHandler)
	//agent.PUT("/agent/username/:id", handlers.UpdateAgentUsernameHandler)
	//agent.DELETE("/agent/username/:id", handlers.DeleteAgentUsernameHandler)

	// Routes for Role operations
	agent.POST("/agent/role", handlers.CreateRoleHandler)
	agent.GET("/agent/role/:role", handlers.GetRoleHandler)
	agent.GET("/agent/role/:id", handlers.GetRoleHandler)
	agent.PUT("/agent/role/:id", handlers.UpdateRoleHandler)
	agent.DELETE("/agent/role/:id", handlers.DeleteRoleHandler)

	// Routes for Unit operations
	agent.POST("/agent/unit", handlers.CreateUnitHandler)
	agent.GET("/agent/unit/:unit", handlers.GetUnitHandler)
	agent.GET("/agent/unit/:id", handlers.GetUnitHandler)
	agent.PUT("/agent/unit/:id", handlers.UpdateUnitHandler)
	agent.DELETE("/agent/unit/:id", handlers.DeleteUnitHandler)

	// Additional routes for Username
	//agent.GET("/agent/username/:username", handlers.GetAgentByUsernameHandler)

	// Additional routes for Role
	//agent.GET("/agent/role/:role", handlers.GetAgentsByRoleHandler)

	// Additional routes for Unit
	//agent.GET("/agent/unit/:unit", handlers.GetAgentsByUnitHandler)
}
