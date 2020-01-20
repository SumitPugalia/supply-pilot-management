package repository

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(param domain.CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(param domain.UpdatePilotParams) (entity.Pilot, error)
	DeletePilot(id string) error
	ChangeStatePilot(id string, state entity.PilotState) (entity.Pilot, error)
}
