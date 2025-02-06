package main

import (
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/chacha20"
)

// CryptoProvider 提供不同加密算法的适配层
type CryptoProvider interface {
	NewEncryptStream(key, iv []byte) cipher.Stream
	NewDecryptStream(key, iv []byte) cipher.Stream
}

// AESCFBProvider 使用 AES CFB 模式
type AESCFBProvider struct{}

func (p *AESCFBProvider) NewEncryptStream(key, iv []byte) cipher.Stream {
	block, _ := aes.NewCipher(key)
	return cipher.NewCFBEncrypter(block, iv)
}

func (p *AESCFBProvider) NewDecryptStream(key, iv []byte) cipher.Stream {
	block, _ := aes.NewCipher(key)
	return cipher.NewCFBDecrypter(block, iv)
}

// AESCTRProvider 使用 AES CTR 模式
type AESCTRProvider struct{}

func (p *AESCTRProvider) NewEncryptStream(key, iv []byte) cipher.Stream {
	block, _ := aes.NewCipher(key)
	return cipher.NewCTR(block, iv)
}

func (p *AESCTRProvider) NewDecryptStream(key, iv []byte) cipher.Stream {
	block, _ := aes.NewCipher(key)
	return cipher.NewCTR(block, iv)
}

// ChaCha20Provider 使用 ChaCha20 加密
type ChaCha20Provider struct{}

func (p *ChaCha20Provider) NewEncryptStream(key, iv []byte) cipher.Stream {
	stream, _ := chacha20.NewUnauthenticatedCipher(key, iv)
	return stream
}

func (p *ChaCha20Provider) NewDecryptStream(key, iv []byte) cipher.Stream {
	stream, _ := chacha20.NewUnauthenticatedCipher(key, iv)
	return stream
}
