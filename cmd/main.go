package main

import (
	"spotsync/internal/config"
	"spotsync/internal/server"
)

func main() {
	// Load environment variables
	env := config.LoadEnv()
	// Configure postgres database connection
	db := config.ConnectDatabase(env)
	// Start the server
	server.Start(db, env)

}
