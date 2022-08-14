package repositories

import (
	"database/sql"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDBAgentsRepository struct {
	db *sql.DB
}

func NewMariaDBAgentsRepository() *MariaDBAgentsRepository {
	db, err := sql.Open("mysql", "root:Password123!@tcp(localhost:3306)/agentsdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MariaDBAgentsRepository{
		db,
	}

	return repo
}

func (repo *MariaDBAgentsRepository) Close() {
	repo.db.Close()
}

func (repo *MariaDBAgentsRepository) InsertAgentPlayer(agentID, playerID uint32) error {
	rows, err := repo.db.Query("INSERT INTO manager_player(manager_id, player_id) VALUES (?, ?);", agentID, playerID)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (repo *MariaDBAgentsRepository) GetAgentByID(agentID uint32) (*entities.Agent, error) {
	m := &entities.Agent{}
	row := repo.db.QueryRow("SELECT id, manager, account FROM manager WHERE id=?", agentID)
	err := row.Scan(&m.ID, &m.Name, &m.Account)
	if err != nil {
		return nil, err
	}
	return m, nil
}
