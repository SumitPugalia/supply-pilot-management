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
	Id         string             `db:"id"`
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
		pilots = append(pilots, entity.Pilot(pilot))
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
	return entity.Pilot(pilot), nil
}

func (repo *PilotRepo) CreatePilot(entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := Pilot(entity_pilot)
	_, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return entity_pilot, nil
}

func (repo *PilotRepo) UpdatePilot(id string, entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := Pilot(entity_pilot)
	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}

	return entity_pilot, nil
}
