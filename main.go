package main

import (
	"ParkingLot/driver"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		query := scanner.Text()
		driver.RunQuery(query)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
