package application

import (
  "github.com/paarig/goyagi/pkg/config"
  "github.com/paarig/goyagi/pkg/database"
  "github.com/go-pg/pg"
  "github.com/pkg/errors"
)

type App struct {
  Config config.Config
  DB     *pg.DB
}

func New() (App, error) {
  cfg := config.New()

  db, err := database.New(cfg)
  if err != nil {
    return App{}, errors.Wrap(err, "application")
  }

  return App{cfg, db}, nil
}
