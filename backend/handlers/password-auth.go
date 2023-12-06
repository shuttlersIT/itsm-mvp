package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
	"golang.org/x/crypto/bcrypt"
)

/*
func hash(pwd string) []byte {
	return []byte(pwd)
}
*/

// Get Position by Name
func getPositionByName(c *gin.Context, positionName string) (*structs.Position, int, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get position by name handler"})
		return nil, 0, fmt.Errorf("unable to reach DB")
	}

	var position structs.Position
	err := db.QueryRow("SELECT id, position_name FROM positions WHERE position_name = ?", positionName).
		Scan(&position.PositionID, &position.PositionName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return nil, 0, fmt.Errorf("position not found")
	}

	return &position, position.PositionID, nil
}

// Get Department by Name
func getDepartmentByName(c *gin.Context, departmentName string) (*structs.Department, int, error) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get department by name handler"})
		return nil, 0, fmt.Errorf("position not found")
	}

	var department structs.Department
	err := db.QueryRow("SELECT id, department_name FROM departments WHERE department_name = ?", departmentName).
		Scan(&department.DepartmentID, &department.DepartmentName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return nil, 0, fmt.Errorf("department not found")
	}

	return &department, department.DepartmentID, nil
}

// Get a user ID from database
func getCred(c *gin.Context, username string) (int, string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, "Error"
	}
	var s structs.StaffLoginCredentials
	err := db.QueryRow("SELECT id, username, password FROM staff_credentials WHERE username = ?", username).
		Scan(&s.CredentialID, s.Username, s.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, "Error"
	}
	staffCredentials := s
	c.JSON(http.StatusOK, staffCredentials)

	return staffCredentials.CredentialID, staffCredentials.Password
}

// Get a username from database
func GetCred2(c *gin.Context, username int) (*structs.StaffLoginCredentials, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to read db")
	}

	var s structs.StaffLoginCredentials
	err := db.QueryRow("SELECT id, username, password FROM staff_credentials WHERE id = ?", username).
		Scan(&s.CredentialID, s.Username, s.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return nil, fmt.Errorf("user credentials not found")
	}
	staffCredentials := s
	c.JSON(http.StatusOK, staffCredentials)

	return &staffCredentials, nil
}

// Get a user ID from database
func GetStaff(c *gin.Context, id int) (int, string, string, string, int, int, int) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, "", "", "", 0, 0, 0
	}
	var s structs.Staff
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username_id, position_id, department_id FROM staff WHERE username_id = ?", id).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0, "", "", "", 0, 0, 0
	}
	c.JSON(http.StatusOK, s)
	return s.StaffID, s.FirstName, s.LastName, s.StaffEmail, s.Username, s.PositionID, s.DepartmentID
}

// Register with Username & Password
func RegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	staff_email := c.PostForm("email")
	position := c.PostForm("position")
	department := c.PostForm("department")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}
	hash := string(hashedPassword)
	_, positionID, _ := getPositionByName(c, position)
	_, departmentID, _ := getDepartmentByName(c, department)
	usernameID := addStaff(c, username, hash)

	// Add User to DB
	staffID := createStaffByForm(c, first_name, last_name, staff_email, usernameID, positionID, departmentID)

	var staff structs.Staff
	staff.StaffID = staffID
	staff.FirstName = first_name
	staff.LastName = last_name
	staff.StaffEmail = staff_email
	staff.Username = usernameID
	staff.PositionID = positionID
	staff.DepartmentID = departmentID

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login with Username & Password
func PasswordLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// Retrieve the hashed password from the user store (replace with database query)
	id, hashedPassword := getCred(c, username)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	var staff structs.Staff

	staff.StaffID, staff.FirstName, staff.LastName, staff.StaffEmail, staff.Username, staff.PositionID, staff.DepartmentID = GetStaff(c, id)
	// Authentication successful; you can generate a token or set a session here
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"message": "Invalid session state."})
		return
	}
	session.Set("id", staff.StaffID)
	session.Set("user-email", staff.StaffEmail)
	session.Set("user-firstName", staff.FirstName)
	session.Set("user-lastName", staff.LastName)
	session.Set("user-position", getPosition(c, staff.PositionID))
	session.Set("user-department", getDepartment(c, staff.DepartmentID))
}

