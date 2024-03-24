package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/nacl/secretbox"
)

// generateKey generates a random 32-byte key
func generateKey() *[32]byte {
	key := new([32]byte)
	_, err := rand.Read(key[:])
	if err != nil {
		log.Fatal("Error generating key:", err)
	}
	return key
}

// encryptMessage encrypts a message using a key
func encryptMessage(message string, key *[32]byte) ([]byte, error) {
	// Nonce should never be reused with the same key
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, err
	}

	encrypted := secretbox.Seal(nonce[:], []byte(message), &nonce, key)
	return encrypted, nil
}

// decryptMessage decrypts an encrypted message using a key
func decryptMessage(encrypted []byte, key *[32]byte) (string, error) {
	var nonce [24]byte
	copy(nonce[:], encrypted[:24])

	decrypted, ok := secretbox.Open(nil, encrypted[24:], &nonce, key)
	if !ok {
		return "", fmt.Errorf("decryption failed")
	}

	return string(decrypted), nil
}

func main() {
	// Generate a random key for Alice and Bob
	// aliceKey := generateKey()
	bobKey := generateKey()

	// Alice sends a message to Bob
	aliceMessage := "Hello, Bob, How are you today?"
	encryptedMessage, err := encryptMessage(aliceMessage, bobKey)
	if err != nil {
		log.Fatal("Error encrypting message:", err)
	}

	// Bob receives and decrypts the message
	decryptedMessage, err := decryptMessage(encryptedMessage, bobKey)
	if err != nil {
		log.Fatal("Error decrypting message:", err)
	}

	// Display the results
	fmt.Printf("Original message: %s\n", aliceMessage)
	fmt.Printf("Encrypted message: %s\n", base64.StdEncoding.EncodeToString(encryptedMessage))
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)
}
