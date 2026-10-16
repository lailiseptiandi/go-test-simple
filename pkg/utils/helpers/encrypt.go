package helpers

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func Encrypt(valueEncrypt uint64) (string, error) {
	plaintext := []byte(strconv.FormatUint(valueEncrypt, 10))

	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		return "", fmt.Errorf("invalid encryption key: %w", err)
	}

	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	plaintext = append(
		plaintext,
		bytes.Repeat([]byte{byte(padding)}, padding)...,
	)

	ciphertext := make([]byte, len(plaintext))
	block.Encrypt(ciphertext, plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedString string) (uint64, error) {
	ciphertext, err := hex.DecodeString(encryptedString)
	if err != nil {
		return 0, fmt.Errorf("invalid ciphertext hex: %w", err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return 0, fmt.Errorf("invalid ciphertext length")
	}

	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		return 0, fmt.Errorf("invalid encryption key: %w", err)
	}

	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)

	padding := int(plaintext[len(plaintext)-1])
	if padding < 1 || padding > aes.BlockSize || padding > len(plaintext) {
		return 0, fmt.Errorf("invalid ciphertext padding")
	}

	for _, value := range plaintext[len(plaintext)-padding:] {
		if value != byte(padding) {
			return 0, fmt.Errorf("invalid ciphertext padding")
		}
	}

	plaintext = plaintext[:len(plaintext)-padding]

	value, err := strconv.ParseUint(string(plaintext), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid decrypted uint64: %w", err)
	}

	return value, nil
}
