package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
)

func checkProduct(product model.Product) (e error) {
	if getProductID(product) == "" || getProductName(product) == "" || getProductUnitValue(product) <= 0 || getProductUnits(product) <= 0 {
		return fmt.Errorf("Product is invalid")
	} else {
		return nil
	}
}

func setProductID(product model.Product, ID string) (e error) {
	if ID == "" {
		return fmt.Errorf("must provide a ID")
	}
	product.ID = ID
	return nil
}

func getProductID(product model.Product) (ID string) {
	return product.ID
}

func setProductName(product model.Product, name string) (e error) {
	if name == "" {
		return fmt.Errorf("must provide a name")
	}
	product.Name = name
	return nil
}

func getProductName(product model.Product) (name string) {
	return product.Name
}

func setProductUnitValue(product model.Product, value float32) (e error) {
	if value <= 0 {
		return fmt.Errorf("value must be positive")
	}
	product.UnitValue = value
	return nil
}

func getProductUnitValue(product model.Product) (unitValue float32) {
	return product.UnitValue
}

func setProductUnits(product model.Product, units int) (e error) {
	if units <= 0 {
		return fmt.Errorf("units must be positive")
	}
	product.Units = units
	return nil
}

func getProductUnits(product model.Product) (units int) {
	return product.Units
}
