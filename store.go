package main

import "fmt"

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
