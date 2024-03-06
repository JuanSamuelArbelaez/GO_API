package services

import (
    "module github.com/JuanSamuelArbelaez/GO_API/model"
)

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