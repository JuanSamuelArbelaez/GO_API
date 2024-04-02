package services

/*
import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
)

func SellProduct(ID string, units int) (total float32, e error) {
	if p, e := GetPerson(ID); e != nil {
		return 0, e
	} else {
		if units > p.Units {
			return 0, fmt.Errorf("not enough units to sell (%d). currently available: %d", units, p.Units)
		}
			if e := SetProductUnits(p, GetProductUnits(p)-units); e != nil {
				return 0, e
			}
		return float32(units) * p.UnitValue, nil
	}
}

func AddProductUnits(ID string, units int) (e error) {
	if p, e := GetPerson(ID); e != nil {
		return e
	} else {
		if units <= 0 {
			return fmt.Errorf("cannot add 0 or negative units")
		}

			if e := SetProductUnits(p, GetProductUnits(p)+units); e != nil {
				return e
			}

		err := CheckPerson(model.PersonRequest{Name: p.Name, UnitValue: p.UnitValue, Units: p.Units})
		if err != nil {
			return err
		}
		return nil
	}
}
*/
