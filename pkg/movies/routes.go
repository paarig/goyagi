package movies

import (
  "github.com/paarig/goyagi/pkg/application"
  "github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo, app application.App) {
  g := e.Group("/movies")

  h := handler{app}

  g.GET("", h.listHandler)
  g.GET("/:id", h.retrieveHandler)
}
