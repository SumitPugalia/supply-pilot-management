package entity

import "time"

// pilot struct for the entity
type Pilot struct {
	Id         string      `json:"id"`
	UserId     string      `json:"user_id"`
	CodeName   string      `json:"code_name"`
	SupplierId string      `json:"supplier_id"`
	MarketId   string      `json:"market_id"`
	ServiceId  string      `json:"service_id"`
	Status     PilotStatus `json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted    bool
}

type PilotStatus string

// all possible pilot status
const (
	IdlePilotStatus    PilotStatus = "IDLE"
	ActivePilotStatus  PilotStatus = "ACTIVE"
	OffDutyPilotStatus PilotStatus = "OFFDUTY"
	BreakPilotStatus   PilotStatus = "BREAK"
	SuspendPilotStatus PilotStatus = "SUSPEND"
)
