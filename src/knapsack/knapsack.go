package knapsack

import "time"

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

func SolveRecursive(items []Item, knapsackSize int) (Solution, time.Duration) {
	start := time.Now()
	result, w, v := m(items, len(items)-1, knapsackSize)
	elapsed := time.Since(start)

	return Solution {result, v, w}, elapsed;
}

func generate(items []Item) {
	//var partOne = items[:len(items)/2]
	//var partTwo = items[len(items)/2:]

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
