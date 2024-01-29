package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i := range numbersToSum {
		sums[i] = Sum(numbersToSum[i])
	}

	return sums
}

func SumTail(numbers ...[]int) []int {
	var tailSums []int

	for _, num := range numbers {
		sum := 0

		if len(num) > 0 {
			tail := num[1:]
			sum = Sum(tail)
		}

		tailSums = append(tailSums, sum)
	}

	return tailSums
}
