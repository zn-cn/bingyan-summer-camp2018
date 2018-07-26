

//des 加密算法
package models

import (
	"fmt"
	"bytes"
	"crypto/des"
	"crypto/cipher"	
	"encoding/base64"	
)

func GetDes(value string) (string){
	key := []byte("sfe023f_")

	result, err := DesEncrypt([]byte(value), key)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(result))

	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(origData))

	return base64.StdEncoding.EncodeToString(result)
}

//des加密
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		 return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	 // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//DES解密代码如下：
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		 return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool{
		return r == rune(0)
	})
}

//以上代码使用DES加密（des.NewCipher），加密模式为CBC（cipher.NewCBCEncrypter(block, key)），填充方式PKCS5Padding，该函数的代码如下：
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//可见，解密无非是调用cipher.NewCBCDecrypter，最后unpadding，其他跟加密几乎一样。相应的PKCS5UnPadding：

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


