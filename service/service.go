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

func (s ServiceImpl) UpdatePilot(params domain.UpdatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.UpdatePilot(params)
}

func (s ServiceImpl) DeletePilot(id string) error {
	return s.pilotRepo.DeletePilot(id)
}

func (s ServiceImpl) ChangeStatePilot(id string, state string) (entity.Pilot, error) {
	switch state {
	case "idle":
		return s.pilotRepo.ChangeStatePilot(id, entity.IdlePilotState)
	case "active":
		return s.pilotRepo.ChangeStatePilot(id, entity.ActivePilotState)
	case "offduty":
		return s.pilotRepo.ChangeStatePilot(id, entity.OffDutyPilotState)
	case "break":
		return s.pilotRepo.ChangeStatePilot(id, entity.BreakPilotState)
	case "suspend":
		return s.pilotRepo.ChangeStatePilot(id, entity.SuspendPilotState)
	default:
		return entity.Pilot{}, entity.InvalidPilotState
	}
}
