package service

import (
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
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		Status:     "IDLE",
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		CreatedAt:  now,
		UpdatedAt:  now,
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
	return s.pilotRepo.DeletePilot(id)
}

func (s ServiceImpl) ChangePilotStatus(id string, status string) (entity.Pilot, error) {
	switch status {
	case "idle":
		return s.pilotRepo.ChangePilotStatus(id, entity.IdlePilotStatus)
	case "active":
		return s.pilotRepo.ChangePilotStatus(id, entity.ActivePilotStatus)
	case "offduty":
		return s.pilotRepo.ChangePilotStatus(id, entity.OffDutyPilotStatus)
	case "break":
		return s.pilotRepo.ChangePilotStatus(id, entity.BreakPilotStatus)
	case "suspend":
		return s.pilotRepo.ChangePilotStatus(id, entity.SuspendPilotStatus)
	default:
		return entity.Pilot{}, entity.InvalidPilotStatus
	}
}

func genUUID() string {
	id := guuid.New()
	return id.String()
}
