package main

import (
	"assignment-2/databases"
	"fmt"
)

func main() {
	db, err := databases.Start()

	if err != nil {
		fmt.Println("Start Databases failed", err)
	}

	_ = db

}
