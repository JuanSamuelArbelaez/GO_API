package services

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
)

func GetInventory(store model.Store) (inventory map[string]model.Product) {
	return store.Inventory
}

func GetInventorySize(store model.Store) (size int, e error) {
	if e := CheckInventory(store); e != nil {
		return -1, e
	} else {
		return len(GetInventory(store)), nil
	}
}

func AddProduct(store model.Store, product model.Product) (e error) {
	if e := CheckProduct(product); e != nil {
		return e
	}

	if contained, e := ContainsProduct(store, product.ID); e != nil {
		return e
	} else if contained {
		return fmt.Errorf("inventory already contains product '%s'", product.Name)
	} else {
		GetInventory(store)[product.ID] = product
		return nil
	}
}

func RemoveProduct(store model.Store, ID string) (e error) {
	if contained, e := ContainsProduct(store, ID); e != nil {
		return e
	} else if !contained {
		return fmt.Errorf("product with ID:'%s' not found", ID)
	} else {
		delete(GetInventory(store), ID)
		return nil
	}
}

func SellProduct(store model.Store, ID string, units int) (total float32, e error) {
	er := CheckInventory(store)
	if er != nil {
		return 0, er
	}

	if p, e := GetProduct(store, ID); e != nil {
		return 0, e
	} else {
		if units > GetProductUnits(p) {
			return 0, fmt.Errorf("not enough units to sell (%d). currently available: %d", units, p.Units)
		}
		if e := SetProductUnits(p, GetProductUnits(p)-units); e != nil {
			return 0, e
		}
		return float32(units) * p.UnitValue, nil
	}
}

func AddProductUnits(store model.Store, ID string, units int) (e error) {
	er := CheckInventory(store)
	if er != nil {
		return er
	}

	if p, e := GetProduct(store, ID); e != nil {
		return e
	} else {
		if units <= 0 {
			return fmt.Errorf("cannot add 0 or negative units")
		}
		if e := SetProductUnits(p, GetProductUnits(p)+units); e != nil {
			return e
		}
		return nil
	}
}

func CheckInventory(store model.Store) (e error) {
	if GetInventory(store) == nil {
		return fmt.Errorf("inventory is nil")
	} else {
		return nil
	}
}

func ContainsProduct(store model.Store, ID string) (isContained bool, e error) {
	if e := CheckInventory(store); e != nil {
		return false, e
	} else {
		_, found := GetInventory(store)[ID]
		return found, nil
	}
}

func GetProduct(store model.Store, ID string) (product model.Product, e error) {
	if p, found := GetInventory(store)[ID]; found {
		return p, nil
	}
	return model.Product{}, fmt.Errorf("product not found")
}
