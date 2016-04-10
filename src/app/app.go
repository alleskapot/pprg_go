package main

import (
	"fmt"
	"strconv"
	"filereader"
	"knapsack"
	"time"
)



func main() {

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("Wir arbeiten mit Google Go Go - https://www.youtube.com/watch?v=pIgZ7gMze7A \n")
	fmt.Printf("Beispiel: Knapsack Problem (https://en.wikipedia.org/wiki/Knapsack_problem) \n")
	fmt.Println("---------------------------------------------------------------------------")


	// Read Knapsack Capacity from file
	var capacity int
	capacity = filereader.ReadCapacity("testdata/capacity")

	// Read Weight of each item from File
	var weight []int
	weight = filereader.ReadWeightAndProfit("testdata/weight")

	// Read Profit of each item from File
	var profit []int
	profit = filereader.ReadWeightAndProfit("testdata/profit")

	// Create a map with key and values
	items := []knapsack.Item {}
	for i := 0; i < len(weight); i++ {
		key := "item" + strconv.Itoa(i)
		items = append(items, knapsack.Item {key, weight[i], profit[i]})
	}

	fmt.Printf("Capacity: %[1]d \n", capacity)
	fmt.Println("Items: ", items)


	// RECURSIVE
	startRecursive := time.Now()
	s := knapsack.SolveRecursive(items, capacity)
	elapsedRecursive := time.Since(startRecursive)

	fmt.Printf("Using SolveRecursive\n")
	fmt.Printf("################################# RESULT ##################################\n")
	fmt.Println("Take the following items: ",s.Items)
	fmt.Println("weight:", s.TotalWeight)
	fmt.Println("value:", s.TotalValue)
	fmt.Printf("Time elapsed: %s\n", elapsedRecursive)
	fmt.Printf("###########################################################################\n")



	// PARALLEL
	startParallel := time.Now()
	resultDynamic := knapsack.KnapsackParallel(items, capacity)
	ra, rb := knapsack.ShowOptimalSolution(items, resultDynamic, capacity)
	elapsedParallel := time.Since(startParallel)

	fmt.Printf("\n")
	fmt.Printf("Using SolveParallel\n")
	fmt.Printf("################################# RESULT ##################################\n")
	fmt.Println("Take the following items: ", rb)
	fmt.Println("weight:", ra)
	fmt.Println("value:", resultDynamic[len(items)-1][capacity])
	fmt.Printf("Time elapsed: %s\n", elapsedParallel)
	fmt.Printf("###########################################################################\n")



}



