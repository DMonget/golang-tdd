package arrays

// Sum adds the value of a slice of integers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll adds the values of the slices of integers passed to it
func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAll adds the values of the slices of integers passed to it
func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
