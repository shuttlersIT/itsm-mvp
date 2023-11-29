package structs

import (
	"time"
)

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

type Ticket struct {
	ID              int       `json:"ticket_id"`
	Subject         string    `json:"subject"`
	Description     string    `json:"description"`
	Category        int       `json:"category"`
	SubCategory     int       `json:"sub_category"`
	Priority        int       `json:"priority"`
	SLA             int       `json:"sla"`
	StaffID         int       `json:"staff_id"`
	AgentID         int       `json:"agent_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DueAt           time.Time `json:"due_at"`
	AssetID         int       `json:"asset_id"`
	RelatedTicketID int       `json:"related_ticket_id"`
	Tag             string    `json:"tag"`
	Site            string    `json:"site"`
	Status          string    `json:"status"`
	AttachmentID    int       `json:"attachment"`
}

type Asset struct {
	ID            int       `json:"id"`
	AssetID       string    `json:"asset_id"`
	AssetType     string    `json:"asset_type"`
	AssetName     string    `json:"asset_name"`
	Description   string    `json:"description"`
	Manufacturer  string    `json:"manufacturer"`
	Model         string    `json:"model"`
	SerialNumber  string    `json:"serial_number"`
	PurchaseDate  string    `json:"purchase_date"`
	PurchasePrice string    `json:"purchase_price"`
	Vendor        string    `json:"vendor"`
	Site          string    `json:"site"`
	Status        string    `json:"status"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Sla struct {
	SlaID          int       `json:"sla_id"`
	SlaName        string    `json:"sla_name"`
	PriorityID     int       `json:"priority_id"`
	SatisfactionID int       `json:"satisfaction_id"`
	PolicyID       int       `json:"policy_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Priority struct {
	PriorityID    int       `json:"priority_id"`
	Name          string    `json:"priority_name"`
	FirstResponse int       `json:"first_response"`
	Colour        string    `json:"red"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Satisfaction struct {
	SatisfactionID int       `json:"satisfaction_id"`
	Name           string    `json:"satisfaction_name"`
	Rank           int       `json:"rank"`
	Emoji          string    `json:"emoji"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Policies struct {
	PolicyID     int       `json:"policy_id"`
	PolicyName   string    `json:"policy_name"`
	EmbeddedLink string    `json:"policy_embed"`
	PolicyUrl    string    `json:"policy_url"`
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

type Category struct {
	CategoryID   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SubCategory struct {
	SubCategoryID   int       `json:"sub_category_id"`
	SubCategoryName string    `json:"sub_category_name"`
	CategoryID      int       `json:"category_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Status struct {
	StatusID   int       `json:"status_id"`
	StatusName string    `json:"status_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

type AssetType struct {
	AssetTypeID int       `json:"asset_type_id"`
	AssetType   string    `json:"asset_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssetAssignment struct {
	AssignmentID   int       `json:"assignment_id"`
	AssetID        int       `json:"asset_id"`
	StaffID        int       `json:"staff_id"`
	AssignedBy     int       `json:"assigned_by"`
	AssignmentType string    `json:"assignment_type"`
	DueAt          time.Time `json:"due_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
