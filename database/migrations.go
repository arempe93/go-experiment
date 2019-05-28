package database

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/arempe93/experiment/models"
)

type Migration struct {
	ID uint `gorm:"primary_key"`
}

var m *gormigrate.Gormigrate
var initM sync.Once

var migrations = []*gormigrate.Migration{
	{
		ID: "20190522211930",
		Migrate: func(tx *gorm.DB) error {
			type Audit struct {
				TestColumn int
			}

			return tx.AutoMigrate(&Audit{}).Error
		},
	},
	{
		ID: "20190522223357",
		Migrate: func(tx *gorm.DB) error {
			return tx.Table("audits").DropColumn("test_column").Error
		},
	},
}

func Migrate() {
	initialize()

	if err := m.Migrate(); err != nil {
		log.Fatal("migration err:", err)
	}

	log.Println("Migrations successful")
}

func EnsureSchema() {
	if len(migrations) == 0 {
		return
	}

	lastMigration := migrations[len(migrations)-1]

	if Instance().First(&Migration{}, lastMigration.ID).RecordNotFound() {
		log.Fatal("Migrations are pending!")
	}
}

func InitSchema(tx *gorm.DB) (err error) {
	err = tx.AutoMigrate(
		&models.Audit{},
	).Error

	return
}

func initialize() {
	initM.Do(func() {
		m = gormigrate.New(Instance(), gormigrate.DefaultOptions, migrations)
		m.InitSchema(InitSchema)
	})
}
