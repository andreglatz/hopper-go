package main

import (
	"context"

	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driven/db/repositories"
	"github.com/andreglatz/hopper-go/internal/driving/http/handlers"
	"github.com/andreglatz/hopper-go/internal/settings"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := gin.Default()
	s := settings.New()

	en := en.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_translations.RegisterDefaultTranslations(v, translator)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.Database.URL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	repository := repositories.NewPostgresLinkRepository(conn)
	useCase := usecases.NewCreateShortLinkUseCase(repository)
	handler := handlers.NewCreateShortLinkHandler(translator, useCase)

	r.POST("/links", handler.Handle)

	return r.Run(":3000")
}
