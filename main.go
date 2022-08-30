package main

import (
	"github.com/VJ-Vijay77/Rest-with-Gin/api"
	"github.com/VJ-Vijay77/Rest-with-Gin/db"
	_ "github.com/lib/pq"
)

func main() {

	db := db.InitDb()

	engine := api.NewServer(db)

	engine.Run(":8080")

}
