package endpoint

//------------------------------------------------------------
// This is the main file that contains the router.
//-------------------------------------------------------------
import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/service/impl"
)

//------------------------------------------------------------
// This is the entry/starting point of our application.
//-------------------------------------------------------------
func StartApp(port string) {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	fmt.Println("About to start the server at port", port)
	http.ListenAndServe(port, nil)
}

//------------------------------------------------------------
// Routers that maps the routes to the endpoints for our system.
//-------------------------------------------------------------
func assignRoutes(router *mux.Router) *mux.Router {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc:", "pilot-management",
			"ts:", log.DefaultTimestampUTC,
			"caller:", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	service := impl.MakeServiceImpl(logger)

	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(EncodeErrorResponse),
		httpTransport.ServerErrorLogger(logger),
	}

	statusHandler := httpTransport.NewServer(
		MakeStatusEndpoint(service),
		DecodeStatusRequest,
		EncodeResponse,
	)

	listPilotsHandler := httpTransport.NewServer(
		MakeListPilotsEndpoint(service),
		DecodeListPilotsRequest,
		EncodeResponse,
		options...,
	)

	getPilotHandler := httpTransport.NewServer(
		MakeGetPilotEndpoint(service),
		DecodeGetPilotRequest,
		EncodeResponse,
		options...,
	)

	CreatePilotHandler := httpTransport.NewServer(
		MakeCreatePilotEndpoint(service),
		DecodeCreatePilotRequest,
		EncodeResponse,
		options...,
	)

	UpdatePilotHandler := httpTransport.NewServer(
		MakeUpdatePilotEndpoint(service),
		DecodeUpdatePilotRequest,
		EncodeResponse,
		options...,
	)

	DeletePilotHandler := httpTransport.NewServer(
		MakeDeletePilotEndpoint(service),
		DecodeDeletePilotRequest,
		EncodeResponse,
		options...,
	)

	ChangePilotStatusHandler := httpTransport.NewServer(
		MakeChangePilotStatusEndpoint(service),
		DecodeChangePilotStatusRequest,
		EncodeResponse,
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
