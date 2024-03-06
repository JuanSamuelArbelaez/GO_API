package services

import (
    "fmt"
    "github.com/JuanSamuelArbelaez/GO_API/model"
)

func getInventory(store model.Store) (inventory map[string]model.Product) {
    return store.Inventory
}

func getInventorySize(store model.Store) (size int, e error) {
    if e := checkInventory(store); e != nil {
        return -1, e
    } else {
        return len(getInventory(store)), nil
    }
}

func addProduct(store model.Store, product model.Product) (e error) {
    if e:=checkProduct(product); e != nil {
        return e
    }

    if contained, e := containsProduct(store, product.ID); e!=nil {
        return e
    } else if contained {
        return fmt.Errorf("Inventory already contains product '%s'", product.Name)
    } else {
        getInventory(store)[product.ID] = product
        return nil
    }
}

func removeProduct(store model.Store, ID string) (e error) {
    if contained, e := containsProduct(store, ID); e!=nil {
        return e
    } else if !contained {
        return fmt.Errorf("Product with ID:'%s' not found", ID)
    } else {
        delete(getInventory(store), ID)
        return nil
    }
}

func sellProduct(store model.Store, ID string, units int) (total float32, e error) {
    er := checkInventory(store)
    if er != nil {
        return 0, er
    }

    if p, e := getProduct(store, ID); e != nil {
        return 0, e
    } else {
        if units <= 0 {
            return 0, fmt.Errorf("units cannot be negative or 0")
        }
        if units > p.Units {
            return 0, fmt.Errorf("not enough units to sell (%d). currently available: %d", units, p.Units)
        }
        p.Units -= units
        return float32(units) * p.UnitValue, nil
    }
}

func addProductUnits(store model.Store, ID string, units int) (e error) {
    er := checkInventory(store)
    if er != nil {
        return er
    }

    if p, e := getProduct(store, ID); e != nil {
        return e
    } else {
        p.Units += units
        return nil
    }
}

func checkInventory(store model.Store) (e error) {
    if getInventory(store) == nil {
        return fmt.Errorf("inventory is nil")
    } else {
        return nil
    }
}

func containsProduct(store model.Store, ID string) (isContained bool, e error) {
    if e := checkInventory(store); e != nil {
        return false, e
    } else {
        _, found := getInventory(store)[ID]
        return found, nil
    }
}

func getProduct(store model.Store, ID string) (product model.Product, e error) {
    if p, found := getInventory(store)[ID]; found {
        return p, nil
    }
    return model.Product{}, fmt.Errorf("product not found")
}