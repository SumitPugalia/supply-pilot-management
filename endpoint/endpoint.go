package endpoint

import (
	"context"

	"gitlab.intelligentb.com/cafu/supply/pilot-management/domain"

	"github.com/go-kit/kit/endpoint"
)

func MakeListPilotsEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := s.ListPilots()
		if err != nil {
			return nil, err
		}
		pilots := make([]PilotView, 0)
		for _, pilot := range v {
			pilots = append(pilots, toPilotView(pilot))
		}
		return Response{Data: pilots}, nil
	}
}

func MakeGetPilotEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPilotRequest)
		pilot, err := s.GetPilot(req.Id)
		if err != nil {
			return nil, err
		}
		return Response{Data: toPilotView(pilot)}, nil
	}
}

func MakeCreatePilotEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePilotRequest)
		pilot, err := s.CreatePilot(domain.CreatePilotParams(req))
		if err != nil {
			return nil, err
		}
		return Response{Data: toPilotView(pilot)}, nil
	}
}

func MakeUpdatePilotEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePilotRequest)
		pilot, err := s.UpdatePilot(domain.UpdatePilotParams(req))
		if err != nil {
			return nil, err
		}
		return Response{Data: toPilotView(pilot)}, nil
	}
}

func MakeChangeStatePilotEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeStatePilotRequest)
		pilot, err := s.ChangeStatePilot(req.Id, req.State)
		if err != nil {
			return nil, err
		}
		return Response{Data: toPilotView(pilot)}, nil
	}
}

func MakeDeletePilotEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePilotRequest)
		err := s.DeletePilot(req.Id)
		if err != nil {
			return nil, err
		}
		return Response{Data: nil}, nil
	}
}
