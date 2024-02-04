package main

import (
	"cleanarchitecturego/cockroach/entities"
	"cleanarchitecturego/config"
	"cleanarchitecturego/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	//server.NewEchoServer(&cfg, db.GetDb()).Start()
	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
	db.GetDb().CreateInBatches([]entities.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)
}
