package impl

//------------------------------------------------------------
// This file contains the implementation of the services.
//-------------------------------------------------------------
import (
	"time"

	"pilot-management/domain"
	"pilot-management/repository"
	"pilot-management/repository/impl/postgresql"
	"pilot-management/service"

	guuid "github.com/google/uuid"
)

type ServiceImpl struct {
	pilotRepo repository.PilotRepo
}

func MakeServiceImpl() ServiceImpl {
	pilotRepo := postgresql.MakePostgresPilotRepo()
	return ServiceImpl{pilotRepo: &pilotRepo}
}

//------------------------------------------------------------
// It returns the list of all the pilots with the filters
// and the pagination.
// Parameter: service.ListPilotParams
// Response: list of domain.Pilot, total entries, total pages, error
//-------------------------------------------------------------
func (s ServiceImpl) ListPilots(params service.ListPilotParams) ([]domain.Pilot, uint, uint, error) {
	var status domain.PilotStatus
	var err error
	if params.Status != "" {
		status, err = pilotStatus(params.Status)
		if err != nil {
			return []domain.Pilot{}, 0, 0, err
		}
	}

	return s.pilotRepo.ListPilots(
		params.SupplierId,
		params.MarketId,
		params.ServiceId,
		params.CodeName,
		status,
		params.Page,
		params.PageSize,
	)
}

//------------------------------------------------------------
// It returns the details of the pilot identified by the id.
// Parameter: id
// Response: domain.Pilot, error
//-------------------------------------------------------------
func (s ServiceImpl) GetPilot(id guuid.UUID) (domain.Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

//------------------------------------------------------------
// It creates the pilot in our system as per the parameter received.
// Parameter: service.CreatePilotParams
// Response: domain.Pilot, error
//-------------------------------------------------------------
func (s ServiceImpl) CreatePilot(params service.CreatePilotParams) (domain.Pilot, error) {
	now := time.Now()
	pilot := domain.Pilot{
		Id:         genUUID(),
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		Status:     "IDLE",
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		CreatedAt:  now,
		UpdatedAt:  now,
		Deleted:    false,
	}
	return s.pilotRepo.CreatePilot(pilot)
}

//------------------------------------------------------------
// It updates the detail of the pilot as the new params received
// Parameter: id, service.UpdatePilotParams
// Response: domain.Pilot, error
//-------------------------------------------------------------
func (s ServiceImpl) UpdatePilot(id guuid.UUID, params service.UpdatePilotParams) (domain.Pilot, error) {
	pilot := domain.Pilot{}
	pilot.CodeName = params.CodeName
	pilot.MarketId = params.MarketId
	pilot.ServiceId = params.ServiceId
	pilot.UpdatedAt = time.Now()
	return s.pilotRepo.UpdatePilot(id, pilot)
}

//------------------------------------------------------------
// It soft deletes the pilot in our system as per the id.
// Parameter: id
// Response: error
//-------------------------------------------------------------
func (s ServiceImpl) DeletePilot(id guuid.UUID) error {
	pilot := domain.Pilot{}
	pilot.UpdatedAt = time.Now()
	pilot.Deleted = true
	_, err := s.pilotRepo.UpdatePilot(id, pilot)
	return err
}

//------------------------------------------------------------
// It changes the status of the pilot.
// Parameter: id, status
// Response: domain.Pilot, error
//-------------------------------------------------------------
func (s ServiceImpl) ChangePilotStatus(id guuid.UUID, status string) (domain.Pilot, error) {
	pilotStatus, err := pilotStatus(status)
	if err != nil {
		return domain.Pilot{}, err
	}
	pilot := domain.Pilot{}
	pilot.Status = pilotStatus
	pilot.UpdatedAt = time.Now()
	return s.pilotRepo.UpdatePilot(id, pilot)
}

//------------------------------------------------------------
// Internal/Helper function to convert the status(string) to
// status(domain.PilotStatus)
//-------------------------------------------------------------
func pilotStatus(status string) (domain.PilotStatus, error) {
	switch status {
	case "idle":
		return domain.IdlePilotStatus, nil
	case "active":
		return domain.ActivePilotStatus, nil
	case "offduty":
		return domain.OffDutyPilotStatus, nil
	case "break":
		return domain.BreakPilotStatus, nil
	case "suspend":
		return domain.SuspendPilotStatus, nil
	default:
		return domain.IdlePilotStatus, domain.InvalidPilotStatus
	}
}

//------------------------------------------------------------
// Internal/Helper function to create uuid.
//-------------------------------------------------------------
func genUUID() guuid.UUID {
	id := guuid.New()
	return id
}
