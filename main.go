package main

import (
	"ecommerce-ums/cmd"
	"ecommerce-ums/helpers"
)

func main() {
	// Load config
	helpers.SetupConfig()

	// Load log 
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// load redis 
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	// run http
	cmd.ServeHTTP()
}