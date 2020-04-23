package utils

// GetPrimeNumbers - получаем N первых простых чисел
func GetPrimeNumbers(num int) []int {
	primes := []int{}

Loop:
	for n := 2; len(primes)+1 <= num; n++ {
		for _, prime := range primes {
			if n%prime == 0 {
				continue Loop
			}
		}

		primes = append(primes, n)
	}

	return primes
}

// GetPrimeNumber - получаем i-ое простое число
func GetPrimeNumber(index int) int {
	primes := GetPrimeNumbers(index + 1)

	return primes[index]
}
