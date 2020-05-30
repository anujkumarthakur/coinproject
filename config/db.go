package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//"strconv"
	_ "github.com/go-sql-driver/mysql"
)

//DB db object
var DB *sql.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

//BuildDBConfig builds db config object from environment variables
func BuildDBConfig() DBConfig {

	//port, _ := strconv.ParseInt(os.Getenv("DBPORT"), 10, 0)

	dbConfig := DBConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		DBName:   os.Getenv("DBNAME"),
		Password: os.Getenv("DBPASS"),
	}
	return dbConfig
}

//DbURL get db connection string
// username:password@tcp(127.0.0.1:3306)/test")

func (dbConfig DBConfig) DbURL() string {
	return fmt.Sprintf(dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName)
}

func openDB() {
	var err error
	DB, err = sql.Open("mysql", BuildDBConfig().DbURL())
	if err != nil {
		log.Println("db connect error")
		log.Println(err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

//GetDB get db object
func GetDB() *sql.DB {
	if DB == nil {
		openDB()
	}

	return DB
}
