package SQL

import (
	"database/sql"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func InsertDataSet() {
	mlk := model.Product{ID: "100", Name: "Milk Carton", UnitValue: 2500, Units: 98}
	eg := model.Product{ID: "101", Name: "Egg Carton", UnitValue: 7800, Units: 75}
	flr := model.Product{ID: "102", Name: "Flour Bag", UnitValue: 3200, Units: 63}

	InsertProduct(mlk)
	InsertProduct(eg)
	InsertProduct(flr)
}
