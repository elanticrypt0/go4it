package go4it

import (
	"fmt"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect2DB(app *App, connName string) {
	// check if config exists
	dbConfig, exists := app.Config.DB[connName]
	if exists {
		newConn := DBActive{
			Name: connName,
			Conn: Connect2Engine(&dbConfig),
		}
		app.DB.Actives = append(app.DB.Actives, newConn)
	}
	if len(app.DB.Actives) == 1 {
		app.DB.SetPrimaryDB(0)
	}

}

func Connect2DBOnly(dbo *DBOnly, connName string) {
	// check if config exists
	dbConfig, exists := dbo.Config.Connection[connName]
	if exists {
		newConn := DBActive{
			Name: connName,
			Conn: Connect2Engine(&dbConfig),
		}
		dbo.DB.Actives = append(dbo.DB.Actives, newConn)
	}
	if len(dbo.DB.Actives) == 1 {
		dbo.DB.SetPrimaryDB(0)
	}
}

func Connect2Engine(dbconfig *DatabaseConfig) *gorm.DB {

	var conn *gorm.DB

	switch dbconfig.Engine {
	case "postgres":
		conn = DbConnectPostgres(dbconfig)
	case "mysql":
		conn = DbConnectMySql(dbconfig)
	case "sqlite":
		conn = DbConnectSqlite(dbconfig)
	}
	return conn
}

// Engines:

func DbConnectPostgres(dbconfig *DatabaseConfig) *gorm.DB {

	const dns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable "
	dnsConfig := fmt.Sprintf(dns, dbconfig.Host, dbconfig.User, dbconfig.Password, dbconfig.DBName, dbconfig.Port)
	// connect to gorn
	conn, err := gorm.Open(postgres.Open(dnsConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database Postgres")
	}
	return conn
}

func DbConnectMySql(dbconfig *DatabaseConfig) *gorm.DB {

	const dns = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dnsConfig := fmt.Sprintf(dns, dbconfig.User, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.DBName)
	// connect to gorn
	conn, err := gorm.Open(mysql.Open(dnsConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database MySQL")
	}
	return conn
}

// obtiene una conexión hacia la base de datos sqlite que se encuentra en la carpeta "libs" con el nombre que se le pase como parámetro
func DbConnectSqlite(dbconfig *DatabaseConfig) *gorm.DB {
	dbname := strings.ToLower(dbconfig.DBName)
	// gorm create sqlite db
	conn, err := gorm.Open(sqlite.Open(dbname+".db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database SQLITE")
	}
	return conn
}
