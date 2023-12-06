package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get positions from database
func getPosition(c *gin.Context, pid int) string {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return ""
	}

	//session := sessions.Default(c)
	id := pid
	var s structs.Position
	err := db.QueryRow("SELECT id, position_name, emoji FROM positions WHERE id = ?", id).
		Scan(&s.PositionID, &s.PositionName, &s.CadreName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return ""
	}
	c.JSON(http.StatusOK, s)
	return s.PositionName
}

// Get a department from database
func getDepartment(c *gin.Context, did int) string {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return ""
	}

	//session := sessions.Default(c)
	id := did
	var s structs.Department
	err := db.QueryRow("SELECT id, department_name, emoji FROM departments WHERE id = ?", id).
		Scan(&s.DepartmentID, &s.DepartmentName, &s.Emoji)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return ""
	}
	c.JSON(http.StatusOK, s)
	return s.DepartmentName
}

/*
// Get a department from database
func getRole(c *gin.Context, rid int) string {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return ""
	}

	//session := sessions.Default(c)
	id := rid
	var s structs.Role
	err := db.QueryRow("SELECT id, role_name FROM departments WHERE id = ?", id).
		Scan(&s.RoleID, &s.RoleName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return ""
	}
	c.JSON(http.StatusOK, s)
	return s.RoleName
}
*/
