package main

import (
	"collect/database"
	"collect/router"
)

func main() {
	database.InitDB()
	router.Init()
}
