package service

import (
	"github.com/anhbkpro/go-microservices-go-kit/src/api/entities"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/repositories"
)

// AgentsService Service Interface
type AgentsService interface {
	InsertAgentPlayer(agentID, playerID uint32) error
	GetAgentByID(agentID uint32) (*entities.Agent, error)
}

// AgentsServiceImpl Service Implementation
type AgentsServiceImpl struct {
	Repo *repositories.MariaDBAgentsRepository
}

func (srv AgentsServiceImpl) InsertAgentPlayer(agentID, playerID uint32) error {
	err := srv.Repo.InsertAgentPlayer(agentID, playerID)
	return err
}

func (srv AgentsServiceImpl) GetAgentByID(agentID uint32) (*entities.Agent, error) {
	agent, err := srv.Repo.GetAgentByID(agentID)
	return agent, err
}

// ServiceMiddleware is a chainable behavior modifier for AgentsService
type ServiceMiddleware func(AgentsService) AgentsServiceImpl
