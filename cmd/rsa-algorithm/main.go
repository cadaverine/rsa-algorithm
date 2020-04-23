package main

import (
	"github.com/cadaverine/rsa-algorithm/utils"
)

// Алгоритм RSA
// 1) выбрать 2 произвольных простых числа

const (
	// диапазон номеров простых чисел
	min = 1
	max = 100
)

// Key - структура для хранения ключа
type Key struct {
	E int64
	N int64
}

func GetRandomKeys() (Key, Key, error) {
	return Key{}, Key{}, nil
}

func EncryptMessage(message string, publicKey Key) string {
	return ""
}

func DecryptMessage(message string, privateKey Key) string {
	return ""
}

func main() {
	// Алгоритм RSA
	// 1) выбрать 2 произвольных простых числа
	p := utils.GetPrimeNumber(utils.GetRandInt(min, max))
	q := utils.GetPrimeNumber(utils.GetRandInt(min, max))

	// 2) находим их произведение
	n := p * q

	println(p)
	println(q)
	println(n)

	// 3) функция Эйлера F(n)=(p-1)*(q-1)"

}
