package endpoint

//------------------------------------------------------------
// This file contains all the requests type that our system
// expect, functions to convert the http request to our system
// request type and run the validations on the attributes
//-------------------------------------------------------------

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

//------------------------------------------------------------
// Constants:
// 	BadRequestError - malformed request
//  VALIDATE        - instance of 'validate' with sane defaults.
//-------------------------------------------------------------

var (
	BadRequestError = errors.New("bad request")
	VALIDATE        = validator.New()
)

//------------------------------------------------------------
//	Valid Request Types that our system expects.
//  We can use our validator logic for each of the attributes
//	of the request
//-------------------------------------------------------------

type ListPilotsRequest struct {
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
	CodeName   string `json:"codeName"`
	Status     string `json:"status"`
	Page       uint   `json:"page"`
	PageSize   uint   `json:"pageSize"`
}

type StatusRequest struct{}

type GetPilotRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeletePilotRequest struct {
	Id string `json:"id" validate:"required"`
}

type CreatePilotRequest struct {
	UserId     string `json:"userId" validate:"required"`
	CodeName   string `json:"codeName" validate:"required"`
	SupplierId string `json:"supplierId" validate:"required"`
	MarketId   string `json:"marketId" validate:"required"`
	ServiceId  string `json:"serviceId" validate:"required"`
}

type UpdatePilotRequest struct {
	Id        string `json:"id" validate:"required"`
	CodeName  string `json:"codeName"`
	MarketId  string `json:"marketId"`
	ServiceId string `json:"serviceId"`
}

type ChangePilotStatusRequest struct {
	Id     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

//------------------------------------------------------------
//	Functions to Decode / Map the request to Valid Request Types.
//  It intakes two paramters context and http request
//  Responds with interface{} that is one of Request Type
//  or error incase of invalid http request received
//-------------------------------------------------------------

func DecodeStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request StatusRequest
	return request, nil
}

func DecodeListPilotsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var err error
	var page, pageSize uint64
	pageQuery := r.URL.Query().Get("page")
	if pageQuery != "" {
		page, err = strconv.ParseUint(pageQuery, 10, 32)
	}

	pageSizeQuery := r.URL.Query().Get("pageSize")
	if pageSizeQuery != "" {
		pageSize, err = strconv.ParseUint(pageSizeQuery, 10, 32)
	}
	if err != nil {
		return nil, err
	}

	request := ListPilotsRequest{
		SupplierId: r.URL.Query().Get("supplierId"),
		MarketId:   r.URL.Query().Get("marketId"),
		ServiceId:  r.URL.Query().Get("serviceId"),
		CodeName:   r.URL.Query().Get("codeName"),
		Status:     r.URL.Query().Get("status"),
		Page:       uint(page),
		PageSize:   uint(pageSize),
	}
	err = validateReq(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetPilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, BadRequestError
	}
	return GetPilotRequest{Id: id}, nil
}

func DecodeCreatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePilotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	err = validateReq(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, BadRequestError
	}

	var req UpdatePilotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.Id = id
	err = validateReq(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeChangePilotStatusRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	status, okk := vars["status"]
	if !ok || !okk {
		return nil, BadRequestError
	}

	var req ChangePilotStatusRequest
	req.Id = id
	req.Status = status
	return req, nil
}

func DecodeDeletePilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, BadRequestError
	}
	return DeletePilotRequest{Id: id}, nil
}

//------------------------------------------------------------
//	Private function to check the validations
//-------------------------------------------------------------

func validateReq(req interface{}) error {
	err := VALIDATE.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
