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

func FrontEndAgent(c *gin.Context, t structs.Agent) *frontendstructs.FrontendAgent {
	var agent frontendstructs.FrontendAgent
	agent.AgentID = t.AgentID
	agent.FirstName = t.FirstName
	agent.LastName = t.LastName
	agent.AgentEmail = t.AgentEmail
	agent.Phone = t.Phone
	w, _ := handlers.GetRole(c, t.RoleID)
	agent.Role = w.RoleName
	x, _ := handlers.GetUnit(c, t.Unit)
	agent.Unit = x.UnitName
	y, _ := handlers.GetCred2(c, t.Unit)
	agent.Username = y.Username
	z, _ := handlers.GetAgent(c, t.SupervisorID)
	agent.Supervisor = fmt.Sprintf("%v %v", z.FirstName, z.LastName)
	agent.CreatedAt = t.CreatedAt
	agent.UpdatedAt = t.UpdatedAt

	return &agent
}

func FrontEndAgentList(c *gin.Context, t structs.Agent) *frontendstructs.FrontendAgent {
	var agent frontendstructs.FrontendAgent
	agent.AgentID = t.AgentID
	agent.FirstName = t.FirstName
	agent.LastName = t.LastName
	agent.AgentEmail = t.AgentEmail
	agent.Phone = t.Phone
	w, _ := handlers.GetRole(c, t.RoleID)
	agent.Role = w.RoleName
	x, _ := handlers.GetUnit(c, t.Unit)
	agent.Unit = x.UnitName
	y, _ := handlers.GetCred2(c, t.Unit)
	agent.Username = y.Username
	z, _ := handlers.GetAgent(c, t.SupervisorID)
	agent.Supervisor = fmt.Sprintf("%v %v", z.FirstName, z.LastName)
	agent.CreatedAt = t.CreatedAt
	agent.UpdatedAt = t.UpdatedAt

	return &agent
}

// FrontEndAgent efficiently converts a structs.Agent into a frontendstructs.FrontendAgent
func FrontEndAgentB(c *gin.Context, a *structs.Agent) *frontendstructs.FrontendAgent {
	var agent frontendstructs.FrontendAgent
	agent.AgentID = a.AgentID
	agent.FirstName = a.FirstName
	agent.LastName = a.LastName
	agent.AgentEmail = a.AgentEmail
	y, _ := handlers.GetCred2(c, a.Unit)
	agent.Username = y.Username
	agent.Phone = a.Phone
	r, _ := handlers.GetRole(c, a.RoleID)
	agent.Role = r.RoleName
	x, _ := handlers.GetUnit(c, a.Unit)
	agent.Unit = x.UnitName
	z, _ := handlers.GetAgent(c, a.SupervisorID)
	agent.Supervisor = fmt.Sprintf("%v %v", z.FirstName, z.LastName)
	agent.CreatedAt = a.CreatedAt
	agent.UpdatedAt = a.UpdatedAt

	return &agent
}

// FrontEndAgentList efficiently converts a slice of structs.Agent into a slice of frontendstructs.FrontendAgent
func FrontEndAgentListB(c *gin.Context, agentList []*structs.Agent) []*frontendstructs.FrontendAgent {
	frontendAgentList := make([]*frontendstructs.FrontendAgent, len(agentList))

	for i, a := range agentList {
		frontendAgentList[i] = FrontEndAgentB(c, a)
	}

	return frontendAgentList
}
