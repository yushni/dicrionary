package app

import (
	"dictionary/db"
	"dictionary/facilities"
)

type Migrate interface {
	Run() error
}

type Dependency struct {
	DBMigrate Migrate
}

func NewDependency(conf *facilities.Config) Dependency {
	migrate := db.NewMigrate(conf.DB)

	dep := Dependency{
		DBMigrate: migrate,
	}

	return dep
}
