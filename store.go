package main

import "fmt"

//PRODUCT
type Product struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    UnitValue float32 `json:"value"`
    Units     int     `json:"units"`
}

func (p *Product) checkProduct() (e error) {
    if p.getProductID() == "" || p.getProductName() == "" || p.getProductUnitValue() <= 0 || p.getProductUnits() <= 0 {
        return fmt.Errorf("Product is invalid")
    } else {
        return nil
    }
}

func (p *Product) setProductID(ID string) (e error){
    if ID == "" {
        return fmt.Errorf("must provide a ID")
    }
    p.ID = ID
    return nil
}

func (p *Product) getProductID() (ID string) {
    return p.ID
}

func (p *Product) setProductName(name string) (e error){
    if name == "" {
        return fmt.Errorf("must provide a name")
    }
    p.Name = name
    return nil
}

func (p *Product) getProductName() (name string) {
    return p.Name
}

func (p *Product) setProductUnitValue(value float32) (e error) {
    if value <= 0{
        return fmt.Errorf("value must be positive")
    }
    p.UnitValue = value
    return nil
}

func (p *Product) getProductUnitValue() (unitValue float32) {
    return p.UnitValue
}

func (p *Product) setProductUnits(units int) (e error){
    if units <= 0 {
        return fmt.Errorf("units must be positive")
    }
    p.Units = units
    return nil
}

func (p *Product) getProductUnits() (units int) {
    return p.Units
}

//STORE
type Store struct {
    Inventory map[string]*Product
}

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


func main() {
    var store = Store{make(map[string]*Product)}

    mlk := &Product{"100", "Milk Carton", 2500, 98}
    eg := &Product{"101", "Egg Carton", 7800, 75}
    flr := &Product{"102", "Flour Bag", 3200, 63}

    store.addProduct(mlk)
    store.addProduct(eg)
    store.addProduct(flr)

    store.removeProduct(eg.ID)

    fmt.Println(store.getInventorySize())

    fmt.Println(store.containsProduct("100"))

    fmt.Println(store.getProduct("100"))
}
