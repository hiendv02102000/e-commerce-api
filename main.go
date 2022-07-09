package main

import (
	"api/pkg/infrastucture/db"
	"fmt"
)

func main() {
	d, err := db.NewDB()
	d.MigrateDBWithGorm()
	fmt.Println(err)
}
