package repositories

import (
	"context"

	"github.com/andreglatz/hopper-go/internal/application/entities"
	"github.com/andreglatz/hopper-go/internal/driven/db/models"
	sql "github.com/andreglatz/hopper-go/tools/sqlc"
	"github.com/jackc/pgx/v5"
)

type GetLinksParams struct {
	Offset   int32
	Limit    int32
	Short    string
	Original string
}

type LinkRepository interface {
	Create(*entities.Link) error
	Update(*entities.Link) error
	GetByShort(string) (*entities.Link, error)
	GetLinks(GetLinksParams) ([]*entities.Link, uint, error)
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

func (r *PostgresLinkRepository) GetLinks(params GetLinksParams) ([]*entities.Link, uint, error) {
	dbParams := sql.GetLinksParams{
		Limit:    params.Limit,
		Offset:   params.Offset,
		Short:    params.Short,
		Original: params.Original,
	}

	links, err := r.db.GetLinks(context.Background(), dbParams)
	if err != nil {
		return nil, 0, err
	}

	var entitiesLinks []*entities.Link
	for _, link := range links {
		entitiesLinks = append(entitiesLinks, models.NewLink(link).ToEntity())
	}

	count, err := r.count(params)
	if err != nil {
		return nil, 0, err
	}

	return entitiesLinks, count, nil
}

func (r *PostgresLinkRepository) count(params GetLinksParams) (uint, error) {
	dbParams := sql.GetLinksCountParams{
		Short:    params.Short,
		Original: params.Original,
	}

	count, err := r.db.GetLinksCount(context.Background(), dbParams)
	if err != nil {
		return 0, err
	}

	return uint(count), nil
}
