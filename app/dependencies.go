package app

import (
	"dictionary/app/facilities"
	"dictionary/migration"
)

type Migrate interface {
	Run() error
}

type Dependency struct {
	DBMigrate Migrate
}

func NewDependency(conf facilities.Config) Dependency {
	migrate := migration.NewMigrate(conf.DB)

	dep := Dependency{
		DBMigrate: migrate,
	}

	return dep
}
