package converter

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
	frontendstructs "github.com/shuttlersIT/itsm-mvp/backend/structs/frontend"
)

func FrontEndStaff(c *gin.Context, t structs.Staff) *frontendstructs.FrontendStaff {
	var staff frontendstructs.FrontendStaff
	staff.StaffID = t.StaffID
	staff.FirstName = t.FirstName
	staff.LastName = t.LastName
	staff.StaffEmail = t.StaffEmail
	staff.Phone = t.Phone
	w, _ := handlers.GetPosition(c, t.PositionID)
	staff.Position = w.PositionName
	x, _ := handlers.GetDepartment(c, t.DepartmentID)
	staff.Department = x.DepartmentName
	y, _ := handlers.GetAgentCred2(c, t.Username)
	staff.Username = y.Username
	staff.CreatedAt = t.CreatedAt
	staff.UpdatedAt = t.UpdatedAt

	return &staff
}

func FrontEndStaffList(c *gin.Context, t structs.Staff) *frontendstructs.FrontendStaff {
	var staff frontendstructs.FrontendStaff
	staff.StaffID = t.StaffID
	staff.FirstName = t.FirstName
	staff.LastName = t.LastName
	staff.StaffEmail = t.StaffEmail
	staff.Phone = t.Phone
	w, _ := handlers.GetPosition(c, t.PositionID)
	staff.Position = w.PositionName
	x, _ := handlers.GetDepartment(c, t.DepartmentID)
	staff.Department = x.DepartmentName
	y, _ := handlers.GetAgentCred2(c, t.Username)
	staff.Username = y.Username
	staff.CreatedAt = t.CreatedAt
	staff.UpdatedAt = t.UpdatedAt

	return &staff
}

// FrontEndStaff efficiently converts a structs.Staff into a frontendstructs.FrontendStaff
func FrontEndStaffB(c *gin.Context, s *structs.Staff) *frontendstructs.FrontendStaff {
	var staff frontendstructs.FrontendStaff
	staff.StaffID = s.StaffID
	staff.FirstName = s.FirstName
	staff.LastName = s.LastName
	staff.StaffEmail = s.StaffEmail
	y, _ := handlers.GetAgentCred2(c, s.Username)
	staff.Username = y.Username
	staff.Phone = s.Phone
	p, _ := handlers.GetPosition(c, s.PositionID)
	staff.Position = p.PositionName
	d, _ := handlers.GetDepartment(c, s.DepartmentID)
	staff.Department = d.DepartmentName
	staff.CreatedAt = s.CreatedAt
	staff.UpdatedAt = s.UpdatedAt

	return &staff
}

// FrontEndStaffList efficiently converts a slice of structs.Staff into a slice of frontendstructs.FrontendStaff
func FrontEndStaffListB(c *gin.Context, staffList []*structs.Staff) []*frontendstructs.FrontendStaff {
	frontendStaffList := make([]*frontendstructs.FrontendStaff, len(staffList))

	for i, s := range staffList {
		frontendStaffList[i] = FrontEndStaffB(c, s)
	}

	return frontendStaffList
}
