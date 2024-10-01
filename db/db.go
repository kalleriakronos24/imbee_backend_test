package database

import (
	"log"
	"time"

	"gorm.io/gorm/logger"

	"github.com/kalleriakronos24/imbee-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	SortDesc   string      `json:"sortDesc"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"rows"`
}

func GetDatabaseConnection() *gorm.DB {
	var db *gorm.DB

	loggerConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}

	if config.AppConfig.Environment == "PRODUCTION" {
		loggerConfig = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		}
	}

	db, err := gorm.Open(mysql.Open(config.AppConfig.DBUrl), loggerConfig)

	if con, _ := db.DB(); err != nil {
		log.Println("[INIT] failed connecting to MySQL")
		return nil
	} else {
		con.SetConnMaxLifetime(15 * time.Second)
		con.SetMaxOpenConns(90)
		con.SetMaxIdleConns(20)
	}

	return db
}

func DropUnusedColumns(dst interface{}) {

	db := GetDatabaseConnection()
	stmt := db.Statement
	err := stmt.Parse(dst)

	if err != nil {
		log.Println("[INIT] failed parse column on the database ", err.Error())
		return
	}
	fields := stmt.Schema.Fields
	columns, _ := db.Debug().Migrator().ColumnTypes(dst)

	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			err := db.Migrator().DropColumn(dst, columns[i].Name())
			if err != nil {
				log.Println("[INIT] failed drop column on the database ", err.Error())
				return
			}
		}
	}
}
