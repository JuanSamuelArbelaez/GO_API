package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
)

func CheckProduct(product model.Product) (e error) {
	if GetProductID(product) == "" || GetProductName(product) == "" || GetProductUnitValue(product) <= 0 || GetProductUnits(product) <= 0 {
		return fmt.Errorf("product is invalid")
	} else {
		return nil
	}
}

func SetProductID(product model.Product, ID string) (e error) {
	if ID == "" {
		return fmt.Errorf("must provide a ID")
	}
	product.ID = ID
	return nil
}

func GetProductID(product model.Product) (ID string) {
	return product.ID
}

func SetProductName(product model.Product, name string) (e error) {
	if name == "" {
		return fmt.Errorf("must provide a name")
	}
	product.Name = name
	return nil
}

func GetProductName(product model.Product) (name string) {
	return product.Name
}

func SetProductUnitValue(product model.Product, value float32) (e error) {
	if value <= 0 {
		return fmt.Errorf("value must be positive")
	}
	product.UnitValue = value
	return nil
}

func GetProductUnitValue(product model.Product) (unitValue float32) {
	return product.UnitValue
}

func SetProductUnits(product model.Product, units int) (e error) {
	if units <= 0 {
		return fmt.Errorf("units must be positive")
	}
	product.Units = units
	return nil
}

func GetProductUnits(product model.Product) (units int) {
	return product.Units
}
