package converter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
	frontendstructs "github.com/shuttlersIT/itsm-mvp/backend/structs/frontend"
)

func FrontEndTicket(c *gin.Context, t *structs.Ticket) *frontendstructs.FrontendTicket {
	var ticket frontendstructs.FrontendTicket
	ticket.ID = t.ID
	_, a := handlers.GetAsset2(c, t.AssetID)
	ticket.Asset = a.AssetName
	ticket.Subject = t.Subject
	ticket.Description = t.Description
	b, _ := handlers.GetCategory(c, t.Category)
	ticket.Category = b.CategoryName
	f, _ := handlers.GetSubCategory(c, t.SubCategory)
	ticket.SubCategory = f.SubCategoryName
	g, _ := handlers.GetSla(c, t.SLA)
	ticket.SLA = g.SlaName
	h, _ := handlers.GetPriority(c, t.Priority)
	ticket.Priority = h.Name
	i, _ := handlers.GetUserByID(c, t.StaffID)
	ticket.Staff = fmt.Sprintf("%v %v", i.FirstName, i.LastName)
	j, _ := handlers.GetAgent(c, t.AgentID)
	ticket.Agent = fmt.Sprintf("%v %v", j.FirstName, j.LastName)
	ticket.Tag = t.Tag
	ticket.Status = t.Status
	ticket.Site = t.Site
	ticket.RelatedTicket = t.RelatedTicketID
	ticket.CreatedAt = t.CreatedAt
	ticket.UpdatedAt = t.UpdatedAt
	ticket.DueAt = t.DueAt

	return &ticket
}

// FrontEndTicketList efficiently converts a slice of structs.Ticket into a slice of frontendstructs.FrontendTicket
func FrontEndTicketList(c *gin.Context, tkts []*structs.Ticket) []*frontendstructs.FrontendTicket {
	tickets := make([]*frontendstructs.FrontendTicket, len(tkts))

	for i, t := range tkts {
		tickets[i] = convertToFrontEndTicket(c, t)
	}

	return tickets
}

// convertToFrontEndTicket is a helper function to convert a single structs.Ticket to frontendstructs.FrontendTicket
func convertToFrontEndTicket(c *gin.Context, t *structs.Ticket) *frontendstructs.FrontendTicket {
	var ticket frontendstructs.FrontendTicket
	ticket.ID = t.ID
	_, a := handlers.GetAsset2(c, t.AssetID)
	ticket.Asset = a.AssetName
	ticket.Subject = t.Subject
	ticket.Description = t.Description
	b, _ := handlers.GetCategory(c, t.Category)
	ticket.Category = b.CategoryName
	f, _ := handlers.GetSubCategory(c, t.SubCategory)
	ticket.SubCategory = f.SubCategoryName
	g, _ := handlers.GetSla(c, t.SLA)
	ticket.SLA = g.SlaName
	h, _ := handlers.GetPriority(c, t.Priority)
	ticket.Priority = h.Name
	i, _ := handlers.GetUserByID(c, t.StaffID)
	ticket.Staff = fmt.Sprintf("%v %v", i.FirstName, i.LastName)
	j, _ := handlers.GetAgent(c, t.AgentID)
	ticket.Agent = fmt.Sprintf("%v %v", j.FirstName, j.LastName)
	ticket.Tag = t.Tag
	ticket.Status = t.Status
	ticket.Site = t.Site
	ticket.RelatedTicket = t.RelatedTicketID
	ticket.CreatedAt = t.CreatedAt
	ticket.UpdatedAt = t.UpdatedAt
	ticket.DueAt = t.DueAt

	return &ticket
}
