package server

import (
  "context"
  "fmt"
  "net/http"

  "github.com/paarig/goyagi/pkg/application"
  "github.com/paarig/goyagi/pkg/movies"  
  "github.com/paarig/goyagi/pkg/health"
  "github.com/paarig/goyagi/pkg/signals"
  "github.com/labstack/echo"
  "github.com/lob/logger-go"
)

func New(app application.App) *http.Server {
  log := logger.New()

  e := echo.New()

  health.RegisterRoutes(e)

  movies.RegisterRoutes(e, app)

  srv := &http.Server{
    Addr: fmt.Sprintf(":%d", app.Config.Port),
    Handler: e,
  }

  graceful := signals.Setup()

  go func() {
    <-graceful
    err := srv.Shutdown(context.Background())
    if err != nil {
      log.Err(err).Error("server shutdown")
    }
  }()


  return srv
}
