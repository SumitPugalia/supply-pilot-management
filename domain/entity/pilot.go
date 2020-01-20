package entity

import "time"

type Pilot struct {
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
	State      PilotState
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PilotState string

const (
	IdlePilotState    PilotState = "IDLE"
	ActivePilotState  PilotState = "ACTIVE"
	OffDutyPilotState PilotState = "OFFDUTY"
	BreakPilotState   PilotState = "BREAK"
	SuspendPilotState PilotState = "SUSPEND"
)
