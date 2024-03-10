package main

import (
	"vistaverse/src"
	"vistaverse/src/db"
)

func main() {
	// to setup db connection
	db.GetStorage()
	// start listening
	src.RunAPI(":3030")
}
