package endpoint

import (
	"pilot-management/domain"

	guuid "github.com/google/uuid"
)

//------------------------------------------------------------
// This file contains view or the presentation of the response.
//-------------------------------------------------------------

//------------------------------------------------------------
// PilotView struct for representation of pilot for response.
//-------------------------------------------------------------
type PilotView struct {
	Id         guuid.UUID `json:"id"`
	UserId     guuid.UUID `json:"userId"`
	CodeName   string     `json:"codeName"`
	SupplierId guuid.UUID `json:"supplierId"`
	MarketId   guuid.UUID `json:"marketId"`
	ServiceId  guuid.UUID `json:"serviceId"`
	Status     string     `json:"status"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
}

func ToPilotView(pilot domain.Pilot) PilotView {
	return PilotView{
		Id:         pilot.Id,
		UserId:     pilot.UserId,
		CodeName:   pilot.CodeName,
		SupplierId: pilot.SupplierId,
		MarketId:   pilot.MarketId,
		ServiceId:  pilot.ServiceId,
		Status:     string(pilot.Status),
		CreatedAt:  pilot.CreatedAt.Unix(),
		UpdatedAt:  pilot.UpdatedAt.Unix(),
	}
}

func ToPilotViews(pilots []domain.Pilot) []PilotView {
	views := make([]PilotView, 0)
	for _, pilot := range pilots {
		views = append(views, ToPilotView(pilot))
	}
	return views
}
