package main

import (
    "fmt"
    "pantry/pantryObjects"
    "database/sql"
    _ "modernc.org/sqlite"
)

func main () {
    i := pantryObjects.Ingredient{Name: "hallo", Perishable: true, ExpDate:""}
    db, _ := sql.Open("sqlite", "./test.sql")
    fmt.Println(i.GetName())
    fmt.Println(db)
}
