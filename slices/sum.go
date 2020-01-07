package slices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumber := len(numbersToSum)
	sums = make([]int, lengthOfNumber)
	for idx, numbers := range numbersToSum {
		sums[idx] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := Sum(numbers[1:])
			sums = append(sums, tail)
		}
	}
	return sums
}
