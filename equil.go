package main

import "fmt"

func main() {
	data := []int64{-1, 3, -4, 5, 1, -6, 2, 1}
	equill := solution(data)
	fmt.Println(equill)
}

func solution(data []int64) []int64 {
	var lSum, rSum int64
	var p int64
	var ret []int64

	for _, value := range data {
		rSum += value
	}

	if rSum-data[p] == 0 {
		return append(ret, p)
	}

	lSum = data[0]

	for i := 0; i < len(data) - 1; i++ {
		p++
		rSum -= (data[p] + lSum)
		if lSum == rSum {
			ret = append(ret, p)
		}

		lSum += data[p]
	}
	return ret
}
