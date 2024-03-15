package main

import (
	"eventopia/src"
	"eventopia/src/db"
)

func main() {
	// to setup db connection
	db.GetStorage()
	// start listening
	src.RunAPI(":3030")
}
