package utils

import "math/rand"

// GetRandInt - получить случайное число в диапазоне
func GetRandInt(min, max int) int {
	return rand.Intn(max-min) + min
}
