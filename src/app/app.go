package main

import (
	"fmt"
	"strconv"
	"filereader"
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
	capacity = filereader.ReadCapacity("testdata/capacity")
	fmt.Printf("Capacity: %[1]d \n", capacity)

	// Read Weight of each item from File
	var weight []int
	weight = filereader.ReadWeightAndProfit("testdata/weight")

	// Read Profit of each item from File
	var profit []int
	profit = filereader.ReadWeightAndProfit("testdata/profit")

	// Create a map with key and values
	items := []Item {}
	for i := 0; i < len(weight); i++ {
		key := "item" + strconv.Itoa(i)
		items = append(items, Item {key, weight[i], profit[i]})
	}
	fmt.Println("Items: ", items)

	knapsack(items, capacity)

}


func knapsack(items []Item, knapsackSize int) Solution {
	return Solution {nil, 0, 0};
}
