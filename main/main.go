package main

import (
	"github.com/arempe93/experiment"
	"github.com/arempe93/experiment/database"
	"github.com/arempe93/experiment/server"
)

func main() {
	experiment.ReadConfig()
	defer database.Close()

	database.Migrate()

	server.Run()
}
