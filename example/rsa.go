package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

type RsaLib struct {
	PubKey []byte
	PriKey []byte
}

// 获取文件内容
func GetFileContent(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Func:GetFileContent();Go:Open() Err:", err)
		return nil
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Func:GetFileContent();Go:ReadAll() Err:", err)
		return nil
	}

	return content
}

// 公钥加密
func (r *RsaLib) PublicKeyEncrypt(text []byte) string {
	block, _ := pem.Decode(r.PubKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := pubInterface.(*rsa.PublicKey)
	encryptedBytes, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, text)
	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

// 私钥解密
func (r *RsaLib) PrivateKeyEncrypt(text string) string {
	block, _ := pem.Decode(r.PriKey)
	res, _ := base64.StdEncoding.DecodeString(text)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, res)
	return string(plainText)
}

func main() {
	message := []byte("B6mG5wy8cwSgbqwWeFrJFWxugKAQDzD3ex5Lkv6PtypSriRWXFWUxtb3wXNVzZed")

	PubKeyContent := GetFileContent("key/ios/public_key.pem")
	PriKeyContent := GetFileContent("key/ios/private_key.pem")
	rsaLib := RsaLib{
		PubKey: PubKeyContent,
		PriKey: PriKeyContent,
	}

	// 加密
	cipherText := rsaLib.PublicKeyEncrypt(message)
	fmt.Println("加密结果：", cipherText)

	// 解密
	result := rsaLib.PrivateKeyEncrypt(cipherText)
	fmt.Println("解密结果：", result)
}
