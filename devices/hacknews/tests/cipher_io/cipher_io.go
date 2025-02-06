package main

import (
	"io"
)

// CipherIO 适配加密/解密，封装读写操作
type CipherIO struct {
	rw       io.ReadWriter
	provider CryptoProvider
	key      []byte
	iv       []byte
}

// NewCipherIO 创建 CipherIO 适配层
func NewCipherIO(rw io.ReadWriter, provider CryptoProvider, key, iv []byte) *CipherIO {
	return &CipherIO{
		rw:       rw,
		provider: provider,
		key:      key,
		iv:       iv,
	}
}

// Write 加密并写入数据
func (c *CipherIO) Write(p []byte) (int, error) {
	encryptStream := c.provider.NewEncryptStream(c.key, c.iv)
	dst := make([]byte, len(p))
	encryptStream.XORKeyStream(dst, p)
	return c.rw.Write(dst)
}

// Read 读取并解密数据
func (c *CipherIO) Read(p []byte) (int, error) {
	n, err := c.rw.Read(p)
	if n > 0 {
		decryptStream := c.provider.NewDecryptStream(c.key, c.iv)
		decryptStream.XORKeyStream(p[:n], p[:n])
	}
	return n, err
}
