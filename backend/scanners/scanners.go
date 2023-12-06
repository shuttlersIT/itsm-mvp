package scanners

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

func ScanIntoStaff(rows *sql.Rows) (*structs.Staff, error) {
	staff := new(structs.Staff)
	err := rows.Scan(
		&staff.StaffID,
		&staff.FirstName,
		&staff.LastName,
		&staff.StaffEmail,
		&staff.Username,
		&staff.PositionID,
		&staff.DepartmentID,
		&staff.CreatedAt,
		&staff.UpdatedAt)
	return staff, err
}

func ScanIntoAgent(rows *sql.Rows) (*structs.Agent, error) {
	agent := new(structs.Agent)
	err := rows.Scan(
		&agent.AgentID,
		&agent.FirstName,
		&agent.LastName,
		&agent.AgentEmail,
		&agent.Username,
		&agent.RoleID,
		&agent.Unit,
		&agent.SupervisorID,
		&agent.CreatedAt,
		&agent.UpdatedAt)
	return agent, err
}

func ScanIntoTicket(rows *sql.Rows) (*structs.Ticket, error) {
	ticket := new(structs.Ticket)
	err := rows.Scan(
		&ticket.ID,
		&ticket.Subject,
		&ticket.Description,
		&ticket.Category,
		&ticket.SubCategory,
		&ticket.Priority,
		&ticket.SLA,
		&ticket.StaffID,
		&ticket.AgentID,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
		&ticket.DueAt,
		&ticket.AssetID,
		&ticket.RelatedTicketID,
		&ticket.Tag,
		&ticket.Site,
		&ticket.Status,
		&ticket.AttachmentID)

	return ticket, err
}

func ScanIntoAsset(rows *sql.Rows) (*structs.Asset, error) {
	asset := new(structs.Asset)
	err := rows.Scan(
		&asset.ID,
		&asset.AssetID,
		&asset.AssetType,
		&asset.AssetName,
		&asset.Description,
		&asset.Manufacturer,
		&asset.Model,
		&asset.SerialNumber,
		&asset.PurchaseDate,
		&asset.PurchasePrice,
		&asset.Vendor,
		&asset.Site,
		&asset.Status,
		&asset.CreatedBy,
		&asset.CreatedAt,
		&asset.UpdatedAt)
	return asset, err
}

func ScanIntoSla(rows *sql.Rows) (*structs.Sla, error) {
	sla := new(structs.Sla)
	err := rows.Scan(
		&sla.SlaID,
		&sla.SlaName,
		&sla.PriorityID,
		&sla.SatisfactionID,
		&sla.PolicyID,
		&sla.CreatedAt,
		&sla.UpdatedAt)
	return sla, err
}

func ScanIntoPriority(rows *sql.Rows) (*structs.Priority, error) {
	priority := new(structs.Priority)
	err := rows.Scan(
		&priority.PriorityID,
		&priority.Name,
		&priority.FirstResponse,
		&priority.Colour,
		&priority.CreatedAt,
		&priority.UpdatedAt)
	return priority, err
}

func ScanIntoSatisfaction(rows *sql.Rows) (*structs.Satisfaction, error) {
	satisfaction := new(structs.Satisfaction)
	err := rows.Scan(
		&satisfaction.SatisfactionID,
		&satisfaction.Name,
		&satisfaction.Rank,
		&satisfaction.Emoji,
		&satisfaction.CreatedAt,
		&satisfaction.UpdatedAt)
	return satisfaction, err
}

func ScanIntoPolicies(rows *sql.Rows) (*structs.Policies, error) {
	policy := new(structs.Policies)
	err := rows.Scan(
		&policy.PolicyID,
		&policy.PolicyName,
		&policy.EmbeddedLink,
		&policy.PolicyUrl,
		&policy.CreatedAt,
		&policy.UpdatedAt)
	return policy, err
}

func ScanIntoPosition(rows *sql.Rows) (*structs.Position, error) {
	position := new(structs.Position)
	err := rows.Scan(
		&position.PositionID,
		&position.PositionName,
		&position.CadreName,
		&position.CreatedAt,
		&position.UpdatedAt)
	return position, err
}

func ScanIntoDepartment(rows *sql.Rows) (*structs.Department, error) {
	department := new(structs.Department)
	err := rows.Scan(
		&department.DepartmentID,
		&department.DepartmentName,
		&department.Emoji,
		&department.CreatedAt,
		&department.UpdatedAt)
	return department, err
}

func ScanIntoUnit(rows *sql.Rows) (*structs.Unit, error) {
	unit := new(structs.Unit)
	err := rows.Scan(
		&unit.UnitID,
		&unit.UnitName,
		&unit.Emoji,
		&unit.CreatedAt,
		&unit.UpdatedAt)
	return unit, err
}

func ScanIntoRole(rows *sql.Rows) (*structs.Role, error) {
	role := new(structs.Role)
	err := rows.Scan(
		&role.RoleID,
		&role.RoleName,
		&role.CreatedAt,
		&role.UpdatedAt)
	return role, err
}

func ScanIntoCategories(rows *sql.Rows) (*structs.Category, error) {
	category := new(structs.Category)
	err := rows.Scan(
		&category.CategoryID,
		&category.CategoryName,
		&category.CreatedAt,
		&category.UpdatedAt)
	return category, err
}

func ScanIntoSubCategories(rows *sql.Rows) (*structs.SubCategory, error) {
	subCategory := new(structs.SubCategory)
	err := rows.Scan(
		&subCategory.SubCategoryID,
		&subCategory.SubCategoryName,
		&subCategory.CategoryID,
		&subCategory.CreatedAt,
		&subCategory.UpdatedAt)
	return subCategory, err
}

func ScanIntoStatus(rows *sql.Rows) (*structs.Status, error) {
	status := new(structs.Status)
	err := rows.Scan(
		&status.StatusID,
		&status.StatusName,
		&status.CreatedAt,
		&status.UpdatedAt)
	return status, err
}

func ScanIntoStaffLoginCredentials(rows *sql.Rows) (*structs.StaffLoginCredentials, error) {
	credentials := new(structs.StaffLoginCredentials)
	err := rows.Scan(
		&credentials.CredentialID,
		&credentials.Username,
		&credentials.Password,
		&credentials.CreatedAt,
		&credentials.UpdatedAt)
	return credentials, err
}

func ScanIntoAgentLoginCredentials(rows *sql.Rows) (*structs.AgentLoginCredentials, error) {
	credentials := new(structs.AgentLoginCredentials)
	err := rows.Scan(
		&credentials.CredentialID,
		&credentials.Username,
		&credentials.Password,
		&credentials.AgentID,
		&credentials.CreatedAt,
		&credentials.UpdatedAt)
	return credentials, err
}

func ScanIntoAssetType(rows *sql.Rows) (*structs.AssetType, error) {
	assetType := new(structs.AssetType)
	err := rows.Scan(
		&assetType.AssetTypeID,
		&assetType.AssetType,
		&assetType.CreatedAt,
		&assetType.UpdatedAt)
	return assetType, err
}

func ScanIntoAssetAssignment(rows *sql.Rows) (*structs.AssetAssignment, error) {
	assignment := new(structs.AssetAssignment)
	err := rows.Scan(
		&assignment.AssignmentID,
		&assignment.AssetID,
		&assignment.StaffID,
		&assignment.AssignedBy,
		&assignment.AssignmentType,
		&assignment.DueAt,
		&assignment.CreatedAt,
		&assignment.UpdatedAt)
	return assignment, err
}
