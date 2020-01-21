package postgresql

import (
	"time"

	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id         string             `db:"id,omitempty"`
	UserId     string             `db:"user_id"`
	CodeName   string             `db:"code_name"`
	SupplierId string             `db:"supplier_id"`
	MarketId   string             `db:"market_id"`
	ServiceId  string             `db:"service_id"`
	Status     entity.PilotStatus `db:"status"`
	CreatedAt  time.Time          `db:"created_at"`
	UpdatedAt  time.Time          `db:"updated_at"`
	Deleted    bool               `db:"deleted"`
}

func MakePostgresPilotRepo() PilotRepo {
	return PilotRepo{
		readConn:  getReadConn(),
		writeConn: getWriteConn(),
	}
}

func (repo *PilotRepo) ListPilots() ([]entity.Pilot, error) {
	resultSet := make([]Pilot, 0)
	err := repo.readConn.Collection("pilots").Find(db.Cond{"deleted =": false}).All(&resultSet)
	if err != nil {
		return nil, err
	}
	pilots := make([]entity.Pilot, 0)
	for _, pilot := range resultSet {
		pilots = append(pilots, pilotRowToPilot(pilot))
	}
	return pilots, nil
}

func (repo *PilotRepo) GetPilot(id string) (entity.Pilot, error) {
	var pilot Pilot
	err := repo.readConn.Collection("pilots").Find(db.Cond{"id =": id, "deleted =": false}).One(&pilot)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return entity.Pilot{}, entity.PilotDoesNotExistError
		}
		return entity.Pilot{}, err
	}
	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) CreatePilot(entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := entity_pilot_to_db_pilot(entity_pilot)
	id, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	entity_pilot.Id = string(id.([]byte))
	return entity_pilot, nil
}

func (repo *PilotRepo) UpdatePilot(id string, entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := entity_pilot_to_db_pilot(entity_pilot)
	// err := repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	// if err != nil {
	// 	if err == db.ErrNoMoreRows {
	// 		return entity.Pilot{}, entity.PilotDoesNotExistError
	// 	}
	// }

	// pilot.UserId = params.UserId
	// pilot.CodeName = params.CodeName
	// pilot.SupplierId = params.SupplierId
	// pilot.MarketId = params.MarketId
	// pilot.ServiceId = params.ServiceId
	// pilot.UpdatedAt = time.Now()

	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}

	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) ChangePilotStatus(id string, status entity.PilotStatus) (entity.Pilot, error) {
	pilot := Pilot{}
	err := repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return entity.Pilot{}, entity.PilotDoesNotExistError
		}
	}

	pilot.Status = status
	pilot.UpdatedAt = time.Now()

	res := repo.writeConn.Collection("pilots").Find("id", id)
	err = res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}
	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) DeletePilot(id string) error {
	pilot := Pilot{}
	err := repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return entity.PilotDoesNotExistError
		}
	}

	pilot.Deleted = true
	res := repo.writeConn.Collection("pilots").Find("id", id)
	err = res.Update(pilot)
	return err
}

func pilotRowToPilot(row Pilot) entity.Pilot {
	return entity.Pilot{
		Id:         row.Id,
		UserId:     row.UserId,
		SupplierId: row.SupplierId,
		MarketId:   row.MarketId,
		ServiceId:  row.ServiceId,
		CodeName:   row.CodeName,
		Status:     entity.PilotStatus(row.Status),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func entity_pilot_to_db_pilot(entity_pilot entity.Pilot) Pilot {
	return Pilot{
		Id:         entity_pilot.Id,
		UserId:     entity_pilot.UserId,
		SupplierId: entity_pilot.SupplierId,
		MarketId:   entity_pilot.MarketId,
		ServiceId:  entity_pilot.ServiceId,
		CodeName:   entity_pilot.CodeName,
		Status:     entity_pilot.Status,
		CreatedAt:  entity_pilot.CreatedAt,
		UpdatedAt:  entity_pilot.UpdatedAt,
	}
}
