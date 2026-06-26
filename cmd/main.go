package main

import (
	"spotsync/internal/config"
	"spotsync/internal/server"
)

func main() {
	env := config.LoadEnv()

	// Start the server
	server.Start(env)

}
