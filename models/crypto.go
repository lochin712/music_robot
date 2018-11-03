package models

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
)

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ivDefValue = "0102030405060708"
	modulus    = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	nonce      = "0CoJUm6Qyw8W8jud"
	pubKey     = "010001"
)

// 加密参考：https://github.com/cosven/cosven.github.io/issues/30
func Encrypt(text string) (enText string, encSecKey string) {
	secKey := CreateSecretKey(16)
	fmt.Println(secKey)
	enText = AesEncrypt(AesEncrypt(text, nonce), secKey)
	encSecKey = RsaEncrypt(secKey, pubKey, modulus)

	return
}

func AesEncrypt(plaintextStr string, keyStr string) string {
	plaintext := []byte(plaintextStr)
	key := []byte(keyStr)
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	iv := []byte(ivDefValue)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func CreateSecretKey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RsaEncrypt(text string, pubKey string, modulus string) string {
	text = ReverseString(text)              // 反转字符串
	text = hex.EncodeToString([]byte(text)) //hex转码
	p := new(big.Int)
	p.SetString(pubKey, 16)
	m := new(big.Int)
	m.SetString(modulus, 16)
	ts := new(big.Int)
	res, _ := ts.SetString(text, 16)
	modpow := ModPow(res, p, m)
	lochin := fmt.Sprintf("%X", modpow)
	return ZFill(lochin, 256)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func ZFill(str string, size int) string {
	for i := len(str); i < size; i++ {
		str = "0" + str
	}
	return str
}

// 幂模运算 （快速幂算法：http://www.jb51.net/article/54947.htm）
func ModPow(a, b, c *big.Int) *big.Int {
	ans := big.NewInt(1)
	a = a.Mod(a, c)
	for {
		if b.Sign() != 1 {
			break
		}
		one := big.NewInt(1)
		two := big.NewInt(2)
		temp1 := new(big.Int)
		temp2 := new(big.Int)
		if temp2.Sub(temp1.Mod(b, two), one).Sign() == 0 {
			ans = ans.Mod(ans.Mul(ans, a), c)
		}
		b = b.Quo(b, two)
		a = a.Mod(a.Mul(a, a), c)
	}
	return ans
}