// Get an agent id from database
func GetStaffIdHandler(c *gin.Context) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	adminSession := sessions.Default(c)
	email := adminSession.Get("user-email")
	//userID := adminSession.Get("user-id")
	var a structs.Agent
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, username, role_id, unit, supervisor_id FROM agents WHERE email = ?", email).
		Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0
	}
	adminSession.Set("id", a.AgentID)
	agent := a
	id := agent.AgentID
	//c.JSON(http.StatusOK, agent)

	return id
}

// Update Username
func UpdateUserNameHandler(c *gin.Context) {
	var positionError error
	var departmentError error
	//var positionDB *structs.Position
	//var departmentDB *structs.Department

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	username := c.PostForm("username")
	//password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	staff_email := c.PostForm("email")
	position := c.PostForm("position")
	department := c.PostForm("department")

	//session := sessions.Default(c)
	//id := session.Get("user-id")

	//userID := GetUserId(c) // Implement a function to get the current user's ID
	//newName := c.PostForm("name")
	//var staff structs.Staff
	//if err := c.ShouldBindJSON(&staff); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	id := getUserByEmail(c, staff_email)
	staff, _ := GetUserByID(c, id)

	staff.StaffID = getUserByEmail(c, staff_email)
	staff.FirstName = first_name
	staff.LastName = last_name
	staff.StaffEmail = staff_email
	staff.Username = getUserByEmail(c, staff.StaffEmail)
	_, staff.PositionID, positionError = getPositionByName(c, position)
	if positionError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": positionError.Error()})
	}
	_, staff.DepartmentID, departmentError = getDepartmentByName(c, department)
	if departmentError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": departmentError.Error()})
	}

	_, err := db.Exec("UPDATE staff_credentials SET username = ? WHERE id = ?", username, staff.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, error := db.Exec("UPDATE staff SET first_name = ?, last_name = ?, staff_email = ?, username = ?, position_id = ?, department_id = ?, WHERE id = ?", staff.FirstName, staff.LastName, staff.StaffEmail, staff.Username, staff.PositionID, staff.DepartmentID, id)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User updated successfully")

	c.JSON(http.StatusOK, gin.H{"message": "Name updated successfully"})
}

// Update Password
func ChangeUserPasswordHandler(c *gin.Context) {
	username := c.PostForm("unitName")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	var s structs.StaffLoginCredentials
	err := db.QueryRow("SELECT password FROM staff_credentials WHERE username = ?", username).
		Scan(&s.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	//c.JSON(http.StatusOK, s)

	dbPassword := s.Password
	currentPassword := c.PostForm("current_password")
	newPassword := c.PostForm("new_password")
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// Compare the hashed password with the provided password
	error := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(dbPassword))
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// New Password :: Hash to String
	hash := string(hashedPassword)

	_, err2 := db.Exec("UPDATE staff_credentials SET password = ?, WHERE username = ?", hash, username)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// Register a New User and Credentials
func addStaff(c *gin.Context, u string, hp string) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	var t structs.StaffLoginCredentials
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	result, err := db.Exec("INSERT INTO staff_credentials (username, password) VALUES (?, ?)", u, hp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	t.CredentialID = int(lastInsertID)

	return t.CredentialID
}

// Register a New User and Info
func AddStaffInfo(c *gin.Context, fn string, ln string, se string, u int, p int, d int) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	var t structs.Staff
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	result, err := db.Exec("INSERT INTO staff (first_name, last_name, staff_email, username_id, position_id, department_id) VALUES (?, ?, ?, ?, ?, ?)", fn, ln, se, u, p, d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 0
	}

	lastInsertID, _ := result.LastInsertId()
	t.StaffID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
	return t.StaffID
}

// Get an agent id from database
func GetAgentIdHandler(c *gin.Context) int {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0
	}

	adminSession := sessions.Default(c)
	email := adminSession.Get("user-email")
	//userID := adminSession.Get("user-id")
	var a structs.Agent
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, username, role_id, unit, supervisor_id FROM agents WHERE email = ?", email).
		Scan(&a.AgentID, &a.FirstName, &a.LastName, &a.AgentEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return 0
	}
	adminSession.Set("id", a.AgentID)
	agent := a
	id := agent.AgentID
	//c.JSON(http.StatusOK, agent)

	return id
}
