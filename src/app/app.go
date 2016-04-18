package main

import (
	"fmt"
	"strconv"
	"filereader"
	"knapsack"
	"time"
	"io/ioutil"
)

func main() {

	var dataSet = "p01";

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("Wir arbeiten mit Google Go Go - https://www.youtube.com/watch?v=pIgZ7gMze7A \n")
	fmt.Printf("Beispiel: Knapsack Problem (https://en.wikipedia.org/wiki/Knapsack_problem) \n")
	fmt.Println("---------------------------------------------------------------------------")

	// Read Knapsack Capacity from file
	var capacity int
	capacity = filereader.ReadCapacity(fmt.Sprint("testdata/", dataSet,"_c.txt"));

	// Read Weight of each item from File
	var weight []int
	weight = filereader.ReadWeightAndProfit(fmt.Sprint("testdata/", dataSet,"_w.txt"))

	// Read Profit of each item from File
	var profit []int
	profit = filereader.ReadWeightAndProfit(fmt.Sprint("testdata/", dataSet,"_p.txt"))

	// Create a map with key and values
	items := []knapsack.Item {}
	for i := 0; i < len(weight); i++ {
		key := "item" + strconv.Itoa(i)
		items = append(items, knapsack.Item {key, weight[i], profit[i]})
	}

	fmt.Printf("Capacity: %[1]d \n", capacity)
	fmt.Println("Items: ", items)


	SolveItRecursive(items, capacity)
	SolveItDynamicSequential(items, capacity)
	SolveItDynamicParallel(items, capacity)

	dat, _ := ioutil.ReadFile(fmt.Sprint("testdata/", dataSet, "_e.txt"));
	fmt.Printf("Expectation: %s", (string(dat)))
}

func SolveItRecursive(items []knapsack.Item, capacity int) {

	startTime := time.Now()
	item, weight, value := knapsack.SolveRecursive(items, len(items)-1, capacity)
	elapsedTime := time.Since(startTime)

	fmt.Println()
	fmt.Printf("Using SolveRecursive\n")
	fmt.Printf("################################# RESULT ##################################\n")
	fmt.Println("Take the following items: ", item)
	fmt.Println("Weight:", weight)
	fmt.Println("Value:", value)
	fmt.Printf("Time elapsed: %s\n", elapsedTime)
	fmt.Printf("###########################################################################\n")
}

func SolveItDynamicParallel(items []knapsack.Item, capacity int) {

	startTime := time.Now()
	result := knapsack.KnapsackParallel(items, capacity)
	itemNumber, value, weight := knapsack.ShowOptimalComposition(items, result, capacity)
	elapsedTime := time.Since(startTime)

	fmt.Printf("\n")
	fmt.Printf("Using SolveParallel\n")
	fmt.Printf("################################# RESULT ##################################\n")
	fmt.Print("Take the following items: [")
	for i := (len(itemNumber) - 1); i >= 0 ; i-- {
		fmt.Printf("item%d ", itemNumber[i])
	}
	fmt.Println("]");
	fmt.Println("Weight:", weight)
	fmt.Println("Value:", value)
	fmt.Printf("Time elapsed: %s\n", elapsedTime)
	fmt.Printf("###########################################################################\n")
}

func SolveItDynamicSequential(items []knapsack.Item, capacity int) {

	startTime := time.Now()
	result := knapsack.KnapsackDynamic(items, capacity)
	itemNumber, value, weight := knapsack.ShowOptimalComposition(items, result, capacity)
	elapsedTime := time.Since(startTime)

	fmt.Printf("\n")
	fmt.Printf("Using SolveDynamic\n")
	fmt.Printf("################################# RESULT ##################################\n")
	fmt.Print("Take the following items: [")
	for i := (len(itemNumber) - 1); i >= 0 ; i-- {
		fmt.Printf("item%d ", itemNumber[i])
	}
	fmt.Println("]");
	fmt.Println("Weight:", weight)
	fmt.Println("Value:", value)
	fmt.Printf("Time elapsed: %s\n", elapsedTime)
	fmt.Printf("###########################################################################\n")
}