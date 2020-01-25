package endpoint

//------------------------------------------------------------
// This file contains controller for the endpoints.
//-------------------------------------------------------------
import (
	"context"

	"pilot-management/service"

	"github.com/go-kit/kit/endpoint"
)

//------------------------------------------------------------
// All the controllers mapped to the routes
//-------------------------------------------------------------
func MakeStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return Response{Data: "Success"}, nil
	}
}

func MakeListPilotsEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ListPilotsRequest)
		pilots, totalEntries, totalPages, err := s.ListPilots(service.ListPilotParams(req))

		if err != nil {
			return nil, err
		}
		return Response{
			Data:       ToPilotViews(pilots),
			Pagination: PaginationData(totalEntries, totalPages),
		}, nil
	}
}

func MakeGetPilotEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPilotRequest)
		pilot, err := s.GetPilot(req.Id)
		if err != nil {
			return nil, err
		}
		return Response{Data: ToPilotView(pilot)}, nil
	}
}

func MakeCreatePilotEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePilotRequest)
		pilot, err := s.CreatePilot(service.CreatePilotParams(req))
		if err != nil {
			return nil, err
		}
		return Response{Data: ToPilotView(pilot)}, nil
	}
}

func MakeUpdatePilotEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePilotRequest)
		pilot, err := s.UpdatePilot(req.Id, update_pilot_params(req))
		if err != nil {
			return nil, err
		}
		return Response{Data: ToPilotView(pilot)}, nil
	}
}

func MakeChangePilotStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangePilotStatusRequest)
		pilot, err := s.ChangePilotStatus(req.Id, req.Status)
		if err != nil {
			return nil, err
		}
		return Response{Data: ToPilotView(pilot)}, nil
	}
}

func MakeDeletePilotEndpoint(s service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePilotRequest)
		err := s.DeletePilot(req.Id)
		if err != nil {
			return nil, err
		}
		return Response{Data: nil}, nil
	}
}

//------------------------------------------------------------
// Internal helper function
//-------------------------------------------------------------
func update_pilot_params(req UpdatePilotRequest) service.UpdatePilotParams {
	return service.UpdatePilotParams{
		UserId:     req.UserId,
		CodeName:   req.CodeName,
		SupplierId: req.SupplierId,
		MarketId:   req.MarketId,
		ServiceId:  req.ServiceId,
	}
}
