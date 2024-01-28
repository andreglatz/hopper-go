package main

import (
	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driven/repositories"
	"github.com/andreglatz/hopper-go/internal/driving/http/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := gin.Default()

	en := en.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_translations.RegisterDefaultTranslations(v, translator)
	}

	repository := repositories.NewInMemoryLinkRepository()
	useCase := usecases.NewCreateShortLinkUseCase(repository)
	handler := handlers.NewCreateShortLinkHandler(translator, useCase)

	r.POST("/links", handler.Handle)

	return r.Run(":3000")
}
