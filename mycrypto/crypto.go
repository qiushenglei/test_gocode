package mycrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/safe"
	"io"
	"math/big"
	"os"
)

func RSA() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("密钥对生成失败:", err)
		return
	}

	// 将私钥保存到文件
	err = savePrivateKey(privateKey, safe.Path("/cmd/cryptocmd/private.pem"))
	if err != nil {
		fmt.Println("私钥保存失败:", err)
		return
	}

	// 将公钥保存到文件
	err = savePublicKey(privateKey.PublicKey, safe.Path("/cmd/cryptocmd/public.pem"))
	if err != nil {
		fmt.Println("公钥保存失败:", err)
		return
	}

	// 加密消息
	message := []byte("Hello, RSA!")
	ciphertext, err := rsaEncrypt(message, privateKey)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}

	fmt.Println("加密后的消息:", string(ciphertext))

	// 解密消息
	plaintext, err := rsaDecrypt(ciphertext, privateKey)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}

	fmt.Println("解密后的消息:", string(plaintext))
}

// 保存私钥到文件
func savePrivateKey(privateKey *rsa.PrivateKey, filename string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}

	keyFile, err := openFile(filename, "wb")
	if err != nil {
		return err
	}
	defer keyFile.Close()
	pem.Encode(keyFile, pemBlock)
	return nil
}

// 保存公钥到文件
func savePublicKey(publicKey rsa.PublicKey, filename string) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	}
	keyFile, err := openFile(filename, "wb")
	if err != nil {
		return err
	}
	defer keyFile.Close()
	pem.Encode(keyFile, pemBlock)
	return nil
}

func rsaEncrypt(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func rsaDecrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func openFile(filename string, mode string) (file *os.File, err error) {
	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CryptRand() {
	var b io.Reader
	b = bytes.NewReader([]byte("12384567"))
	for i := 0; i < 6; i++ {
		//res, err := rand.Int(rand.Reader, big.NewInt(10))
		res, err := rand.Int(b, big.NewInt(9))
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

}

func CryptRand1() {
	str := "1234567890"
	for i := 0; i < 6; i++ {
		res, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			panic(err)
		}
		num := res.Int64()
		fmt.Println(num)
		fmt.Println(str[num])
	}
}
