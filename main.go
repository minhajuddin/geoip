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

func main() {
	loadDb()
	sortDb()
	for _, x := range db {
		log.Println(x)
	}

	log.Println(SearchLocation(16779265))
}

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
			log.Println("invalid number", tokens[0])
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

type Location struct {
	CountryCode string
	IpStart     int32
	IpEnd       int32
}

//sort interface
type Locations []Location

func (self Locations) Len() int           { return len(self) }
func (self Locations) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self Locations) Less(i, j int) bool { return self[i].IpStart < self[j].IpStart }

//search
//copied straight from sort/search.go
func SearchLocation(ip int32) string {
	i, j := 0, len(db)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if db[h].IpStart < ip {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return db[i].CountryCode
}
