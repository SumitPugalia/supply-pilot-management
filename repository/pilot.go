package repository

//------------------------------------------------------------
// This file contains the repository interface.
//-------------------------------------------------------------
import (
	"pilot-management/domain"

	guuid "github.com/google/uuid"
)

//------------------------------------------------------------
// PilotRepo interface.
//------------------------------------------------------------
type PilotRepo interface {
	ListPilots(supplierId guuid.UUID,
		marketId guuid.UUID,
		serviceId guuid.UUID,
		codeName string,
		status domain.PilotStatus,
		page uint,
		pageSize uint) ([]domain.Pilot, uint, uint, error)
	GetPilot(id guuid.UUID) (domain.Pilot, error)
	CreatePilot(domain_pilot domain.Pilot) (domain.Pilot, error)
	UpdatePilot(id guuid.UUID, domain_pilot domain.Pilot) (domain.Pilot, error)
}
