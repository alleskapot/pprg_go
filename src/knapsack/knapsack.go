package knapsack

import (
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

func SolveRecursive(items []Item, knapsackSize int) Solution {
	result, w, v := m(items, len(items)-1, knapsackSize)

	return Solution {result, v, w}
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



func KnapsackParallel(items []Item, knapsackSize int) [][]int {

	n := len(items)-1
	W := knapsackSize

	m := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		m[i] = make([]int, W+1)
	}
	for i := 0; i <= W; i++ {
		m[0][i] = 0
	}

	done := make(chan bool)

	go func() {

		for i := 1; i <= n; i++ {

			for j := 0; j <= W; j++ {

				if items[i].Weight > j {
					m[i][j] = m[i - 1][j]
				} else if m[i - 1][j] > m[i - 1][j - items[i].Weight] + items[i].Value {
					m[i][j] = m[i - 1][j]
				} else {
					m[i][j] = m[i - 1][j - items[i].Weight] + items[i].Value
				}

			}
			done <- true
		}
	} ();

	for x := 1; x <= n; x++ {
		<- done
	}

	return m
}


func ShowOptimalSolution(items []Item, m [][]int, knapsackSize int) (int, []int) {

	finalValue := 0
	result := []int{}
	W := knapsackSize
	n := len(items)-1

	for W > 0 && n > 0 {
		if m[n][W] != m[n-1][W] {
			result = append(result, n)
			W = W - items[n].Weight
			finalValue += items[n].Value
			n--
		} else {
			n--
		}
	}
	return finalValue, result
}
