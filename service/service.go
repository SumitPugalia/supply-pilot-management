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
	return s.pilotRepo.CreatePilot(params)
}

func (s ServiceImpl) UpdatePilot(id string, params domain.UpdatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.UpdatePilot(id, params)
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
