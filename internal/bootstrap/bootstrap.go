package bootstrap

import (
	"log"

	"github.com/joqd/TPay/internal/adapter/config"
)


func Start() {
	conf := config.Load()
	log.Printf("server running on %d", conf.Server.Port)
}