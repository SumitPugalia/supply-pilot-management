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

func (s ServiceImpl) ListPilots() ([]entity.Pilot, error) {
	return s.pilotRepo.ListPilots()
}

func (s ServiceImpl) GetPilot(id string) (entity.Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

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

func (s ServiceImpl) DeletePilot(id string) error {
	pilot, err := s.pilotRepo.GetPilot(id)
	if err != nil {
		return err
	}
	pilot.Deleted = true
	_, err = s.pilotRepo.UpdatePilot(id, pilot)
	return err
}

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

func genUUID() string {
	id := guuid.New()
	return id.String()
}
