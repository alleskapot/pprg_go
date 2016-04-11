package filereader

import (
	"os"
	"log"
	"bufio"
	"strconv"
)


//TODO:  Müssten noch in ein anderes Files ausgelagert werden
func ReadCapacity(filename string) int {

	var capacity string
	file, err := os.Open(filename)

	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		capacity = scanner.Text()
	}

	if err := scanner.Err(); err != nil { log.Fatal(err) }

	// Convert back to int
	capacityInt, err := strconv.Atoi(capacity)

	return capacityInt
}


func ReadWeightAndProfit(filename string) []int {

	var m []int
	m = append(m, 0)

	fileWeight, err := os.Open(filename)

	if err != nil { log.Fatal(err) }
	defer fileWeight.Close()

	scanner := bufio.NewScanner(fileWeight)
	for scanner.Scan() {

		// Convert to int
		was, err := strconv.Atoi(scanner.Text())
		if err != nil { log.Fatal(err) }

		m = append(m, was)
	}

	if err := scanner.Err(); err != nil { log.Fatal(err) }

	return m
}