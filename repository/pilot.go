package repository

//------------------------------------------------------------
// This file contains the repository interface.
//-------------------------------------------------------------
import (
	"pilot-management/domain"
)

//------------------------------------------------------------
// PilotRepo interface.
//------------------------------------------------------------
type PilotRepo interface {
	ListPilots(supplierId string,
		marketId string,
		serviceId string,
		codeName string,
		status domain.PilotStatus,
		page uint,
		pageSize uint) ([]domain.Pilot, uint, uint, error)
	GetPilot(id string) (domain.Pilot, error)
	CreatePilot(domain_pilot domain.Pilot) (domain.Pilot, error)
	UpdatePilot(id string, domain_pilot domain.Pilot) (domain.Pilot, error)
}
