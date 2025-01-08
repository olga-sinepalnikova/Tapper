package storage

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateStructs(db *gorm.DB, log *logrus.Logger) {
	generator := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	generator.UseDB(db)

	generator.ApplyBasic(generator.GenerateAllTable()...)

	generator.Execute()
	log.Debug("Models created")
}
