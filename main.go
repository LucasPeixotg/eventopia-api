package main

import (
	"github.com/LucasPeixotg/eventopia-api/src"
	"github.com/LucasPeixotg/eventopia-api/src/db"
)

func main() {
	// to setup db connection
	db.GetStorage()
	// start listening
	src.RunAPI(":3030")
}
