package server

import (
  "context"
  "fmt"
  "net/http"

  "github.com/paarig/goyagi/pkg/health"
  "github.com/paarig/goyagi/pkg/signals"
  "github.com/labstack/echo"
  "github.com/lob/logger-go"
)

func New() *http.Server {
  log := logger.New()

  e := echo.New()

  health.RegisterRoutes(e)

  srv := &http.Server{
    Addr: fmt.Sprintf(":%d", 3000),
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
