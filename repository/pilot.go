package repository

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(entity_pilot entity.Pilot) (entity.Pilot, error)
	UpdatePilot(id string, entity_pilot entity.Pilot) (entity.Pilot, error)
	DeletePilot(id string) error
	ChangePilotStatus(id string, status entity.PilotStatus) (entity.Pilot, error)
}
