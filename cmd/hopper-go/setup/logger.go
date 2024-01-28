package setup

import (
	"github.com/andreglatz/hopper-go/internal/settings"
	"go.uber.org/zap"
)

func GetLogger() *zap.SugaredLogger {
	s := settings.Get()

	var logger *zap.Logger
	var err error

	if s.Application.Env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar
}
