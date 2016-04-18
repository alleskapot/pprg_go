package knapsack

type Item struct {
	Name   string
	Weight int
	Value  int
}

func SolveRecursive(items []Item, itemSize int, knapsackSize int) ([] string, int, int) {

	k := knapsackSize

	if itemSize < 0 || k == 0 {
		return nil, 0, 0
	} else if items[itemSize].Weight > k {
		return SolveRecursive(items, itemSize - 1, k)
	}

	item0, weight0, value0 := SolveRecursive(items, itemSize - 1, k)
	item1, weight1, value1 := SolveRecursive(items, itemSize - 1, k - items[itemSize].Weight)

	value1 += items[itemSize].Value

	if value1 > value0 {
		return append(item1, items[itemSize].Name), weight1 + items[itemSize].Weight, value1
	}

	return item0, weight0, value0
}


func KnapsackParallel(items []Item, knapsackSize int) [][]int {

	n := len(items)-1
	k := knapsackSize

	m := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		m[i] = make([]int, k +1)
	}

	done := make(chan bool)

	go func() {
		for i := 1; i <= n; i++ {
			for j := 0; j <= k; j++ {
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



	// Das funktioniert - aber ziemlich langsam -.-
	// nur bei einer hohen Anzahl an Werten > 100 ist Parallel ziemlich ähnlich wie Dynamic
	//for i := 1; i <= n; i++ {
	//	go func() {
	//		for j := 0; j <= k; j++ {
	//			if items[i].Weight > j {
	//				m[i][j] = m[i - 1][j]
	//			} else if m[i - 1][j] > m[i - 1][j - items[i].Weight] + items[i].Value {
	//				m[i][j] = m[i - 1][j]
	//			} else {
	//				m[i][j] = m[i - 1][j - items[i].Weight] + items[i].Value
	//			}
	//			done <- true
	//		}
	//	} ();
	//
	//	for x := 0; x <= k; x++ {
	//		<- done
	//	}
	//}

	//for i := 1; i <= n; i++ {
	//	go func() {
	//		for j := 0; j <= k; j++ {
	//			if items[i].Weight > j {
	//				m[i][j] = m[i - 1][j]
	//			} else if m[i - 1][j] > m[i - 1][j - items[i].Weight] + items[i].Value {
	//				m[i][j] = m[i - 1][j]
	//			} else {
	//				m[i][j] = m[i - 1][j - items[i].Weight] + items[i].Value
	//			}
	//			done <- true
	//		}
	//	} ();
	//
	//	for x := 0; x <= k; x++ {
	//		<- done
	//	}
	//}



	return m
}



func KnapsackDynamic(items []Item, knapsackSize int) [][]int {

	n := len(items)-1
	k := knapsackSize

	m := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		m[i] = make([]int, k +1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if items[i].Weight > j {
				m[i][j] = m[i - 1][j]
			} else if m[i - 1][j] > m[i - 1][j - items[i].Weight] + items[i].Value {
				m[i][j] = m[i - 1][j]
			} else {
				m[i][j] = m[i - 1][j - items[i].Weight] + items[i].Value
			}
		}
	}

	return m
}


func ShowOptimalComposition(items []Item, m [][]int, knapsackSize int) ([]int, int, int) {

	itemIndex := []int {}
	value := 0
	weight := 0

	k := knapsackSize
	n := len(items) - 1

	for k > 0 && n > 0 {
		if m[n][k] != m[n-1][k] {
			// Add index to array
			itemIndex = append(itemIndex, n)

			// Reduce the size of knapsack
			k = k - items[n].Weight

			// Increase value and weight
			value += items[n].Value
			weight += items[n].Weight

			n--
		} else {
			n--
		}
	}
	return itemIndex, value, weight
}




