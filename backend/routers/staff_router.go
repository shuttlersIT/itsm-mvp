package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupStaffRouter sets up the staff router
func SetupStaffRouter(staff *gin.RouterGroup, db *sql.DB) {
	staff.Use(middleware.AuthorizeRequest())

	// Example route with dynamic parameter
	staff.GET("/staff/:id", handlers.GetStaffHandler)

	// Add other routes and handlers as needed

	// Example route for POST request
	staff.POST("/staff", handlers.CreateStaffHandler)

	// Example route for PUT request
	staff.PUT("/staff/:id", handlers.UpdateStaffHandler)

	// Example route for DELETE request
	staff.DELETE("/staff/:id", handlers.DeleteStaffHandler)

	// Routes for Department operations
	staff.POST("/department", handlers.CreateDepartmentHandler)
	staff.GET("/department/:id", handlers.GetDepartmentHandler)
	staff.PUT("/department/:id", handlers.UpdateDepartmentHandler)
	staff.DELETE("/department/:id", handlers.DeleteDepartmentHandler)

	// Routes for Position operations
	staff.POST("/position", handlers.CreatePositionHandler)
	staff.GET("/position/:id", handlers.GetPositionHandler)
	staff.PUT("/position/:id", handlers.UpdatePositionHandler)
	staff.DELETE("/position/:id", handlers.DeletePositionHandler)

	// Routes for Username operations
	staff.POST("/username", handlers.RegisterHandler)
	staff.GET("/username/:id", handlers.PasswordLoginHandler)
	staff.PUT("/username/:id", handlers.UpdateUserNameHandler)
	staff.DELETE("/username/:id", handlers.ChangeUserPasswordHandler)
}
