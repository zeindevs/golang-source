package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	magicHeader = "MYFT"
	version     = 1
)

var secretKey = []byte("e7bac1f5054756b1d1ba395244e29445")

func writeFile(filename string, data []byte) error {
	nonce, chipertext, err := encryptAESGCM(data)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.WriteString(magicHeader)
	buf.WriteByte(byte(version))
	buf.Write(nonce)
	binary.Write(buf, binary.BigEndian, uint32(len(chipertext)))
	buf.Write(chipertext)
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	header := make([]byte, 4)
	if _, err := io.ReadFull(file, header); err != nil {
		return nil, err
	}
	if string(header) != magicHeader {
		return nil, fmt.Errorf("invalid file format")
	}
	ver := make([]byte, 1)
	io.ReadFull(file, ver)
	nonce := make([]byte, 12)
	io.ReadFull(file, nonce)
	lenBuf := make([]byte, 4)
	io.ReadFull(file, lenBuf)
	length := binary.BigEndian.Uint32(lenBuf)
	chipertext := make([]byte, length)
	io.ReadFull(file, chipertext)
	return decryptAESGCM(nonce, chipertext)
}

// func xor(data, key []byte) []byte {
// 	out := make([]byte, len(data))
// 	for i := range data {
// 		out[i] = data[i] ^ key[i%len(key)]
// 	}
// 	return out
// }

func encryptAESGCM(plaintext []byte) (nonce, chipertext []byte, err error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, nil, err
	}
	nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	chipertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return nonce, chipertext, nil
}

func decryptAESGCM(nonce, chipertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, nonce, chipertext, nil)
}

func main() {
	filename := "example.myft"
	content := []byte("This is my secret data!")
	if err := writeFile(filename, content); err != nil {
		panic(err)
	}
	data, err := readFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read from file:", string(data))
}
