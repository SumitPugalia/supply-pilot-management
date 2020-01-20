package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest = errors.New("bad request")
)

var VALIDATE = validator.New()

type ListPilotsRequest struct{}

type GetPilotRequest struct {
	Id string `json:"id"`
}

type DeletePilotRequest struct {
	Id string `json:"id"`
}

type CreatePilotRequest struct {
	UserId     string `json:"userId" validate:"required"`
	CodeName   string `json:"codeName" validate:"required"`
	SupplierId string `json:"supplierId" validate:"required"`
	MarketId   string `json:"marketId" validate:"required"`
	ServiceId  string `json:"serviceId" validate:"required"`
}

type UpdatePilotRequest struct {
	Id         string `json:"id" validate:"required"`
	UserId     string `json:"userId" validate:"required"`
	CodeName   string `json:"codeName" validate:"required"`
	SupplierId string `json:"supplierId" validate:"required"`
	MarketId   string `json:"marketId" validate:"required"`
	ServiceId  string `json:"serviceId" validate:"required"`
}

type ChangeStatePilotRequest struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

func DecodeListPilotsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ListPilotsRequest
	return request, nil
}

func DecodeGetPilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}
	return GetPilotRequest{Id: id}, nil
}

func DecodeCreatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePilotRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	er := validateReq(req)
	if er != nil {
		return nil, er
	}
	return req, nil
}

func DecodeUpdatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	var req UpdatePilotRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.Id = id
	er := validateReq(req)
	if er != nil {
		return nil, er
	}

	return req, nil
}

func DecodeChangeStatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	state, okk := vars["state"]
	if !ok || !okk {
		return nil, ErrBadRequest
	}

	var req ChangeStatePilotRequest
	req.Id = id
	req.State = state
	return req, nil
}

func DecodeDeletePilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}
	return DeletePilotRequest{Id: id}, nil
}

func validateReq(req interface{}) error {
	err := VALIDATE.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
