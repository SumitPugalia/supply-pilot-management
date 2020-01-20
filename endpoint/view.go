package endpoint

import (
	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
)

type PilotView struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	CodeName   string `json:"code_name"`
	SupplierId string `json:"supplier_id"`
	MarketId   string `json:"market_id"`
	ServiceId  string `json:"service_id"`
	State      string `json:"status"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func toPilotView(pilot entity.Pilot) PilotView {
	return PilotView{
		Id:         pilot.Id,
		UserId:     pilot.UserId,
		CodeName:   pilot.CodeName,
		SupplierId: pilot.SupplierId,
		MarketId:   pilot.MarketId,
		ServiceId:  pilot.ServiceId,
		State:      string(pilot.State),
		CreatedAt:  pilot.CreatedAt.Unix(),
		UpdatedAt:  pilot.UpdatedAt.Unix(),
	}
}
