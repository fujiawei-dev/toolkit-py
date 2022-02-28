{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "database/sql/driver"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
//     "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/pkg/fs"
)

const (
	MySQL   = "mysql"
	MariaDB = "mariadb"
	SQLite  = "sqlite3"
	//Postgres = "postgres"
)

func (c *config) Db() *gorm.DB {
	if c.gm == nil {
		panic("config: database not connected")
	}

	return c.gm
}

// InitDb will initialize the database connection and schema.
func (c *config) InitDb() error {
	if err := c.connectDb(); err != nil {
		return err
	}

	entity.SetDbProvider(c)
	entity.MigrateDb()

	return nil
}

// CloseDb closes the db connection (if any).
func (c *config) CloseDb() error {
	if c.gm != nil {

		sqlDB, err := c.gm.DB()
		if err != nil || sqlDB == nil {
			return err
		}

		if err = sqlDB.Close(); err == nil {
			c.gm = nil
		} else {
			return err
		}
	}

	return nil
}

func (c *config) DatabaseType() string {
	switch strings.ToLower(c.settings.Database.DbType) {
	case MySQL, MariaDB:
		c.settings.Database.DbType = MySQL
	case SQLite, "sqlite", "test", "file", "":
		c.settings.Database.DbType = SQLite
	default:
		panic("config: unsupported database type " + c.settings.Database.DbType)
	}

	return c.settings.Database.DbType
}

func (c *config) DatabaseServer() string {
	if c.settings.Database.HostPort != "" {
		return c.settings.Database.HostPort
	}

	return "localhost"
}

// DatabaseHost the database server host.
func (c *config) DatabaseHost() string {
	if s := strings.Split(c.DatabaseServer(), ":"); len(s) > 0 {
		return s[0]
	}

	return c.DatabaseServer()
}

// DatabasePort the database server port.
func (c *config) DatabasePort() int {
	const defaultPort = 3306

	if s := strings.Split(c.DatabaseServer(), ":"); len(s) != 2 {
		return defaultPort
	} else if port, err := strconv.Atoi(s[1]); err != nil {
		return defaultPort
	} else if port < 1 || port > 65535 {
		return defaultPort
	} else {
		return port
	}
}

// DatabasePortString the database server port as string.
func (c *config) DatabasePortString() string {
	return strconv.Itoa(c.DatabasePort())
}

func (c *config) DatabaseUser() string {
	if c.settings.Database.Username != "" {
		return c.settings.Database.Username
	}

	return "root"
}

func (c *config) DatabasePassword() string {
	if c.settings.Database.Password != "" {
		return c.settings.Database.Password
	}

	return "root"
}

func (c *config) DatabaseName() string {
	if c.settings.Database.DbName != "" {
		return c.settings.Database.DbName
	}

	return c.AppName() + ".db"
}

func (c *config) DatabasePath() string {
	if c.settings.Database.DbPath != "" {
		return c.settings.Database.DbPath
	}

	return c.StoragePath()
}

func (c *config) DatabaseDialector() gorm.Dialector {
	switch c.DatabaseType() {
	case MySQL, MariaDB:
		return mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8"+
				"&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
				c.DatabaseUser(),
				c.DatabasePassword(),
				c.DatabaseServer(),
				c.DatabaseName(),
			))
	case SQLite:
		return sqlite.Open(fs.Join(c.DatabasePath(), c.DatabaseName()))
	default:
		panic("config: not currently supported")
	}
}

func (c *config) DatabaseConns() int {
	limit := (runtime.NumCPU() * 2) + 16

	if limit > 1024 {
		limit = 1024
	}

	return limit
}

func (c *config) DatabaseConnsIdle() int {
	limit := runtime.NumCPU() + 8

	if limit > c.DatabaseConns() {
		limit = c.DatabaseConns()
	}

	return limit
}

func (c *config) DatabaseLogger() logger.Interface {
	var prefix string

	if !c.DetachServer() {
		prefix = "\r\n"
	}

	return logger.New(log.New(c.LogWriter(), prefix, log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  c.settings.Database.LogLevel, // 1 silent 2 error 3 warn 4 info
		IgnoreRecordNotFoundError: false,
		Colorful:                  !c.DetachServer(),
	})
}

// connectDb establishes a database connection.
func (c *config) connectDb() error {
	gm, err := gorm.Open(c.DatabaseDialector(), &gorm.Config{
		Logger:                                   c.DatabaseLogger(),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil || gm == nil {
		return err
	}

	sqlDB, err := gm.DB()
	if err != nil || sqlDB == nil {
		return err
	}

	sqlDB.SetMaxOpenConns(c.DatabaseConns())
	sqlDB.SetMaxIdleConns(c.DatabaseConnsIdle())
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	c.gm = gm

	return nil
}

// ImportSQL imports a file to the currently configured database.
func (c *config) ImportSQL(path string) error {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	statements := strings.Split(string(contents), ";\n")
	q := c.Db().Unscoped()

	for _, stmt := range statements {
		// Skip empty lines and comments
		if len(stmt) < 3 || stmt[0] == '#' || stmt[0] == ';' {
			continue
		}

		var result struct{}

		q.Raw(stmt).Scan(&result)
	}

	return nil
}
