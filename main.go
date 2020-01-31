package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var table = 1
var assignment = ""
var waiterTable = 1

func findTable() {
	//find a new table assignment number
	table = rand.Intn(33)
	table++

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	//number of elements = number of tables, integer represents number of people at table
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//when table has seven people, add table to this slice
	usedTables := []int{}

	//necessary for random assignment
	rand.Seed(time.Now().UnixNano())

	//to parse csv data
	csvFile, _ := os.Open("list.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//read csv data
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		findTable()

		for {
			if contains(usedTables, table) {
				findTable()
			} else {
				break
			}
		}

		if table < 32 {
			//everything that satisfies this if statement will be sitting at a table

			if tableFill[table] < 8 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//remove table from tableNumber, and remove table from tableFill
				usedTables = append(usedTables, table)
			}
			fmt.Println(line[0], line[1], table)

		} else if table == 32 {
			//table 32 is kitchen crew

			if tableFill[table] < 6 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//remove table from tableNumber, and remove table from tableFill
				usedTables = append(usedTables, table)
			}
			assignment = "Kitchen Crew"
			fmt.Println(line[0], line[1], assignment)

		} else if table == 33 {
			//table 33 are waiters
			if tableFill[table] < 30 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//remove table from tableNumber, and remove table from tableFill
				usedTables = append(usedTables, table)
			}
			assignment = "Waiter"
			waiterTable++
			//prints list of people with assigned number
			fmt.Println(line[0], line[1], assignment, waiterTable)
		}

	}

}
