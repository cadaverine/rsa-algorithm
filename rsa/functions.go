package rsa

import (
	"github.com/cadaverine/rsa-algorithm/utils"
)

const (
	// диапазон номеров простых чисел
	min = 1
	max = 100
)

// GetRandomRSAKeys - получить ключи шифрования
func GetRandomRSAKeys() (Key, Key, error) {
	// Алгоритм RSA
	// 1) выбрать 2 произвольных простых числа
	p := utils.GetPrimeNumber(utils.GetRandInt(min, max))
	q := utils.GetPrimeNumber(utils.GetRandInt(min, max))

	// 2) находим их произведение
	n := p * q

	// 3) вычисляем функцию Эйлера
	f := utils.EulerFunction(n)

	// 4) находим случайное число e, взаимно простое с f -
	// оно и будет ключом шифрования
	e := utils.GetRandomCoprimeNumber(f)

	// 5) находим число d, которое будет ключом дешифрования
	d := (f + 1) / e

	if d < 0 {
		d += n
	}

	return Key{e, n}, Key{d, n}, nil
}

// HandleRune - шифрование/дешифрование руны (зависит от ключа):
// C = (T ^ E) mod p*q
// T = (C ^ D) mod p*q
// где T — исходная руна, С — зашифрованная
func HandleRune(c rune, key Key) rune {
	result := c

	for i := 1; i < key.A; i++ {
		result = (c * result) % int32(key.N)
	}

	return result
}

// HandleMessage - засшифровать/дешифровать сообщение (зависит от ключа)
func HandleMessage(message string, key Key) string {
	runes := []rune(message)
	result := []rune{}

	for _, rune := range runes {
		result = append(result, HandleRune(rune, key))
	}

	return string(result)
}
