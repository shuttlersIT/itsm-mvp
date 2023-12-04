package structs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

type AccessToken struct {
	UserID int64 `json:"user_id"`
	// Add other claims as needed
	jwt.StandardClaims
}

type RefreshToken struct {
	UserID int64 `json:"user_id"`
	// Add other claims as needed
	jwt.StandardClaims
}

type StaffLoginCredentials struct {
	CredentialID int       `json:"credentials_id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	StaffID      int       `json:"staff_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AgentLoginCredentials struct {
	CredentialID int       `json:"credentials_id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	AgentID      int       `json:"agent_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Staff struct {
	StaffID      int       `json:"staff_id"`
	FirstName    string    `json:"first_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	StaffEmail   string    `json:"staff_email" binding:"required,email"`
	Username     int       `json:"username"`
	Phone        string    `json:"phoneNumber" binding:"required,e164"`
	PositionID   int       `json:"position_id"`
	DepartmentID int       `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Agent struct {
	AgentID      int       `json:"agent_id" binding:"required,email"`
	FirstName    string    `json:"first_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	AgentEmail   string    `json:"agent_email" binding:"required,email"`
	Username     int       `json:"username"`
	Phone        string    `json:"phoneNumber" binding:"required,e164"`
	RoleID       int       `json:"role_id"`
	Unit         int       `json:"unit"`
	SupervisorID int       `json:"supervisor_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Position struct {
	PositionID   int       `json:"position_id"`
	PositionName string    `json:"position_name"`
	CadreName    string    `json:"cadre_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Department struct {
	DepartmentID   int       `json:"department_id"`
	DepartmentName string    `json:"department_name"`
	Emoji          string    `json:"emoji"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Unit struct {
	UnitID    int       `json:"unit_id"`
	UnitName  string    `json:"unit_name"`
	Emoji     string    `json:"emoji"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	RoleID    int       `json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
