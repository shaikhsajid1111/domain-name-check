package main

import(
	"github.com/gorilla/mux"
)


func register() *mux.Router{

	mux := mux.NewRouter()
	mux.SkipClean(true)

	mux.HandleFunc("/json/{url}",jsonHandler)

	return mux

}