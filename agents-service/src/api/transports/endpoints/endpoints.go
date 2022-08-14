package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/service"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/transports/requests"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/transports/responses"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	// ErrNoAgentID agentID parameter is missing. Respond 404.
	ErrNoAgentID = errors.New("agentID is required")
	// ErrAgentIDNotNumber AgentID was not a number. Respond 404
	ErrAgentIDNotNumber = errors.New("agentID is not a number")
)

func MakeInsertAgentPlayerEndpoint(srv service.AgentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.InsertAgentPlayerRequest)
		err := srv.InsertAgentPlayer(req.AgentID, req.PlayerID)
		if err != nil {
			return responses.InsertAgentPlayerResponse{Err: err.Error()}, nil
		}
		return responses.InsertAgentPlayerResponse{Err: ""}, nil
	}
}

func MakeGetAgentByIdRequestEndpoint(srv service.AgentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetAgentByIdRequest)
		v, err := srv.GetAgentByID(req.AgentID)
		if err != nil {
			return responses.GetAgentByIDResponse{Agent: nil, Err: err.Error()}, nil
		}
		return responses.GetAgentByIDResponse{Agent: v, Err: ""}, nil
	}
}

func DecodeInsertAgentPlayerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.InsertAgentPlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetAgentByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		return nil, ErrNoAgentID
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, ErrAgentIDNotNumber
	}
	return requests.GetAgentByIdRequest{AgentID: uint32(id)}, nil
}

func EncodeInsertAgentPlayerResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	res, ok := response.(responses.InsertAgentPlayerResponse)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return errors.New("error when casting response")
	}
	if res.Err != "" {
		rw.WriteHeader(http.StatusInternalServerError)
		return errors.New(res.Err)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("write was successful."))
	return nil
}

func EncodeGetAgentByIdResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	res, ok := response.(responses.GetAgentByIDResponse)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return errors.New("error when casting response")
	}
	if res.Err != "" {
		rw.WriteHeader(http.StatusInternalServerError)
		return errors.New(res.Err)
	}
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return err
	}
	rw.WriteHeader(http.StatusOK)
	return nil
}
