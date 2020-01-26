package domain

//------------------------------------------------------------
// This file contains the domain model of pilot for our system.
//-------------------------------------------------------------
import (
	"time"

	guuid "github.com/google/uuid"
)

//------------------------------------------------------------
// Pilot struct for the domain
//-------------------------------------------------------------
type Pilot struct {
	Id         guuid.UUID
	UserId     guuid.UUID
	CodeName   string
	SupplierId guuid.UUID
	MarketId   guuid.UUID
	ServiceId  guuid.UUID
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
