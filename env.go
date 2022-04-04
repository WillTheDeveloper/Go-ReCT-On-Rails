package main

import (
	"os"
)

// Set the environment variables
func SetEnvironment() {
	os.Setenv("port", "8080")        //PORT OF THE SERVER
	os.Setenv("host", "localhost")   //HOST ADDRESS OF THE SERVER
	os.Setenv("db_user", "")         //USERNAME FOR THE DATABASE
	os.Setenv("db_pass", "")         //PASSWORD FOR THE DATABASE
	os.Setenv("db_url", "")          //DATABASE CONNECTION URL
	os.Setenv("db_present", "false") //DATABASE PRESENCE CHECK
	os.Setenv("db_name", "")         //DATABASE TABLE NAME
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
	if os.Getenv("db_url") != "" {
		println("Database URL has been set.")
		if os.Getenv("db_user") == "" {
			println("Database username has not been set")
		}
		if os.Getenv("db_pass") == "" {
			println("Database password has not been set")
		}
	}
	if os.Getenv("db_url") == "" {
		println("No database url present, assuming its not being used.")
	}
}
