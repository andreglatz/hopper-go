package setup

import (
	"context"

	"github.com/andreglatz/hopper-go/internal/settings"
	"github.com/jackc/pgx/v5"
)

func GetDBConnection() *pgx.Conn {
	s := settings.New()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.Database.URL)
	if err != nil {
		panic(err)
	}

	return conn
}
