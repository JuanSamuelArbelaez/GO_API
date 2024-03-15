package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/Utils"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services/complementary"
)

func GetInventory() ([]model.Product, error) {
	return complementary.SelectAllProducts()
}

func GetInventorySize() (size int, e error) {
	return complementary.CountProducts()
}

func AddProduct(p model.ProductRequest) (id string, e error) {
	e = CheckProduct(p)
	if e != nil {
		return "", e
	}

	newId := ""
	e = Utils.GenerateId(&newId)

	if e != nil {
		return "", e
	} else {
		newProduct := model.Product{}
		newProduct.ID = newId
		model.MapRequest(&newProduct, &p)
		complementary.InsertProduct(newProduct)
		return newId, nil
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
		if units > p.Units {
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
		err := CheckProduct(model.ProductRequest{Name: p.Name, UnitValue: p.UnitValue, Units: p.Units})
		if err != nil {
			return err
		}
		return nil
	}
}

func ContainsProduct(ID string) (isContained bool, e error) {
	return complementary.ContainsProductByID(ID)
}

func GetProduct(ID string) (product model.Product, e error) {
	p, _ := complementary.SelectProductByID(ID)
	fmt.Println(p)
	return p, nil
}

func CheckProduct(product model.ProductRequest) (e error) {
	if product.Name == "" || product.UnitValue <= 0 || product.Units <= 0 {
		return fmt.Errorf("product is invalid")
	} else {
		return nil
	}
}
