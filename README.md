# AES-256-CTR Encryption and Decryption in Go
This is an implementation of AES-256-CTR encryption and decryption in Go. The program uses the Go standard library's `crypto/aes` and `crypto/cipher` packages to perform the encryption and decryption.

This package provides an example implementation of AES-256-CTR file encryption and decryption in Go. The package includes the following functions:

## Encryption
`encryptFile` encrypts the contents of the file at the specified file path and saves the result to a new file with ".enc" appended to the original file name, and also creates a keyfile with the name filename.key

## Decryption
`decryptFile` decrypts the contents of a file with a ".enc" extension at the specified file path and saves the result to a new file with ".enc" removed from the original file name, the function read the key from the keyfile named filename.key

