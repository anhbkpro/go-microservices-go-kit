package responses

import "github.com/anhbkpro/go-microservices-go-kit/src/api/entities"

type InsertAgentPlayerResponse struct {
	Err string `json:"error,omitempty"`
}

type GetAgentByIDResponse struct {
	Agent *entities.Agent `json:"agent,omitempty"`
	Err   string          `json:"error,omitempty"`
}
