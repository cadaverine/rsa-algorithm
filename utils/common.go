package utils

import "math"

const (
	// диапазон номеров простых чисел
	min = 1
	max = 100
)

// GetGCF - поиск наибольшего общего делителя (greatest common factor)
func GetGCF(first, second int) int {
	a := first
	b := second

	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	return a + b
}

// GetExtendedGCF - расширенный алгоритм Евклида
// J×A + K×B = R
// НОД(A, B) = R
func GetExtendedGCF(first, second int) (int, int, int) {
	a, b := first, second
	j, xx, k, yy := 1, 0, 0, 1

	for b > 0 {
		q := a / b
		a, b = b, a%b

		j, xx = xx, j-xx*q
		k, yy = yy, k-yy*q
	}

	return j, k, a
}

// EulerFunction - расчет функции Эйлера
func EulerFunction(value int) int {
	counter := 0

	for i := 1; i < value; i++ {
		if GetGCF(value, i) == 1 {
			counter++
		}
	}

	return counter
}

// GetRandomCoprimeNumber - получить случайное взаимно простое число с заданным
func GetRandomCoprimeNumber(value int) int {
	for i := GetRandInt(min, max); ; i++ {
		if GetGCF(value, i) == 1 {
			return i
		}
	}
}

// Pow - возведение в степень целочисленных переменных
func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
