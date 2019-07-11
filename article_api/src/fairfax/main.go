package main

import (
	"fairfax/http_server"
	"fairfax/search_helper"
	log "fairfax/logging"
)


func main() {
	search_helper.DbInstance()
	log.Info("Greetings to Fairfax, es conn established, http will listen on 8080")
	http_server.ListenAndServe()
}
