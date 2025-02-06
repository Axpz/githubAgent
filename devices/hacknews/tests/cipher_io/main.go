package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	// 生成密钥和 IV
	key := make([]byte, 32) // ChaCha20 需要 32 字节密钥
	iv := make([]byte, 12)  // ChaCha20 推荐 12 字节 IV（Nonce）
	rand.Read(key)
	rand.Read(iv)

	// 选择加密 Provider
	provider := &ChaCha20Provider{}

	// 创建加密文件
	file, err := os.Create("encrypted.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 先写入 IV 以便解密时读取
	file.Write(iv)

	// 创建 CipherIO 进行加密写入
	cipherWriter := NewCipherIO(file, provider, key, iv)
	cipherWriter.Write([]byte("Hello, Encrypted World!"))

	// 读取加密数据并解密
	file, err = os.Open("encrypted.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件头的 IV
	ivHeader := make([]byte, 12)
	_, err = file.Read(ivHeader)
	if err != nil {
		panic("Failed to read IV from file")
	}

	// 创建 CipherIO 进行解密读取
	cipherReader := NewCipherIO(file, provider, key, ivHeader)
	buf := make([]byte, 1024)
	n, _ := cipherReader.Read(buf)
	fmt.Println("Decrypted:", string(buf[:n]))
}
