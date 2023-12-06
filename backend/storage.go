package main

import (
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

type AgentPasswordLoginStorage interface {
	CreateTicketOperation(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type StaffPasswordLoginStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type GoogeAuthStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type TicketStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type StaffStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type AgentStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type AssetStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]**structs.Ticket, error)
	GetTicketByID(int) (**structs.Ticket, error)
	GetTicketByNumber(int) (**structs.Ticket, error)
}

type SlaStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type PriorityStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type SatisfactionStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type PositionStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type DepartmentStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type UnitStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type RoleStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type CategoryStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type SubCategoryStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type StatusStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type AssetTypeStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}

type AssetAssignmentStorage interface {
	CreateTicket(*structs.Ticket) error
	DeleteTicket(int) error
	UpdateTicket(*structs.Ticket) error
	GetTickets() ([]*structs.Ticket, error)
	GetTicketByID(int) (*structs.Ticket, error)
	GetTicketByNumber(int) (*structs.Ticket, error)
}
