package postgresql

import (
	"time"

	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"

	guuid "github.com/google/uuid"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id         string    `db:"id"`
	UserId     string    `db:"user_id"`
	CodeName   string    `db:"code_name"`
	SupplierId string    `db:"supplier_id"`
	MarketId   string    `db:"market_id"`
	ServiceId  string    `db:"service_id"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Deleted    bool      `db:"deleted"`
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

func (repo *PilotRepo) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	now := time.Now()
	pilot := Pilot{
		Id:         genUUID(),
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		Status:     "IDLE",
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	_, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) UpdatePilot(id string, params domain.UpdatePilotParams) (entity.Pilot, error) {
	pilot := Pilot{}
	err := repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return entity.Pilot{}, entity.PilotDoesNotExistError
		}
	}

	pilot.UserId = params.UserId
	pilot.CodeName = params.CodeName
	pilot.SupplierId = params.SupplierId
	pilot.MarketId = params.MarketId
	pilot.ServiceId = params.ServiceId
	pilot.UpdatedAt = time.Now()

	res := repo.writeConn.Collection("pilots").Find("id", id)
	err = res.Update(pilot)

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

	pilot.Status = string(status)
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

func genUUID() string {
	id := guuid.New()
	return id.String()
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
