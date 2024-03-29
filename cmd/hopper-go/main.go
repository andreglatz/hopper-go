package main

import (
	"github.com/andreglatz/hopper-go/cmd/hopper-go/setup"
	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driven/db/repositories"
	"github.com/andreglatz/hopper-go/internal/driving/http/handlers"
	"go.uber.org/fx"

	_ "github.com/lib/pq"
)

func main() {
	fx.New(
		fx.Provide(
			setup.NewHTTPServer,
			setup.GetDBConnection,
			setup.GetTranslator,
			setup.GetLogger,
			repositories.NewPostgresLinkRepository,
			usecases.NewCreateShortLinkUseCase,
			usecases.NewRedirectLinkUseCase,
			usecases.NewGetLinksUseCase,
			handlers.NewCreateShortLinkHandler,
			handlers.NewRedirectLinkHandler,
			handlers.NewGetLinksHandler,
		),
		fx.Invoke(setup.RegisterRoutes),
	).Run()
}
