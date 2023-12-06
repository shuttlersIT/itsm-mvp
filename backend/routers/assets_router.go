package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupAssetsRouter sets up the assets router
func SetupAssetsRouter(assets *gin.RouterGroup, db *sql.DB) {
	assets.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	assets.GET("/assets/:id", handlers.GetAgentHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	assets.POST("/assets", handlers.CreateAssetHandler)

	// Example route for PUT request
	assets.PUT("/assets/:id", handlers.UpdateAssetsHandler)

	// Example route for DELETE request
	assets.DELETE("/assets/:id", handlers.DeleteAssetHandler)

	// Include other routes and handlers as needed
}
