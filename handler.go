package main


import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func jsonHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet{
		/*fetch value from the request object*/
		vars := mux.Vars(request)	/*Vars object for parsing request*/
		domainName := vars["url"]
		/*results map data structure for storing kry,value pairs value of all urls
		structure works like,
		results[key:string] = value:bool
		i.e, results[url] = if_it_exists
		*/
		var results map[string]bool

		results = checkDomains(domainName)
		/*
		response object to store current url's results, that needs to be converted to 
		JSON format
		*/	
		var response = make(map[string]string,23) 

		/*iterating over results and changing the value for key to, more 
		readable results,like "Available" for urls that have got negative response
		 and "Not Available" for urls that have got positive response*/
		for key := range results{
			if results[key] == true{
				response[key] = "Not Available"
			} else{
				response[key] = "Available"
			}
		}
		/*converting the entire data to JSON format*/
		jsonResponse , error := json.Marshal(response)
		
		if error != nil{
			log.Fatal(error)
		}

		writer.Header().Set("Content-Type", "application/json")	/*setting header to JSON*/
		json.NewEncoder(writer).Encode(string(jsonResponse))	/*sending response for the request*/

	}	


}
