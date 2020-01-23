package service

import (
	"time"

	guuid "github.com/google/uuid"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/repository"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/repository/impl/postgresql"
)

type ServiceImpl struct {
	pilotRepo repository.PilotRepo
}

func MakeServiceImpl() ServiceImpl {
	pilotRepo := postgresql.MakePostgresPilotRepo()
	return ServiceImpl{pilotRepo: &pilotRepo}
}

// it returns the list of all the pilots with the filters and the pagination

func (s ServiceImpl) ListPilots(params domain.ListPilotParams) ([]entity.Pilot, uint, uint, error) {
	var status entity.PilotStatus
	var err error
	if params.Status != "" {
		status, err = pilotStatus(params.Status)
		if err != nil {
			return []entity.Pilot{}, 0, 0, err
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

// it returns the details of the pilot identified by the id
func (s ServiceImpl) GetPilot(id string) (entity.Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

// it creates the pilot as the params received
func (s ServiceImpl) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	now := time.Now()
	pilot := entity.Pilot{
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

// it updates detail of the pilot as the new params received
func (s ServiceImpl) UpdatePilot(id string, params domain.UpdatePilotParams) (entity.Pilot, error) {
	pilot, err := s.pilotRepo.GetPilot(id)
	if err != nil {
		return entity.Pilot{}, err
	}
	pilot.UserId = params.UserId
	pilot.CodeName = params.CodeName
	pilot.SupplierId = params.SupplierId
	pilot.MarketId = params.MarketId
	pilot.ServiceId = params.ServiceId
	pilot.UpdatedAt = time.Now()
	return s.pilotRepo.UpdatePilot(id, pilot)
}

// it soft deletes the pilot identified by id
func (s ServiceImpl) DeletePilot(id string) error {
	pilot, err := s.pilotRepo.GetPilot(id)
	if err != nil {
		return err
	}
	pilot.Deleted = true
	_, err = s.pilotRepo.UpdatePilot(id, pilot)
	return err
}

// it changes the status of the pilot
func (s ServiceImpl) ChangePilotStatus(id string, status string) (entity.Pilot, error) {
	pilot, err := s.pilotRepo.GetPilot(id)
	if err != nil {
		return entity.Pilot{}, err
	}

	pilotStatus, err := pilotStatus(status)
	if err != nil {
		return entity.Pilot{}, err
	}
	pilot.Status = pilotStatus
	pilot.UpdatedAt = time.Now()
	return s.pilotRepo.UpdatePilot(id, pilot)
}

// list of the status pilot can hold
func pilotStatus(status string) (entity.PilotStatus, error) {
	switch status {
	case "idle":
		return entity.IdlePilotStatus, nil
	case "active":
		return entity.ActivePilotStatus, nil
	case "offduty":
		return entity.OffDutyPilotStatus, nil
	case "break":
		return entity.BreakPilotStatus, nil
	case "suspend":
		return entity.SuspendPilotStatus, nil
	default:
		return entity.IdlePilotStatus, entity.InvalidPilotStatus
	}
}

// generator function for uuid
func genUUID() string {
	id := guuid.New()
	return id.String()
}
