{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/event"
)

var log = event.Logger()

type Entity interface {
	TableName() string
}

type Entities []Entity

// entities List of database entities.
var entities Entities

func AddEntity(e Entity) {
	entities = append(entities, e)
}

// Truncate removes all data from tables without dropping them.
func (es Entities) Truncate() {
	for _, entity := range es {
		if err := Db().Debug().Exec("DELETE FROM ? WHERE 1", entity.TableName()).Error; err == nil {
			log.Printf("entity: truncated %s successfully", entity.TableName())
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("entity: truncated %s failed (%s)", entity.TableName(), err)
		}
	}
}

// Migrate migrates all database tables of registered entities.
func (es Entities) Migrate() {
	for _, entity := range es {
		if err := UnscopedDb().Debug().AutoMigrate(entity); err != nil {
			log.Printf("entity: migrate %s %s (waiting 1s)", entity.TableName(), err.Error())
			time.Sleep(time.Second)
			if err = UnscopedDb().Debug().AutoMigrate(entity); err != nil {
				panic(err)
			}
		}
	}
}

// Drop drops all database tables of registered entities.
func (es Entities) Drop() {
	Db().Exec("PRAGMA foreign_keys = OFF;")

	for _, entity := range es {
		if err := Db().Debug().Migrator().DropTable(&entity); err != nil {
			panic(err)
		}
	}

	Db().Exec("PRAGMA foreign_keys = ON;")
}

// MigrateDb creates all tables and inserts default entities as needed.
func MigrateDb() {
	entities.Migrate()

	CreateDefaultFixtures()
}

func DropEntities() {
	entities.Drop()
}

// CreateDefaultFixtures inserts default fixtures for test and production.
func CreateDefaultFixtures() {
	CreateDefaultUsers()
}

// ResetDatabase resets the database schema.
func ResetDatabase() error {
	start := time.Now()

	log.Print("entity: dropping existing tables")
	DropEntities()

	log.Print("entity: restoring default schema")
	MigrateDb()

	log.Printf("entity: database reset completed in %s", time.Since(start))

	return nil
}
