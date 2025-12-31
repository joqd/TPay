package server

import (
	"log"

	"github.com/joqd/TPay/internal/adapter/server/handler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e        *echo.Echo
	handlers []handler.Handler
}

func New(handlers ...handler.Handler) *Server {
	e := echo.New()

	for _, h := range handlers {
		h.Register(e)
	}
	log.Println("handlers registered successfully")

	return &Server{
		e:        e,
		handlers: handlers,
	}
}

func (s *Server) Use(mw ...echo.MiddlewareFunc) {
	if len(mw) > 0 {
		s.e.Use(mw...)
		log.Println("middlewares registered successfully")
	}
}

func (s *Server) Run(addr string) error {
	log.Printf("server running on %s", addr)
	return s.e.Start(addr)
}
