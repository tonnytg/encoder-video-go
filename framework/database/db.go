package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/tonnytg/encoder-video-go/domain"

	//_ "gorm.io/driver/postgres"
	//_ "gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}
	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {

	var err error

	if d.Env != "test" {
		//d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		// normal sqlite file
		// d.Db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		// sqlite temporary in memory
		//d.Db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}


	if err != nil {
		return nil, err
	}

	if d.Debug {

		d.Db.LogMode(true)
		// 1 - Silent, 2 - Error, 3 - Warn, 4 - Info
		//d.Db.Logger.LogMode(4)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{}, )
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
