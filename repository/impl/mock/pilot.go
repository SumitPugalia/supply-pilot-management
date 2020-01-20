package mock

import "../../../domain/entity"

type PilotRepoMock struct {
	store map[string]entity.Pilot
}

func (r *PilotRepoMock) ListPilots() ([]entity.Pilot, error) {
	return []entity.Pilot{}, nil
}
