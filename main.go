package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv" // necessary to convert integer to string
	"time"
)

//table1, 2 and 3 so assignment is done all at once
type Person struct {
	Firstname string
	Lastname  string
	Table1    string
	Table2    string
	Table3    string
}

//var table declared globally because it is used in below func findTable()
var table = 1

func findTable() {
	//random assignment of tables
	table = rand.Intn(33)
	//gets rid of table 0
	table++

}

func contains(s []int, e int) bool {
	//searches for integer in slice - checks if table is already filled
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	//assignment is string for either kitchen crew or waiter
	var assignment = ""
	// how much each person changes tables between dinners
	var tablevar = 1
	//appends to struct Person - able to manipulate variable
	var table1disp = ""
	var table2disp = ""
	var table3disp = ""

	//necessary for random assignment 
	rand.Seed(time.Now().UnixNano())

	//each table has 0 members to begin, after assignment they will be filled
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//usedTables is empty slice, moves tables with 8 people into slice
	usedTables := []int{}


	//parses csv data
	csvFile, _ := os.Open("list.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//assign random number to table
		findTable()

		for {
			//continues to find a new number until it finds a table or assignment that isn't filled up yet - 2nd and 3rd seating
			if contains(usedTables, table) {
				findTable()
			} else {
				break
			}
		}
		//sets tablevar at the specific person's order of seating in the table. For example, fifth person at a table gets a tableVar value of 5.
		tablevar = tableFill[table] + 1
		//loopvar is for 1...3 - run the loop 3 times
		var loopvar = 1

		for {
			if table < 32 {
				//only runs on the first loop
				if loopvar == 1 {
					//continues to fill table until there are 8 people
					if tableFill[table] < 8 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//once 8 people are at the table, adds one more and adds the table to usedTables to ignore
						usedTables = append(usedTables, table)
					}
				}
				//converts int to string 
				assignment = strconv.Itoa(table)
			} else if table == 32 {
				//again only runs on the first loop
				if loopvar == 1 {
					//table 32 is kitchen crew
					if tableFill[table] < 6 {
						tableFill[table]++
					} else {
						tableFill[table]++
						usedTables = append(usedTables, table)
					}
				}
				//prints assignment - if table assigned is 32, print Kitchen Crew
				assignment = "Kitchen Crew"
			} else if table == 33 {
				//only runs in first loop
				if loopvar == 1 {
					//table 33 will be assigned to waiters
					if tableFill[table] < 30 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//same as kitchen crew
						usedTables = append(usedTables, table)
					}
				}
				assignment = "Waiter"
			}
			//Each run-through of the loop sets the assignment for one seating
			if loopvar == 1 {
				table1disp = assignment
			} else if loopvar == 2 {
				table2disp = assignment
			} else if loopvar == 3 {
				table3disp = assignment
			}

			//move to the next table assignment and repeat
			table = table + tablevar
			//make sure to avoid index out of range error
			for {
				if table > 33 {
					table = table - 33
				} else {
					break
				}
			}

			//manually continue the loop
			if loopvar < 3 {
				loopvar++
			} else if loopvar == 3 {
				break
			}

		}
		//finished three seatings, append to people and print
		people := Person{
			Firstname: line[1],
			Lastname:  line[0],
			Table1:    table1disp,
			Table2:    table2disp,
			Table3:    table3disp,
		}
		fmt.Println(people)

	}
}
