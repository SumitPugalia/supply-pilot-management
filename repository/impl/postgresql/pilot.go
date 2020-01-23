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

// pilot struct for database
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

// list all the pilots
func (repo *PilotRepo) ListPilots(
	supplierId string,
	marketId string,
	serviceId string,
	codeName string,
	status entity.PilotStatus,
	page uint,
	pageSize uint) ([]entity.Pilot, uint, uint, error) {

	rows := make([]Pilot, 0)
	pilots := make([]entity.Pilot, 0)

	query := repo.readConn.Collection("pilots").Find(db.Cond{"deleted": false})
	query = query.Paginate(pageSize).Page(page)
	if supplierId != "" {
		query = query.And(db.Cond{"supplier_id": supplierId})
	}
	if marketId != "" {
		query = query.And(db.Cond{"market_id": marketId})
	}
	if serviceId != "" {
		query = query.And(db.Cond{"service_id": serviceId})
	}
	if codeName != "" {
		query = query.And(db.Cond{"code_name": codeName})
	}
	if len(status) > 0 {
		query = query.And(db.Cond{"status": status})
	}
	err := query.All(&rows)
	totalEntries, err := query.TotalEntries()
	totalPages, err := query.TotalPages()
	if err != nil {
		if err == db.ErrNoMoreRows {
			return []entity.Pilot{}, 0, 0, entity.PilotDoesNotExistError
		}
		return pilots, 0, 0, err
	}
	for _, pilot := range rows {
		pilots = append(pilots, entity.Pilot(pilot))
	}
	return pilots, uint(totalEntries), totalPages, nil
}

// get the detail of the pilot
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

// create the pilot
func (repo *PilotRepo) CreatePilot(entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := Pilot(entity_pilot)
	_, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return entity_pilot, nil
}

// update the detail of the pilot
func (repo *PilotRepo) UpdatePilot(id string, entity_pilot entity.Pilot) (entity.Pilot, error) {
	pilot := Pilot(entity_pilot)
	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}

	return entity_pilot, nil
}
