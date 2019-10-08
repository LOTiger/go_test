package integers

// Sum is a func
func Sum(numbers []int)  int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll will return sum of slice
func SumAll(numbersToSum ...[] int) (sums []int)  {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}