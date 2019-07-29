package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file %s. Error: %+v", fileName, err)
	}

	return file, nil
}

func WriteFile(writer *csv.Writer, record []string) error {
	if err := writer.Write(record); err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func GetParentDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed in getting current working directory ", err)
		return wd
	}

	return wd
}

func Encrypt(key, text []byte) (string, error) {
	var finalMsg string
	block, err := aes.NewCipher(key)
	if err != nil {
		return finalMsg, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return finalMsg, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	finalMsg = fmt.Sprintf("%0x", ciphertext)

	return finalMsg, nil
}

func Decrypt(strKey, str string) (string, error) {
	plainText := ""
	key := []byte(strKey)

	cipherText, err := hex.DecodeString(str)
	if err != nil {
		log.Println("Error in DecodeString")
		return plainText, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return plainText, err
	}

	if len(cipherText) < aes.BlockSize {
		log.Println(err)
		return plainText, fmt.Errorf("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	data, err := base64.StdEncoding.DecodeString(string(cipherText))
	if err != nil {
		log.Println(err)
		return plainText, err
	}
	plainText = string(data)
	return plainText, nil
}
