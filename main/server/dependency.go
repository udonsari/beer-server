package server

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Dependency struct{}

const (
	envKeyMySQLDataSourceName = "MYSQL_DATA_SOURCE_NAME"
	envKeyBeerCacheDuration   = "BEER_CACHE_DURATION"
	envKeyHost                = "HOST"
	envKeyPort                = "PORT"
	envKeyServerEnv           = "SERVER_ENV"
	envKeyDevToken            = "DEV_TOKEN"

	mysqlMaxConn     = 30
	mysqlMaxIdleConn = 5
)

func NewDependency() Dependency {
	return Dependency{}
}

func (d *Dependency) MysqlDB(logMode bool) *gorm.DB {
	connectionString := d.getEnvOrExit(envKeyMySQLDataSourceName)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed to open mysqldb err: %v", err)
	}

	db.DB().SetMaxOpenConns(mysqlMaxConn)
	db.DB().SetMaxIdleConns(mysqlMaxIdleConn)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.LogMode(logMode)

	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}

	return db
}

func (d *Dependency) BeerCacheDuration() int64 {
	return d.getInt64Env(envKeyBeerCacheDuration)
}

func (d *Dependency) Host() string {
	return d.getEnvOrExit(envKeyHost)
}

func (d *Dependency) PortStr() string {
	return d.getEnvOrExit(envKeyPort)
}

func (d *Dependency) PortInt() int64 {
	return d.getInt64Env(envKeyPort)
}

func (d *Dependency) ServerEnv() string {
	return d.getEnvOrExit(envKeyServerEnv)
}

func (d *Dependency) DevToken() string {
	val, ok := os.LookupEnv(envKeyDevToken)
	if !ok {
		if d.ServerEnv() == "prod" {
			log.Fatalf("In prod env, dev token should be set")
		}
		val = "TEST_DEV_TOKEN"
	}
	return val
}

func (d *Dependency) getEnvOrExit(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("failed to get environment variable. key: %s", key)
	}
	return val
}

func (d *Dependency) getInt64Env(key string) int64 {
	valStr := d.getEnvOrExit(key)
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		log.Fatalf("failed to parse val to int. key: %s, val: %s", key, valStr)
	}
	return val
}
