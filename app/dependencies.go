package app

import (
	"dictionary/app/facilities"
	"dictionary/migration"
)

type IMigrate interface {
	Run() error
}

type Dependency struct {
	DBMigrate IMigrate
}

func NewDependency(conf facilities.Config) Dependency {
	migrate := migration.NewMigrate(conf.DB)

	dep := Dependency{
		DBMigrate: migrate,
	}

	return dep
}
