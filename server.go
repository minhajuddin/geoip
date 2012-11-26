package main

import (
	"net/http"
	"encoding/binary"
	"net"
	"fmt"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	location := SearchLocation(ipToInt(ip))
	fmt.Fprintln(w, location)
}

func startWebServer(port string) {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(port, nil)
}


func ipToInt(ip string) int32 {
	ipa := net.ParseIP(ip)
	if ipa == nil {
		return 0
	}
	ipint, n := binary.Varint(ipa)
	if n > 0 {
		return int32(ipint)
	}
	return 0
}
