package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func encryptFile(filePath string) error {
	// Open the original file
	originalFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	// Create the new encrypted file
	newFile, err := os.Create(filePath + ".enc")
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Generate a new 256-bit (32-byte) AES key
	key := make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, key)
	if err != nil {
		return err
	}

	// Print the key for debugging
	fmt.Println("Encryption key:", key)

	// Save the key to a binary file
	keyFile, err := os.Create(filePath + ".key")
	if err != nil {
		return err
	}
	defer keyFile.Close()

	//Please also make sure that you handle the key file appropriately, as the encryption key is essential for decryption.
	//It's recommended to save it in a secure place and make sure that only authorized personnel can access it,
	//as well as protect it from un-authorized access or modification.
	_, err = keyFile.Write(key)
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

	// Encrypt the contents of the original file
	_, err = io.Copy(newFile, cipher.StreamReader{S: stream, R: originalFile})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Check if a file name was passed as a command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run encrypt.go [fileName]")
		os.Exit(1)
	}

	// Get the file name from the command-line argument
	fileName := os.Args[1]

	// Encrypt the file
	err := encryptFile(fileName)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		os.Exit(1)
	}

	fmt.Println("File successfully encrypted.")
}
