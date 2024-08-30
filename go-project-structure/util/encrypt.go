package util

import (
	"crypto/aes"
	"encoding/hex"
)

func EncryptAES(plaintext string, password string) (string, error) {
  c, err := aes.NewCipher([]byte(password))
  if err != nil {
    return "", err
  }
  out := make([]byte, len(plaintext))
  c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out), nil
}

func DecryptAES(encryptedText string, password string) (string, error) {
  ct, _ := hex.DecodeString(encryptedText)
  c, err := aes.NewCipher([]byte(password))
  if err != nil {
    return "",err
  }
  buf := make([]byte, len(ct))
  c.Decrypt(buf, ct)
  return string(buf), nil
}
