package services

import (
    "module github.com/JuanSamuelArbelaez/GO_API/model"
)

func (s *Store) getInventory() (inventory map[string]*Product) {
    return s.Inventory
}

func (s *Store) getInventorySize() (size int, e error) {
    if e := s.checkInventory(); e != nil {
        return -1, e
    } else {
        return len(s.getInventory()), nil
    }
}

func (s *Store) addProduct(product *Product) (e error) {
    if e:=product.checkProduct(); e != nil {
        return e
    }

    if contained, e := s.containsProduct(product.ID); e!=nil {
        return e
    } else if contained {
        return fmt.Errorf("Inventory already contains product '%s'", product.Name)
    } else {
        s.getInventory()[product.ID] = product
        return nil
    }
}

func (s *Store) removeProduct(ID string) (e error) {
    if contained, e := s.containsProduct(ID); e!=nil {
        return e
    } else if !contained {
        return fmt.Errorf("Product with ID:'%s' not found", ID)
    } else {
        delete(s.getInventory(), ID)
        return nil
    }
}

func (s *Store) sellProduct(ID string, units int) (float32, error) {
    er := s.checkInventory()
    if er != nil {
        return 0, er
    }

    if p, e := s.getProduct(ID); e != nil {
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

func (s *Store) addProductUnits(ID string, units int) (e error) {
    er := s.checkInventory()
    if er != nil {
        return er
    }

    if p, e := s.getProduct(ID); e != nil {
        return e
    } else {
        p.Units += units
        return nil
    }
}

func (s *Store) checkInventory() (e error) {
    if s.getInventory() == nil {
        return fmt.Errorf("inventory is nil")
    } else {
        return nil
    }
}

func (s *Store) containsProduct(ID string) (isContained bool, e error) {
    if e := s.checkInventory(); e != nil {
        return false, e
    } else {
        _, found := s.getInventory()[ID]
        return found, nil
    }
}

func (s *Store) getProduct(ID string) (product *Product, e error) {
    if p, found := s.getInventory()[ID]; found {
        return p, nil
    }
    return nil, fmt.Errorf("Product not found")
}