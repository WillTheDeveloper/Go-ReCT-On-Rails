package main

import (
	"os"
)

// Set the environment variables
func SetEnvironment() {
	os.Setenv("port", "8080")      //PORT OF THE SERVER
	os.Setenv("host", "localhost") //HOST ADDRESS OF THE SERVER
}
