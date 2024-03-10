package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/SQL"
	"github.com/JuanSamuelArbelaez/GO_API/model"
)

func GetInventory() ([]model.Product, error) {
	return SQL.SelectAllProducts()
}

func GetInventorySize() (size int, e error) {
	return SQL.CountProducts()
}

func AddProduct(product model.Product) (e error) {
	if e := CheckProduct(product); e != nil {
		return e
	}

	if contained, e := ContainsProduct(product.ID); e != nil {
		return e
	} else if contained {
		return fmt.Errorf("inventory already contains product '%s'", product.Name)
	} else {
		SQL.InsertProduct(product)
		return nil
	}
}

func RemoveProduct(ID string) (e error) {
	if contained, e := ContainsProduct(ID); e != nil {
		return e
	} else if !contained {
		return fmt.Errorf("product with ID:'%s' not found", ID)
	} else {
		return nil
	}
}

func SellProduct(ID string, units int) (total float32, e error) {
	if p, e := GetProduct(ID); e != nil {
		return 0, e
	} else {
		if units > GetProductUnits(p) {
			return 0, fmt.Errorf("not enough units to sell (%d). currently available: %d", units, p.Units)
		}
		/*
			if e := SetProductUnits(p, GetProductUnits(p)-units); e != nil {
				return 0, e
			}
		*/
		return float32(units) * p.UnitValue, nil
	}
}

func AddProductUnits(ID string, units int) (e error) {
	if p, e := GetProduct(ID); e != nil {
		return e
	} else {
		if units <= 0 {
			return fmt.Errorf("cannot add 0 or negative units")
		}
		/*
			if e := SetProductUnits(p, GetProductUnits(p)+units); e != nil {
				return e
			}
		*/
		CheckProduct(p)
		return nil
	}
}

func ContainsProduct(ID string) (isContained bool, e error) {
	return SQL.ContainsProductByID(ID)
}

func GetProduct(ID string) (product model.Product, e error) {
	p, _ := SQL.SelectProductByID(ID)
	fmt.Println(p)
	return p, nil
}

func CheckProduct(product model.Product) (e error) {
	if GetProductID(product) == "" || GetProductName(product) == "" || GetProductUnitValue(product) <= 0 || GetProductUnits(product) <= 0 {
		return fmt.Errorf("product is invalid")
	} else {
		return nil
	}
}

func GetProductID(product model.Product) (ID string) {
	return product.ID
}

func GetProductName(product model.Product) (name string) {
	return product.Name
}

func GetProductUnitValue(product model.Product) (unitValue float32) {
	return product.UnitValue
}

func GetProductUnits(product model.Product) (units int) {
	return product.Units
}
