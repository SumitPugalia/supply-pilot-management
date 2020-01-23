package domain

import "gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"

// list of service interfaces
type Service interface {
	ListPilots(params ListPilotParams) ([]entity.Pilot, uint, uint, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(params CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(id string, params UpdatePilotParams) (entity.Pilot, error)
	ChangePilotStatus(id string, status string) (entity.Pilot, error)
	DeletePilot(id string) error
}

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
