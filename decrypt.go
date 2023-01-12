package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"strings"
)

func decryptFile(filePath string) error {
	// Open the encrypted file
	encryptedFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	// Create the new decrypted file
	if !strings.HasSuffix(filePath, ".enc") {
		return fmt.Errorf("Invalid file extension %s expected .enc", filePath)
	}

	newFile, err := os.Create(strings.TrimSuffix(filePath, ".enc"))
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Read the key from the key file
	keyFile, err := os.Open(strings.TrimSuffix(filePath, ".enc") + ".key")
	if err != nil {
		return err
	}
	defer keyFile.Close()

	key := make([]byte, 32)
	_, err = keyFile.Read(key)
	if err != nil {
		return err
	}

	// Create the AES-256-CTR cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new CTR (counter mode)
	stream := cipher.NewCTR(block, make([]byte, aes.BlockSize))

	// Decrypt the contents of the encrypted file
	_, err = io.Copy(newFile, cipher.StreamReader{S: stream, R: encryptedFile})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Check if a file name was passed as a command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run decrypt.go [fileName]")
		os.Exit(1)
	}

	// Get the file name from the command-line argument
	fileName := os.Args[1]

	// Encrypt the file
	err := decryptFile(fileName)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		os.Exit(1)
	}

	fmt.Println("File successfully decrypted.")
}
