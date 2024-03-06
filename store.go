package main

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services"
)

func main() {
	store := model.Store{Inventory: make(map[string]model.Product)}

	mlk := model.Product{ID: "100", Name: "Milk Carton", UnitValue: 2500, Units: 98}
	eg := model.Product{ID: "101", Name: "Egg Carton", UnitValue: 7800, Units: 75}
	flr := model.Product{ID: "102", Name: "Flour Bag", UnitValue: 3200, Units: 63}

	services.AddProduct(store, mlk)
	services.AddProduct(store, eg)
	services.AddProduct(store, flr)

	services.RemoveProduct(store, eg.ID)

	fmt.Println(services.GetInventorySize(store))

	fmt.Println(services.ContainsProduct(store, "100"))

	fmt.Println(services.GetProduct(store, "100"))
}
