package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
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

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("generate IV: %w", err)
	}

	cipher.NewCBCEncrypter(block, iv).
		CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedString string) (uint64, error) {
	ciphertext, err := hex.DecodeString(encryptedString)
	if err != nil {
		return 0, fmt.Errorf("invalid ciphertext hex: %w", err)
	}

	// Minimal: satu blok IV dan satu blok data.
	if len(ciphertext) < 2*aes.BlockSize {
		return 0, fmt.Errorf("ciphertext too short")
	}

	if (len(ciphertext)-aes.BlockSize)%aes.BlockSize != 0 {
		return 0, fmt.Errorf("invalid ciphertext length")
	}

	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		return 0, fmt.Errorf("invalid encryption key: %w", err)
	}

	iv := ciphertext[:aes.BlockSize]
	plaintext := ciphertext[aes.BlockSize:]

	cipher.NewCBCDecrypter(block, iv).
		CryptBlocks(plaintext, plaintext)

	// Validasi seluruh byte padding PKCS#7 sebelum dipotong.
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
