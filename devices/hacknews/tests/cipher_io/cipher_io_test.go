package main

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestCipherIO(t *testing.T) {
	providers := []CryptoProvider{
		&AESCFBProvider{},
		&AESCTRProvider{},
		&ChaCha20Provider{},
	}

	key := make([]byte, 32) // 生成 32 字节密钥
	iv := make([]byte, 16)  // 生成 16 字节 IV（AES CFB/CTR 需要 16 字节，ChaCha20 需要 12 字节）
	rand.Read(key)
	rand.Read(iv)

	for _, provider := range providers {
		providerName := getProviderName(provider)
		t.Run("Testing "+providerName, func(t *testing.T) {
			plaintext := []byte("Hello, Secure World!")

			// 加密
			var buf bytes.Buffer
			cipherWriter := NewCipherIO(&buf, provider, key, iv[:getIVSize(provider)])
			_, err := cipherWriter.Write(plaintext)
			if err != nil {
				t.Fatalf("[%s] 加密失败: %v", providerName, err)
			}

			// 解密
			// cipherReader := NewCipherIO(&buf, provider, key, iv[:getIVSize(provider)])
			cipherReader := cipherWriter
			decrypted := make([]byte, len(plaintext))
			_, err = cipherReader.Read(decrypted)
			if err != nil {
				t.Fatalf("[%s] 解密失败: %v", providerName, err)
			}

			// 验证
			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("[%s] 解密数据不匹配: 期望 %s, 但得到 %s", providerName, plaintext, decrypted)
			}
		})
	}
}

// 获取加密模式名称
func getProviderName(p CryptoProvider) string {
	switch p.(type) {
	case *AESCFBProvider:
		return "AES-CFB"
	case *AESCTRProvider:
		return "AES-CTR"
	case *ChaCha20Provider:
		return "ChaCha20"
	default:
		return "Unknown"
	}
}

// 获取适用于不同算法的 IV 长度
func getIVSize(p CryptoProvider) int {
	switch p.(type) {
	case *AESCFBProvider, *AESCTRProvider:
		return 16 // AES 块大小固定 16 字节
	case *ChaCha20Provider:
		return 12 // ChaCha20 推荐 12 字节 Nonce
	default:
		return 16
	}
}
