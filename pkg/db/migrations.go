package db

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

func Migrate(log *zap.Logger, fs *embed.FS, dbUrl string) {
	source, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Error("failed to read migrations source", zap.Error(err))

		return
	}

	instance, err := migrate.NewWithSourceInstance("iofs", source, makeMigrateUrl(dbUrl))
	if err != nil {
		log.Error("failed to initialization the migrate instance", zap.Error(err))

		return
	}

	err = instance.Up()

	switch err {
	case nil:
		log.Info("the migration schema successfully upgraded!")
	case migrate.ErrNoChange:
		log.Info("the migration schema not changed")
	default:
		log.Error("could not apply the migration schema", zap.Error(err))
	}
}
