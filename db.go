package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//TODO: make this a runtime value?
const DB_SIZE = 110000

//NOTE:Can this be stored in a map somehow?
var db Locations

func loadDb() {
	file, err := os.Open("countries.txt")
	defer file.Close()
	r := bufio.NewReader(file)
	if err != nil {
		log.Fatalln("unable to open countries.txt", err)
	}
	db = make(Locations, 0, DB_SIZE)
	for {
		line, err := r.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			//EOF probably reached
			log.Println("EOF probably reached", err)
			break
		}

		tokens := strings.Split(line, " ")

		if len(tokens) != 3 {
			log.Println("invalid number of tokens", line)
			continue
		}

		start, err := strconv.ParseInt(tokens[0], 10, 32)
		if err != nil {
			log.Println("invalid number", tokens[0], err)
			continue
		}

		end, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			log.Println("invalid number", tokens[1])
			continue
		}

		l := Location{
			IpStart:     int32(start),
			IpEnd:       int32(end),
			CountryCode: tokens[2],
		}

		db = append(db, l)
	}
}

func sortDb() {
	sort.Sort(db)
}
