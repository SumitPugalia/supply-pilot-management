package repository

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

type PilotRepo interface {
	ListPilots(supplierId string,
		marketId string,
		serviceId string,
		codeName string,
		status entity.PilotStatus,
		page uint,
		pageSize uint) ([]entity.Pilot, uint, uint, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(entity_pilot entity.Pilot) (entity.Pilot, error)
	UpdatePilot(id string, entity_pilot entity.Pilot) (entity.Pilot, error)
}
