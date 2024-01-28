package setup

import (
	"context"

	"github.com/andreglatz/hopper-go/internal/driving/http/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterRoutesParams struct {
	fx.In

	CreateShortLinkHandler *handlers.CreateShortLinkHandler
	RedirectLinkHandler    *handlers.RedirectLinkHandler
	GetLinksHandler        *handlers.GetLinksHandler
}

func RegisterRoutes(r *gin.Engine, handlers RegisterRoutesParams) {
	r.POST("/links", handlers.CreateShortLinkHandler.Handle)
	r.GET("/:short", handlers.RedirectLinkHandler.Handle)
	r.GET("/", handlers.GetLinksHandler.Handle)
}

func NewHTTPServer(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.Run(":3000"); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return r
}
