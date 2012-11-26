package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	ipint := ipToInt(ip)
	log.Printf("handling request for %#v %#v\n", ipint, ip)
	location := SearchLocation(ipint)
	fmt.Fprintln(w, location)
}

func startWebServer(port string) {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(port, nil)
}

func ipToInt(ip string) uint32 {
	log.Println("parsing", ip)
	ipa := net.ParseIP(ip)
	if ipa == nil {
		log.Println("unable to parse ip")
		return 0
	}
	ip4 := ipa[len(ipa)-4:]
	ipint := int64(ip4[0])*256*256*256 + int64(ip4[1])*256*256 + int64(ip4[2])*256 + int64(ip4[3])
	log.Printf("ipint: %#v %#v\n", ipint, ipa[len(ipa)-4:])
	return uint32(ipint)
}
