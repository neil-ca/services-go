package main

import (
	"github.com/Neil-uli/Restful-go/encryptString/utils"
	"log"
)

// AES keys should be of length 16, 24, 32
func main() {
	key := "111023043350789514532147"
	message := "I am a message"
	log.Println("Original message: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("decrypted message: ", decryptedString)
}
