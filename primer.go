package primer

func PrimeFactor(n int) []int {
	factors, primeNumbers := []int{}, []int{2, 3, 5, 7, 11, 13}
	for index := 0; n != 1; index++ {
		primeNumber := primeNumbers[index]
		for ; n%primeNumber == 0; n /= primeNumber {
			factors = append(factors, primeNumber)
		}
	}
	return factors
}
