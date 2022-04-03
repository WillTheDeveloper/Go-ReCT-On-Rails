package main

import (
	"os"
)

// Set the environment variables
func SetEnvironment() {
	os.Setenv("port", "8080")      //PORT OF THE SERVER
	os.Setenv("host", "localhost") //HOST ADDRESS OF THE SERVER
}

// Checks that the environment has been setup correctly
func CheckEnvironment() {
	if os.Getenv("port") == "" {
		println("Server port has not been set.")
		os.Exit(69)
	}
	if os.Getenv("host") == "" {
		println("Server host has not been set.")
		os.Exit(70)
	}
}
