package main

import "log"

func main() {
	log.Println("loading db")
	loadDb()
	log.Println("sorting db")
	sortDb()
	log.Println("starting server at http://localhost:3000/")
	startWebServer(":3000")
}
