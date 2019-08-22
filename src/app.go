package main

import (
	"GoFirst/src/core/Injector"
	"log"
)

func main() {
	controller, err := Injector.NewCommonController()
	if err != nil {
		log.Fatal(err)
	}
	err = controller.RunAPI("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
