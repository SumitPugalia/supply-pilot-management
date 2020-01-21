package entity

import "errors"

var PilotDoesNotExistError = errors.New("pilot does not exist")
var InvalidPilotStatus = errors.New("invalid pilot status")
