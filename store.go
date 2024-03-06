package main

import (
    "fmt"
    "github.com/JuanSamuelArbelaez/GO_API/model"
    "github.com/JuanSamuelArbelaez/GO_API/services"
)

func main() {
    var store = model.Store{make(map[string]model.Product)}

    mlk := &model.Product{"100", "Milk Carton", 2500, 98}
    eg := &model.Product{"101", "Egg Carton", 7800, 75}
    flr := &model.Product{"102", "Flour Bag", 3200, 63}

    services.addProduct(store, mlk)
    services.addProduct(store, eg)
    services.addProduct(store, flr)

    services.removeProduct(eg.ID)

    fmt.Println(services.getInventorySize(store))

    fmt.Println(services.containsProduct(store, "100"))

    fmt.Println(services.getProduct("100"))
}
