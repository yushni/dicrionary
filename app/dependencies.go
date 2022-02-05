package app

import (
	"dictionary/app/facilities"
	"dictionary/db"
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
