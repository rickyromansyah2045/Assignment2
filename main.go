package main

import (
	"assignment-2/router"
)

func main() {
	r := router.StartApp()

	r.Run(":8080")
}
