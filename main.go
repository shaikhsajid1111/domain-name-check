package main

import (
	"net/http"
)

func main(){
	mux := register()
	http.ListenAndServe(":4000",mux)

}