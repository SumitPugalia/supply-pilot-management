package sqlite

import (
	"../../../domain/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id         string `db:"id"`
	UserId     string `db:"user_id"`
	CodeName   string `db:"code_name"`
	SupplierId string `db:"supplier_id"`
	MarketId   string `db:"market_id"`
	ServiceId  string `db:"service_id"`
}

func MakeSqlitePilotRepo() PilotRepo {
	return PilotRepo{
		readConn:  getReadConn(),
		writeConn: getWriteConn(),
	}
}

func (repo *PilotRepo) ListPilots() ([]entity.Pilot, error) {
	resultSet := make([]Pilot, 0)
	err := repo.readConn.Collection("pilots").Find().All(&resultSet)
	if err != nil {
		return nil, err
	}
	pilots := make([]entity.Pilot, 0)
	for _, pilot := range resultSet {
		pilots = append(pilots, entity.Pilot(pilot))
	}
	return pilots, nil
}
