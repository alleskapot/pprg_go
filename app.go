package main

import (
	"fmt"
	"strconv"
	"log"
	"os"
	"bufio"
)

type Item struct {
	name   string
	weight int
	value  int
}

type Solution struct {
	items       *[]Item // we just need the reference to save memory
	totalValue  int
	totalWeight int
}

func main() {

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("Wir arbeiten mit Google Go Go - https://www.youtube.com/watch?v=pIgZ7gMze7A \n")
	fmt.Println("---------------------------------------------------------------------------")


	// Read Knapsack Capacity from file
	var capacity int
	capacity = readCapacity("testdata/capacity")
	fmt.Printf("Capacity: %[1]d \n", capacity)

	// Read Weight of each item from File
	var weight []int
	weight = readWeightAndProfit("testdata/weight")

	// Read Profit of each item from File
	var profit []int
	profit = readWeightAndProfit("testdata/profit")

	// Create a map with key and values
	items := []Item {}
	for i := 0; i < len(weight); i++ {
		key := "item" + strconv.Itoa(i)
		items = Extend(items, Item {key, weight[i], profit[i]})
	}
	fmt.Println("Items: ", items)

	knapsack(items, capacity)

}


//TODO:  Müssten noch in ein anderes Files ausgelagert werden
func readCapacity(filename string) int {

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


func readWeightAndProfit(filename string) []int {

	var m []int

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

func Extend(slice []Item, element Item) []Item {
	n := len(slice)
	if n == cap(slice) {
		// +1 because there can be a slice with cap 0
		newSlice := make([]Item, len(slice), len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func knapsack(items []Item, knapsackSize int) Solution {
	return Solution {nil, 0, 0};
}
