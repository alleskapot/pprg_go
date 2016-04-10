package knapsack

import (
	"time"
	"fmt"
)

type Item struct {
	Name   string
	Weight int
	Value  int
}

type Solution struct {
	Items       []string
	TotalValue  int
	TotalWeight int
}

type SubsetSum struct {
	subset []Item
	sum    int
}

// needed to enable parallel execution
type empty struct {}
type semaphore chan empty

func SolveRecursive(items []Item, knapsackSize int) (Solution, time.Duration) {
	start := time.Now()
	result, w, v := m(items, len(items)-1, knapsackSize)
	elapsed := time.Since(start)

	return Solution {result, v, w}, elapsed
}

func SolveParallel(items []Item, knapsackSize int) (Solution, time.Duration) {
	// http://stackoverflow.com/questions/22373663/1-0-knapsack-how-to-make-it-parallel-with-priority-queue
	// http://www.bogotobogo.com/Algorithms/knapsack.php
	// https://github.com/mauromoura/0-1-knapsack-problem/blob/master/knapsack.go

	start := time.Now()

	//done := make(chan bool)

	// we only benefit from a parallel solution if we have a bigger list
	ps := []SubsetSum{{nil, 0}}
	for _, i := range items {
		itemSubsetSum := SubsetSum{[]Item{i}, i.Weight}
		itemSubset := []SubsetSum{itemSubsetSum}
		for _, otherItem := range items {
			if(contains(itemSubsetSum.subset, otherItem) == false) {
				subset := append([]Item {i}, otherItem)
				sum := i.Weight + otherItem.Weight
				if(sum <= knapsackSize) {
					itemSubset = append(itemSubset, SubsetSum{subset, sum})
				}
			}
		}
		fmt.Println(itemSubset)
	}

	/*for _ = range items {
		<-done
	}*/

	elapsed := time.Since(start)

	fmt.Println("\nSubset-Sums: ",ps)

	// now we need to find the "perfect" subset sum for our needs
	// if we put the value in the SubSet struct we can just find the maximum Value and the highest possible Weight from ps

	return Solution {nil, 0, 0}, elapsed
}

func m(items []Item, itemSize int, maxWeight int) ([]string, int, int) {
	if itemSize < 0 || maxWeight == 0 {
		return nil, 0, 0
	} else if items[itemSize].Weight > maxWeight {
		return m(items,itemSize-1, maxWeight)
	}
	i0, w0, v0 := m(items,itemSize-1, maxWeight)
	i1, w1, v1 := m(items, itemSize-1, maxWeight-items[itemSize].Weight)
	v1 += items[itemSize].Value
	if v1 > v0 {
		return append(i1, items[itemSize].Name), w1 + items[itemSize].Weight, v1
	}
	return i0, w0, v0
}

func contains(s []Item, e Item) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
