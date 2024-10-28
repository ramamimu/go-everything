package others

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	math_rand "math/rand"
	"testing"
	"time"
)

// Encrypt encrypts plaintext using AES-GCM with the provided key.
func Encrypt(plaintext, keyHex string) (string, error) {
	// Convert hex key to bytes and ensure it is 32 bytes for AES-256
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", errors.New("failed to decode key")
	}
	if len(keyBytes) != 32 {
		return "", errors.New("key must be 32 bytes (64 hex characters) for AES-256")
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Use GCM for authenticated encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	// Return the encrypted data as a base64-encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts an AES-GCM encrypted base64-encoded string with the provided key.
func Decrypt(encryptedText, keyHex string) (string, error) {
	// Convert hex key to bytes and ensure it is 32 bytes for AES-256
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", errors.New("failed to decode key")
	}
	if len(keyBytes) != 32 {
		return "", errors.New("key must be 32 bytes (64 hex characters) for AES-256")
	}

	// Decode the base64 string
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Use GCM for authenticated encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Split nonce and actual ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func TestEncDec(t *testing.T) {
	// A 256-bit key in hexadecimal format
	key := ""
	fmt.Println(len(key))
	message := ""

	// Encrypt the message
	encrypted, err := Encrypt(message, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	// Decrypt the message
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", decrypted)
}

const charset = "abcdefABCDEF0123456789"

func init() {
	math_rand.Seed(time.Now().UnixNano())
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[math_rand.Intn(len(charset))]
	}
	return string(b)
}

// the character should 0-9 and a-f (or A-F if uppercase).
func TestGenerateRandomeString(t *testing.T) {
	randomString := generateRandomString(64) // Generates a random string of length 64
	println(randomString)
}
