package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/scanners"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// Get a user ID from database
func GetUser(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	session := sessions.Default(c)
	id := session.Get("user-id")
	var s structs.Staff
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE id = ?", id).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Get a user ID from database
func GetUserByID(c *gin.Context, id int) (*structs.Staff, error) {
	//var s structs.Staff
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("db unreacheable")
	}

	rows, err := db.Query("SELECT * FROM staff WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "staff not found"})
		return nil, err
	}
	for rows.Next() {
		s, e := scanners.ScanIntoStaff(rows)
		if e != nil {
			c.JSON(http.StatusNotFound, "an error occured when getting staff from db")
			return nil, e
		} else {
			c.JSON(http.StatusOK, "staff retrieval successfull")
			return s, nil
		}
	}
	return nil, fmt.Errorf("staff %d not found", id)
}

// Get a user ID from database
func getUserByEmail(c *gin.Context, e string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	//session := sessions.Default(c)
	//id := session.Get("user-id")
	var s structs.Staff
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE email = ?", e).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0
	}
	//c.JSON(http.StatusOK, s)

	return s.StaffID
}

// Update a update by Struct
func UpdateUser(c *gin.Context, user structs.Staff) (int, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0, fmt.Errorf("unable to reach DB")
	}

	// session := sessions.Default(c)
	// id := session.Get("id")
	// var s structs.Staff
	id := user.StaffID
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, fmt.Errorf("unable to bind json")
	}
	_, err := db.Exec("UPDATE staff SET first_name = ?, last_name = ?, staff_email = ?, username_id = ?, position_id = ?, department_id = ?, WHERE id = ?", user.FirstName, user.LastName, user.StaffEmail, user.Username, user.PositionID, user.DepartmentID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0, fmt.Errorf("user not found")
	}
	c.JSON(http.StatusOK, "User updated successfully")

	return id, fmt.Errorf("successful")
}

// Update a update by Struct
func UpdateUsername(c *gin.Context, staffID int, username string, password string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	// session := sessions.Default(c)
	// id := session.Get("id")
	// var s structs.Staff
	id := staffID
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	_, err := db.Exec("UPDATE staff_credentials SET id = ?, username_id = ?, password = ? WHERE id = ?", id, username, password, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}
	c.JSON(http.StatusOK, "User updated successfully")

	return id
}

// Delete a ticket by ID
func DeleteUserOperation(c *gin.Context, tid int) (bool, error) {
	id := tid

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return false, fmt.Errorf("unable to reach DB")
	}
	_, err := db.Exec("DELETE FROM staff WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false, err
	}
	return true, nil
}

// Create staff
func CreateUser2(c *gin.Context, staff structs.Staff) (*structs.Staff, int, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return nil, 0, fmt.Errorf("db unreacheable")
	}

	//session := sessions.Default(c)
	//email := session.Get("user-email")
	//username := session.Get("user-name")
	//first_name := session.Get("user-firstName")
	//last_name := session.Get("user-lastName")
	//sub := session.Get("user-sub")

	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("invalid request")
	}

	result, err := db.Exec("INSERT INTO staff (first_name, last_name, staff_email, phone, username_id, position_id, department_id) VALUES (?, ?, ?, ?, ?, ?, ?)", staff.FirstName, staff.LastName, staff.StaffEmail, staff.Phone, staff.Username, staff.PositionID, staff.DepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("failed to create staff")
	}

	lastInsertID, _ := result.LastInsertId()
	st, e := GetUserByID(c, int(lastInsertID))
	if e != nil {
		c.JSON(http.StatusNotFound, "Staff creation failed")
		return nil, 0, e
	}
	c.JSON(http.StatusCreated, st)
	c.JSON(http.StatusOK, "Staff created successfully")
	return st, st.StaffID, nil

	/*
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, fmt.Errorf("Add Staff failed not found")
		}

		//lastInsertID, _ := result.LastInsertId()
		s = int(lastInsertID)
		c.JSON(http.StatusCreated, s)

		c.JSON(http.StatusOK, "User created successfully")

		return s.StaffID
	*/
}

// Create staff
func createStaffByForm(c *gin.Context, fn string, ln string, se string, u int, pid int, did int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return 0
	}

	var s structs.Staff
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	result, err := db.Exec("INSERT INTO staff (first_name, last_name, staff_email, username_id, position_id, department_id) VALUES (?, ?, ?, ?, ?, ?)", fn, ln, se, u, pid, did)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	s.StaffID = int(lastInsertID)
	staff := s
	c.JSON(http.StatusCreated, staff)

	c.JSON(http.StatusOK, "User created successfully")

	return s.StaffID
}

// List all Agents
func ListUsers(c *gin.Context) ([]*structs.Staff, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	rows, err := db.Query("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("unable to retrieve agents at query")
	}
	defer rows.Close()

	var staff []*structs.Staff
	for rows.Next() {
		//var a structs.Agent
		a, err := scanners.ScanIntoStaff(rows)
		//if err := rows.Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID); err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return nil, fmt.Errorf("unable to retrieve agents")
		//}
		if err != nil {
			return nil, fmt.Errorf("unable to add staff to array")
		}
		staff = append(staff, a)
	}
	c.JSON(http.StatusOK, "Staff Listed successfully")
	return staff, nil
}
