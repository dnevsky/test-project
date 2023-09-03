package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// test
	reader := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var wanted, numBanks int
	fmt.Fscan(reader, &wanted, &numBanks)

	denominations := make([]int, numBanks)
	for i := 0; i < numBanks; i++ {
		var bank int

		fmt.Fscan(reader, &bank)
		denominations[i] = bank
	}

	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))

	remaining := wanted
	stolen := make([]int, numBanks)

	for i := 0; i < numBanks; i++ {
		for remaining >= denominations[i] {
			if st := stolen[i]; st < 2 {
				remaining -= denominations[i]
				stolen[i]++
			} else {
				break
			}
		}
	}

	res := make([]int, 0)

	if remaining == 0 {
		for i := 0; i < numBanks; i++ {
			for j := 0; j < stolen[i]; j++ {
				res = append(res, denominations[i])
			}
		}
	} else {
		writer.WriteString("-1")
	}

	sort.Ints(res)

	writer.WriteString(fmt.Sprintf("%d\n", sum(stolen)))
	for i := 0; i < len(res); i++ {
		writer.WriteString(fmt.Sprintf("%d ", res[i]))
	}

}

func sum(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}
