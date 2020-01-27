package service

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"

	guuid "github.com/google/uuid"
)

//------------------------------------------------------------
// Service interface for pilot in our system.
//-------------------------------------------------------------
type Service interface {
	ListPilots(params ListPilotParams) ([]domain.Pilot, uint, uint, error)
	GetPilot(id guuid.UUID) (domain.Pilot, error)
	CreatePilot(params CreatePilotParams) (domain.Pilot, error)
	UpdatePilot(id guuid.UUID, params UpdatePilotParams) (domain.Pilot, error)
	ChangePilotStatus(id guuid.UUID, status string) (domain.Pilot, error)
	DeletePilot(id guuid.UUID) error
}

//------------------------------------------------------------
// Input Parameters type in our system.
//-------------------------------------------------------------

type CreatePilotParams struct {
	UserId     guuid.UUID
	CodeName   string
	SupplierId guuid.UUID
	MarketId   guuid.UUID
	ServiceId  guuid.UUID
}

type UpdatePilotParams struct {
	CodeName  string
	MarketId  guuid.UUID
	ServiceId guuid.UUID
}

type ListPilotParams struct {
	SupplierId guuid.UUID
	MarketId   guuid.UUID
	ServiceId  guuid.UUID
	CodeName   string
	Status     string
	Page       uint
	PageSize   uint
}
