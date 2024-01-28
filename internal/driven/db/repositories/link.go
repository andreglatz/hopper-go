package repositories

import (
	"context"

	"github.com/andreglatz/hopper-go/internal/application/entities"
	"github.com/andreglatz/hopper-go/internal/driven/db/models"
	sql "github.com/andreglatz/hopper-go/tools/sqlc"
	"github.com/jackc/pgx/v5"
)

type LinkRepository interface {
	Create(*entities.Link) error
	Update(*entities.Link) error
	GetByShort(string) (*entities.Link, error)
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

func (r *PostgresLinkRepository) Create(link *entities.Link) error {
	result, err := r.db.CreateLink(context.Background(), sql.CreateLinkParams{Short: link.Short, Original: link.Original})
	if err != nil {
		return err
	}

	link.ID = uint(result.ID)

	return nil
}

func (r *PostgresLinkRepository) Update(link *entities.Link) error {
	params := sql.UpdateLinkParams{
		ID:       int32(link.ID),
		Short:    link.Short,
		Original: link.Original,
		Clicks:   int32(link.Clicks),
	}

	err := r.db.UpdateLink(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresLinkRepository) GetByShort(short string) (*entities.Link, error) {
	link, err := r.db.GetLinkByShort(context.Background(), short)
	if err != nil {
		return nil, err
	}

	return models.NewLink(link).ToEntity(), nil
}
