package domain

//------------------------------------------------------------
// This file contains the domain model of pilot for our system.
//-------------------------------------------------------------
import "time"

//------------------------------------------------------------
// Pilot struct for the domain
//-------------------------------------------------------------
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

//------------------------------------------------------------
// Pilot Status type and the values it can hold
//-------------------------------------------------------------
type PilotStatus string

const (
	IdlePilotStatus    PilotStatus = "IDLE"
	ActivePilotStatus  PilotStatus = "ACTIVE"
	OffDutyPilotStatus PilotStatus = "OFFDUTY"
	BreakPilotStatus   PilotStatus = "BREAK"
	SuspendPilotStatus PilotStatus = "SUSPEND"
)
