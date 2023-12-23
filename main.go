package main

import (
	"github.com/krisnafirdaus/go-project/database"
	"github.com/krisnaifrdaus/go-project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}