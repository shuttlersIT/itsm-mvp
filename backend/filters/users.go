package filters

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get a Staff by First Name
func GetStaffByFirstName(c *gin.Context, fname string) (int, structs.Staff) {
	var s structs.Staff
	first_name := fname

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE first_name = ?", first_name).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get a Staff by Last Name
func GetStaffByLastName(c *gin.Context, lname string) (int, structs.Staff) {
	var s structs.Staff
	last_name := lname

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE last_name = ?", last_name).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get a Staff by Email
func GetStaffByEmail(c *gin.Context, e string) (int, structs.Staff) {
	var s structs.Staff
	email := e

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE email = ?", email).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get a Staff by Username
func GetStaffByUsername(c *gin.Context, u int) (int, structs.Staff) {
	var s structs.Staff
	username := u

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, s
	}
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE username = ?", username).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, s
	}
	c.JSON(http.StatusOK, s)
	return 1, s
}

// Get a Staff by Position
func GetStaffByPosition(c *gin.Context, p int) (int, []structs.Staff) {
	var staff []structs.Staff
	position := p

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, staff
	}
	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE position_id = ?", position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, staff
	}
	defer rows.Close()

	for rows.Next() {
		var s structs.Staff
		if err := rows.Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, staff
		}
		staff = append(staff, s)
	}

	c.JSON(http.StatusOK, staff)
	return 1, staff
}

// Get a Staff by Department
func GetStaffByDepartment(c *gin.Context, dept int) (int, []structs.Staff) {
	var staff []structs.Staff
	department := dept

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, staff
	}
	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username, position_id, department_id FROM staff WHERE department_id = ?", department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, staff
	}
	defer rows.Close()

	for rows.Next() {
		var s structs.Staff
		if err := rows.Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 0, staff
		}
		staff = append(staff, s)
	}

	c.JSON(http.StatusOK, staff)
	return 1, staff
}
