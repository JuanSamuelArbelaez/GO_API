package SQL

import (
	"database/sql"
	"errors"
	"fmt"
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

func InsertProduct(p model.Product) {
	_, err := DB.Exec("INSERT INTO product (ID, Name, UnitValue, Units) VALUES (?, ?, ?, ?)", p.ID, p.Name, p.UnitValue, p.Units)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("nuevo producto insertado correctamente.")
}

func UpdateProduct(p model.Product) {
	_, err := DB.Exec("UPDATE product SET Name = ?, UnitValue = ?, Units = ? WHERE ID = ?", p.Name, p.UnitValue, p.Units, p.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("producto actualizado correctamente.")
}

func DeleteProduct(ID string) {
	_, err := DB.Exec("DELETE FROM product WHERE ID = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("producto eliminado correctamente.")
}

func SelectProductByID(ID string) (model.Product, error) {
	rows, err := DB.Query("SELECT * FROM product WHERE ID = ?", ID)
	if err != nil {
		return model.Product{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return model.Product{}, errors.New("producto no encontrado")
	}
	var product model.Product
	if err := rows.Scan(&product.ID, &product.Name, &product.UnitValue, &product.Units); err != nil {
		return model.Product{}, err
	}
	fmt.Println(product)
	return product, nil
}

func ContainsProductByID(ID string) (bool, error) {
	rows, err := DB.Query("SELECT 1 FROM product WHERE ID = ?", ID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func SelectAllProducts() ([]model.Product, error) {
	rows, err := DB.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.UnitValue, &product.Units); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func CountProducts() (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM product").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
