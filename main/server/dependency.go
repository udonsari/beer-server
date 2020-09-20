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
