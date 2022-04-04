package main

import (
	"net/http"
	"os"
)

func GetRequest(suffix string) {
	http.Get(os.Getenv("host") + "/" + suffix)
}

func PostRequest() {

}

func PostFormRequest() {

}
