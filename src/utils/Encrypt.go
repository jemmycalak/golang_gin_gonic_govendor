package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
)

func Encryptor(texts []byte) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte("secret*#key#*for*#AES&encryption"))
	if err != nil {
		log.Println("error cipherBlock")
		return "", err
	}

	endcodeRawString := base64.StdEncoding.EncodeToString(texts)
	cipherText := make([]byte, aes.BlockSize+len(endcodeRawString))

	t1 := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, t1); err != nil {
		log.Println("error Readfull")
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(cipherBlock, t1)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(endcodeRawString))
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decryptor(texts string) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte("secret*#key#*for*#AES&encryption"))
	if err != nil {
		fmt.Println("cipherBlock")
		return "", err
	}
	decodeString, err := base64.StdEncoding.DecodeString(texts)
	if err != nil {
		fmt.Println("decodeString", texts)
		return "", err
	}
	if len(decodeString) < aes.BlockSize {
		fmt.Println("len decodeString")
		return "", errors.New("cipher text is too short")
	}
	t1 := decodeString[:aes.BlockSize]
	decodeString = decodeString[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherBlock, t1)
	cfb.XORKeyStream(decodeString, decodeString)
	decryptedString, err := base64.StdEncoding.DecodeString(string(decodeString))
	if err != nil {
		return "", err
	}

	return string(decryptedString), nil
}

// Takes in the encrypted string and returns the decrypted text
func Decrypt(encryptedString string) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte("secret*#key#*for*#AES&encryption"))
	if err != nil {
		return "", err
	}
	decodedString, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	if len(decodedString) < aes.BlockSize {
		return "", errors.New("cipher text is too short")
	}
	iv := decodedString[:aes.BlockSize]
	decodedString = decodedString[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherBlock, iv)
	cfb.XORKeyStream(decodedString, decodedString)
	decryptedString, err := base64.StdEncoding.DecodeString(string(decodedString))
	if err != nil {
		return "", err
	}
	return string(decryptedString), nil
}
