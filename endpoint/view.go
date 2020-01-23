package endpoint

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

// pilot struct for representation of pilot for response
type PilotView struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

func ToPilotView(pilot entity.Pilot) PilotView {
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

func ToPilotViews(pilots []entity.Pilot) []PilotView {
	views := make([]PilotView, 0)
	for _, pilot := range pilots {
		views = append(views, ToPilotView(pilot))
	}
	return views
}
