package frontendstructs

import (
	"time"
)

type FrontendStaff struct {
	StaffID    int       `json:"staffID"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	StaffEmail string    `json:"staffEmail" binding:"required,email"`
	Username   string    `json:"userName"`
	Phone      string    `json:"phoneNumber" binding:"required,e164"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type FrontendAgent struct {
	AgentID    int       `json:"agentID"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	AgentEmail string    `json:"agentEmail" binding:"required,email"`
	Username   string    `json:"userName"`
	Phone      string    `json:"phoneNumber" binding:"required,e164"`
	Role       string    `json:"role"`
	Unit       string    `json:"unit"`
	Supervisor string    `json:"supervisorID"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type FrontendTicket struct {
	ID            int       `json:"ticketID"`
	Subject       string    `json:"subjectText"`
	Description   string    `json:"descriptionText"`
	Category      string    `json:"categoryName"`
	SubCategory   string    `json:"subCategoryName"`
	Priority      string    `json:"priorityName"`
	SLA           string    `json:"slaName"`
	Staff         string    `json:"staff"`
	Agent         string    `json:"agent"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DueAt         time.Time `json:"dueAt"`
	Asset         string    `json:"assetName"`
	RelatedTicket int       `json:"relatedTicketID"`
	Tag           string    `json:"tagName"`
	Site          string    `json:"siteName"`
	Status        string    `json:"statusText"`
	AttachmentID  int       `json:"attachmentID"`
}

type FrontendAsset struct {
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
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
