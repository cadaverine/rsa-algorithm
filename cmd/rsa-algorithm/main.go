package main

import (
	"fmt"
	"log"

	"github.com/cadaverine/rsa-algorithm/rsa"
)

func main() {
	publicKey, privateKey, err := rsa.GetRandomRSAKeys()
	if err != nil {
		log.Fatal(err)
	}

	message := "Hello, RSA!"

	// зашифровываем сообщение:
	encripted := rsa.HandleMessage(message, publicKey)

	// расшифровываем сообщение:
	decripted := rsa.HandleMessage(encripted, privateKey)

	fmt.Println("Исходное сообщение: ", message)
	fmt.Println("Зашифрованное сообщение: ", encripted)
	fmt.Println("Расшифрованное сообщение: ", decripted)
}
