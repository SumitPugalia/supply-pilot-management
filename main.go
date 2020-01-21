package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain/entity"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/endpoint"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/service"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	fmt.Println("About to start the server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func assignRoutes(router *mux.Router) *mux.Router {
	service := service.MakeServiceImpl()
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(encodeErrorResponse),
	}

	statusHandler := httpTransport.NewServer(
		endpoint.MakeStatusEndpoint(service),
		endpoint.DecodeStatusRequest,
		endpoint.EncodeResponse,
	)

	listPilotsHandler := httpTransport.NewServer(
		endpoint.MakeListPilotsEndpoint(service),
		endpoint.DecodeListPilotsRequest,
		endpoint.EncodeResponse,
		options...,
	)

	getPilotHandler := httpTransport.NewServer(
		endpoint.MakeGetPilotEndpoint(service),
		endpoint.DecodeGetPilotRequest,
		endpoint.EncodeResponse,
		options...,
	)

	CreatePilotHandler := httpTransport.NewServer(
		endpoint.MakeCreatePilotEndpoint(service),
		endpoint.DecodeCreatePilotRequest,
		endpoint.EncodeResponse,
		options...,
	)

	UpdatePilotHandler := httpTransport.NewServer(
		endpoint.MakeUpdatePilotEndpoint(service),
		endpoint.DecodeUpdatePilotRequest,
		endpoint.EncodeResponse,
		options...,
	)

	DeletePilotHandler := httpTransport.NewServer(
		endpoint.MakeDeletePilotEndpoint(service),
		endpoint.DecodeDeletePilotRequest,
		endpoint.EncodeResponse,
		options...,
	)

	ChangePilotStatusHandler := httpTransport.NewServer(
		endpoint.MakeChangePilotStatusEndpoint(service),
		endpoint.DecodeChangePilotStatusRequest,
		endpoint.EncodeResponse,
		options...,
	)

	router.Handle("/supply/pilots/status", statusHandler).Methods("GET")
	router.Handle("/supply/pilots", listPilotsHandler).Methods("GET")
	router.Handle("/supply/pilots/{id}", getPilotHandler).Methods("GET")
	router.Handle("/supply/pilots", CreatePilotHandler).Methods("POST")
	router.Handle("/supply/pilots/{id}", UpdatePilotHandler).Methods("PATCH")
	router.Handle("/supply/pilots/{id}", DeletePilotHandler).Methods("DELETE")
	router.Handle("/supply/pilots/{id}/{status}", ChangePilotStatusHandler).Methods("PATCH")
	return router
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if errs, ok := err.(validator.ValidationErrors); ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(endpoint.Response{Errors: encodeV10Errors(errs)})
		return
	}

	statusCode := codeFrom(err)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(endpoint.Response{Errors: []string{err.Error()}})
}

func codeFrom(err error) int {
	switch err {
	case endpoint.ErrBadRequest,
		entity.InvalidPilotStatus:
		return http.StatusBadRequest
	case entity.PilotDoesNotExistError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func encodeV10Errors(errs validator.ValidationErrors) []string {
	var errorsSlice []string
	for _, err := range errs {
		errorsSlice = append(errorsSlice, err.Field()+":"+err.Tag())
	}
	return errorsSlice
}
