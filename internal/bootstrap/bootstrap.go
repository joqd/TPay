package bootstrap

import (
	"fmt"

	"github.com/joqd/TPay/internal/adapter/config"
	"github.com/joqd/TPay/internal/adapter/server"
	"github.com/joqd/TPay/internal/adapter/server/handler"
)

func Start() {
	conf := config.Load()

	srv := server.New(
		handler.NewHealthHandler(),
	)

	srv.Run(fmt.Sprintf(":%d", conf.Server.Port))
}
