package service

import "pilot-management/domain"

//------------------------------------------------------------
// Service interface for pilot in our system.
//-------------------------------------------------------------
type Service interface {
	ListPilots(params ListPilotParams) ([]domain.Pilot, uint, uint, error)
	GetPilot(id string) (domain.Pilot, error)
	CreatePilot(params CreatePilotParams) (domain.Pilot, error)
	UpdatePilot(id string, params UpdatePilotParams) (domain.Pilot, error)
	ChangePilotStatus(id string, status string) (domain.Pilot, error)
	DeletePilot(id string) error
}

//------------------------------------------------------------
// Input Parameters type in our system.
//-------------------------------------------------------------

type CreatePilotParams struct {
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

type UpdatePilotParams struct {
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

type ListPilotParams struct {
	SupplierId string
	MarketId   string
	ServiceId  string
	CodeName   string
	Status     string
	Page       uint
	PageSize   uint
}
