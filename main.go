package main

//------------------------------------------------------------
// This is the main file that contains the router.
//-------------------------------------------------------------
import (
	"fmt"
	"net/http"

	"pilot-management/endpoint"
	"pilot-management/service/impl"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//------------------------------------------------------------
// This is the entry/starting point of our application.
//-------------------------------------------------------------
func main() {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	fmt.Println("About to start the server at port 8080")
	http.ListenAndServe(":8080", nil)
}

//------------------------------------------------------------
// Routers that maps the routes to the endpoints for our system.
//-------------------------------------------------------------
func assignRoutes(router *mux.Router) *mux.Router {
	service := impl.MakeServiceImpl()
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(endpoint.EncodeErrorResponse),
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
