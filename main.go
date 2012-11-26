package main

func main() {
	loadDb()
	sortDb()
	startWebServer(":3000")
}
