package main

import (
	"github.com/arempe93/experiment"
	"github.com/arempe93/experiment/database"
)

func main() {
	experiment.ReadConfig()
	defer database.Close()

	database.Migrate()
}
