package SQL

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed load of .env file")
	}

	dbHost := os.Getenv("DB_host")
	dbPortStr := os.Getenv("DB_port")
	dbUser := os.Getenv("DB_user")
	dbPassword := os.Getenv("DB_pswd")
	dbName := os.Getenv("DB_name")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPortStr, dbName))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	err = DB.Ping()

	if err != nil {
		log.Fatalln(err)
	}
}

//func InsertDataSet() {
//	mlk := model.Product{ID: "100", Name: "Milk Carton", UnitValue: 2500, Units: 98}
//	eg := model.Product{ID: "101", Name: "Egg Carton", UnitValue: 7800, Units: 75}
//	flr := model.Product{ID: "102", Name: "Flour Bag", UnitValue: 3200, Units: 63}
//
//	complementary.InsertProduct(mlk)
//	complementary.InsertProduct(eg)
//	complementary.InsertProduct(flr)
//}
