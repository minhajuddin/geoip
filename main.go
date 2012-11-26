package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//TODO: make this a runtime value?
const DB_SIZE = 110000

//NOTE:Can this be stored in a map somehow?
var db []Location

func main() {
	file, err := os.Open("countries.txt")
	defer file.Close()
	r := bufio.NewReader(file)
	if err != nil {
		log.Fatalln("unable to open countries.txt", err)
	}
	db = make([]Location, DB_SIZE)
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

		l := &Location{
			Start: int32(start),
			End:   int32(end),
			Code:  tokens[2],
		}

		log.Println(l)

	}
}

type Location struct {
	Code  string
	Start int32
	End   int32
}
