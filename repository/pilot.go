package repository

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(param domain.CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(id string, param domain.UpdatePilotParams) (entity.Pilot, error)
	DeletePilot(id string) error
	ChangePilotStatus(id string, status entity.PilotStatus) (entity.Pilot, error)
}
