package postgresql

//------------------------------------------------------------
// This file contains the implementation of the repo interface.
//-------------------------------------------------------------
import (
	"time"

	"pilot-management/domain"

	guuid "github.com/google/uuid"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	tableName = "pilots"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

//------------------------------------------------------------
// Pilot struct for database.
//-------------------------------------------------------------
type Pilot struct {
	Id         guuid.UUID         `db:"id"`
	UserId     guuid.UUID         `db:"user_id"`
	CodeName   string             `db:"code_name"`
	SupplierId guuid.UUID         `db:"supplier_id"`
	MarketId   guuid.UUID         `db:"market_id"`
	ServiceId  guuid.UUID         `db:"service_id"`
	Status     domain.PilotStatus `db:"status"`
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

//------------------------------------------------------------
// List all the pilots based on the filters and the pagination.
// Updates the pilot information for the id provided.
// Parameter: supplierId, marketId, serviceId, codeName, status,
//            page and pageSize
// Response: list of domain.Pilot, total entries, total pages, error
//-------------------------------------------------------------
func (repo *PilotRepo) ListPilots(
	supplierId guuid.UUID,
	marketId guuid.UUID,
	serviceId guuid.UUID,
	codeName string,
	status domain.PilotStatus,
	page uint,
	pageSize uint) ([]domain.Pilot, uint, uint, error) {

	rows := make([]Pilot, 0)
	pilots := make([]domain.Pilot, 0)

	query := repo.readConn.Collection(tableName).Find(db.Cond{"deleted": false})
	query = query.Paginate(pageSize).Page(page)

	Nil := guuid.UUID{}

	if supplierId != Nil {
		query = query.And(db.Cond{"supplier_id": supplierId})
	}
	if marketId != Nil {
		query = query.And(db.Cond{"market_id": marketId})
	}
	if serviceId != Nil {
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
			return []domain.Pilot{}, 0, 0, domain.PilotDoesNotExistError
		}
		return pilots, 0, 0, err
	}
	for _, pilot := range rows {
		pilots = append(pilots, domain.Pilot(pilot))
	}
	return pilots, uint(totalEntries), totalPages, nil
}

//------------------------------------------------------------
// Get the pilot information for the id provided.
// Parameter: id
// Response: domain.Pilot or error
//-------------------------------------------------------------
func (repo *PilotRepo) GetPilot(id guuid.UUID) (domain.Pilot, error) {
	var pilot Pilot
	err := repo.readConn.Collection(tableName).Find(db.Cond{"id =": id, "deleted =": false}).One(&pilot)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return domain.Pilot{}, domain.PilotDoesNotExistError
		}
		return domain.Pilot{}, err
	}
	return domain.Pilot(pilot), nil
}

//------------------------------------------------------------
// Create the pilot.
// Parameter: domain.Pilot
// Response: domain.Pilot or error
//-------------------------------------------------------------
func (repo *PilotRepo) CreatePilot(domain_pilot domain.Pilot) (domain.Pilot, error) {
	pilot := Pilot(domain_pilot)
	_, err := repo.writeConn.Collection(tableName).Insert(pilot)
	if err != nil {
		return domain.Pilot{}, err
	}

	return domain_pilot, nil
}

//------------------------------------------------------------
// Updates the pilot information for the id provided.
// Parameter: id, domain.Pilot
// Response: domain.Pilot or error
//-------------------------------------------------------------
func (repo *PilotRepo) UpdatePilot(id guuid.UUID, domain_pilot domain.Pilot) (domain.Pilot, error) {
	updates := []interface{}{
		"updated_at",
		domain_pilot.UpdatedAt,
	}

	Nil := guuid.UUID{}

	if len(domain_pilot.CodeName) > 0 {
		updates = append(updates, "code_name", domain_pilot.CodeName)
	}

	if domain_pilot.MarketId != Nil {
		updates = append(updates, "market_id", domain_pilot.MarketId)
	}

	if domain_pilot.ServiceId != Nil {
		updates = append(updates, "service_id", domain_pilot.ServiceId)
	}

	if len(domain_pilot.Status) > 0 {
		updates = append(updates, "status", domain_pilot.Status)
	}

	if domain_pilot.Deleted {
		updates = append(updates, "deleted", domain_pilot.Deleted)
	}

	q := repo.writeConn.Update("pilots").Set(updates...).Where("id = ? AND deleted = ?", id, false)
	res, err := q.Exec()

	rows, _ := res.RowsAffected()

	if err != nil {
		return domain.Pilot{}, err
	}

	if rows == 1 {
		var pilot Pilot
		err = repo.writeConn.Collection(tableName).Find(db.Cond{"id =": id, "deleted =": false}).One(&pilot)
		if err != nil && !domain_pilot.Deleted {
			return domain.Pilot{}, err
		}
		return domain.Pilot(pilot), nil
	}

	return domain.Pilot{}, domain.PilotDoesNotExistError
}
