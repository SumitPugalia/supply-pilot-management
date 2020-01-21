package entity

import "time"

// pilot struct for the entity
type Pilot struct {
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
	Status     PilotStatus
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
