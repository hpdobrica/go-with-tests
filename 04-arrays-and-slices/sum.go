package main

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(slicesToSum ...[]int) []int {
	// numberOfSlices := len(slicesToSum)
	// sums := make([]int, numberOfSlices)

	// for i, slice := range slicesToSum {
	// 	sums[i] = Sum(slice)
	// }
	// return sums

	var sums []int

	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	// numberOfSlices := len(slicesToSum)
	// sums := make([]int, numberOfSlices)

	// for i, slice := range slicesToSum {
	// 	sums[i] = Sum(slice)
	// }
	// return sums

	var sums []int

	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			sums = append(sums, Sum(slice))
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
