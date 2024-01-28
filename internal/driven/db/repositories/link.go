package repositories

import (
	"context"

	"github.com/andreglatz/hopper-go/internal/application/entities"
	sql "github.com/andreglatz/hopper-go/tools/sqlc"
	"github.com/jackc/pgx/v5"
)

type LinkRepository interface {
	Save(*entities.Link) error
}

type PostgresLinkRepository struct {
	db *sql.Queries
}

func NewPostgresLinkRepository(conn *pgx.Conn) LinkRepository {
	db := sql.New(conn)

	return &PostgresLinkRepository{
		db: db,
	}
}

func (r *PostgresLinkRepository) Save(link *entities.Link) error {
	result, err := r.db.CreateLink(context.Background(), sql.CreateLinkParams{Short: link.Short, Original: link.Original})
	if err != nil {
		return err
	}

	link.ID = uint(result.ID)

	return nil
}
