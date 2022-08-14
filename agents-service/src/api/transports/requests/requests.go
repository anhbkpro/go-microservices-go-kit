package requests

type InsertAgentPlayerRequest struct {
	AgentID  uint32 `json:"agent_id"`
	PlayerID uint32 `json:"player_id"`
}

type GetAgentByIdRequest struct {
	AgentID uint32
}
