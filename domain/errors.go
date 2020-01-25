package domain

//------------------------------------------------------------
// This file contains list of system defined errors.
//-------------------------------------------------------------
import "errors"

//------------------------------------------------------------
// This file contains list of system defined errors.
//-------------------------------------------------------------
var (
	PilotDoesNotExistError = errors.New("pilot does not exist")
	InvalidPilotStatus     = errors.New("invalid pilot status")
)
