package db

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {
	godotenv.Load()
}

// DBInit create connection to database
func DBInit() (*gorm.DB, error) {

	var (
		USERNAME = os.Getenv("DB_USERNAME")
		PASSWORD = os.Getenv("DB_PASSWORD")
		HOST     = os.Getenv("DB_HOST")
		PORT     = os.Getenv("DB_PORT")
		DATABASE = os.Getenv("DB_DATABASE")

		mysqlCon = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			USERNAME,
			PASSWORD,
			HOST,
			PORT,
			DATABASE,
		)
	)
	var err error
	DB, err = gorm.Open("mysql", mysqlCon)

	if err != nil {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println("")
		errorOutput.Println("!!! Warning")
		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", mysqlCon))
		errorOutput.Println("")

		log.Println(err)
	} else {

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)

		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", mysqlCon))
		successOutput.Println("")
	}

	fmt.Println("Connection is created")

	return DB, err
}

func GetConnection() *gorm.DB {
	if DB == nil {
		fmt.Println("Initialize database")
		DB, _ = DBInit()
	} else {
		fmt.Println("Get connection database")
	}
	return DB
}
