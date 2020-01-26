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
	"pilot-management/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
	guuid "github.com/google/uuid"
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
	SupplierId guuid.UUID `json:"supplierId"`
	MarketId   guuid.UUID `json:"marketId"`
	ServiceId  guuid.UUID `json:"serviceId"`
	CodeName   string     `json:"codeName"`
	Status     string     `json:"status"`
	Page       uint       `json:"page"`
	PageSize   uint       `json:"pageSize"`
}

type StatusRequest struct{}

type GetPilotRequest struct {
	Id guuid.UUID `json:"id" validate:"required"`
}

type DeletePilotRequest struct {
	Id guuid.UUID `json:"id" validate:"required"`
}

type CreatePilotRequest struct {
	UserId     guuid.UUID `json:"userId" validate:"required"`
	CodeName   string     `json:"codeName" validate:"required"`
	SupplierId guuid.UUID `json:"supplierId" validate:"required"`
	MarketId   guuid.UUID `json:"marketId" validate:"required"`
	ServiceId  guuid.UUID `json:"serviceId" validate:"required"`
}

type UpdatePilotRequest struct {
	Id        guuid.UUID `json:"id" validate:"required"`
	CodeName  string     `json:"codeName"`
	MarketId  guuid.UUID `json:"marketId"`
	ServiceId guuid.UUID `json:"serviceId"`
}

type ChangePilotStatusRequest struct {
	Id     guuid.UUID `json:"id" validate:"required"`
	Status string     `json:"status" validate:"required"`
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
	var supplierId, marketId, serviceId string
	var supplierID, marketID, serviceID guuid.UUID
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

	supplierId = r.URL.Query().Get("supplierId")
	if supplierId != "" {
		supplierID, err = guuid.Parse(supplierId)
		if err != nil {
			return nil, err
		}
	}

	marketId = r.URL.Query().Get("marketId")
	if marketId != "" {
		marketID, err = guuid.Parse(marketId)
		if err != nil {
			return nil, err
		}
	}

	serviceId = r.URL.Query().Get("serviceId")
	if serviceId != "" {
		serviceID, err = guuid.Parse(serviceId)
		if err != nil {
			return nil, err
		}
	}

	request := ListPilotsRequest{
		SupplierId: supplierID,
		MarketId:   marketID,
		ServiceId:  serviceID,
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

	Id, err := guuid.Parse(id)
	if err != nil {
		return nil, domain.PilotDoesNotExistError
	}

	return GetPilotRequest{Id: Id}, nil
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

	Id, err := guuid.Parse(id)
	if err != nil {
		return nil, domain.PilotDoesNotExistError
	}

	var req UpdatePilotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.Id = Id
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

	Id, err := guuid.Parse(id)
	if err != nil {
		return nil, domain.PilotDoesNotExistError
	}

	var req ChangePilotStatusRequest
	req.Id = Id
	req.Status = status
	return req, nil
}

func DecodeDeletePilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, BadRequestError
	}

	Id, err := guuid.Parse(id)
	if err != nil {
		return nil, domain.PilotDoesNotExistError
	}
	return DeletePilotRequest{Id: Id}, nil
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
