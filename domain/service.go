package domain

import "gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"

type Service interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(params CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(params UpdatePilotParams) (entity.Pilot, error)
	ChangeStatePilot(id string, state string) (entity.Pilot, error)
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
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}
