package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func main() {
	fmt.Println("Olá Mundo")
	resultado, err := Decode("Zh4U5AIUycyzeXZXSGADYOS/vmewlw/HMB2x5DlVF9oVQw==", "9ba4244eb81e11974f63a492c87c349e")
	fmt.Println(err)
	fmt.Println(resultado)
}

func Decode(ivCifra string, chave string) (string, error) {

	iv, cifra, err := extractIVAndCipherText(ivCifra)

	if err != nil {
		return "", errors.Wrap(err, "failed to decode IV and cipher text")
	}

	// create the AES cipher with the given secret key and IV
	keyBytes := []byte(chave)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", errors.Wrap(err, "failed to create AES cipher")
	}

	// decrypt the cipher text using GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "failed to create GCM cipher")
	}

	plainText, err := gcm.Open(nil, iv, cifra, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt cipher text")
	}

	return strings.TrimSpace(string(plainText)), nil
}

func extractIVAndCipherText(ivAndCipherText string) ([]byte, []byte, error) {
	ivAndCipherTextBytes, err := base64.StdEncoding.DecodeString(ivAndCipherText)
	if err != nil {
		return nil, nil, err
	}

	ivSize := 12 // assuming 12 bytes IV for AES GCM mode
	if len(ivAndCipherTextBytes) < ivSize {
		return nil, nil, errors.New("invalid IV and CipherText")
	}

	iv := make([]byte, ivSize)
	copy(iv, ivAndCipherTextBytes[:ivSize])

	cipherText := make([]byte, len(ivAndCipherTextBytes)-ivSize)
	copy(cipherText, ivAndCipherTextBytes[ivSize:])

	fmt.Println("IV:", base64.StdEncoding.EncodeToString(iv))
	fmt.Println("CipherText:", string(cipherText))

	return iv, cipherText, nil
}
