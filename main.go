package main

import (
	"encoding/binary"
	"net"
)

func main() {
	loadDb()
	sortDb()
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
