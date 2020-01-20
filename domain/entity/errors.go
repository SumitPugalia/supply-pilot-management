package entity

import "errors"

var PilotDoesNotExistError = errors.New("pilot does not exist")
var InvalidPilotState = errors.New("invalid pilot state")
