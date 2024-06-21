package main

import (
	"crypto/des"
	"fmt"
)

func main() {
  key := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd}
  cipherBlock, err :=  des.NewCipher(key)
  if err != nil {
    fmt.Println("error")
  }
  src := []byte("see you lala")
  encrptDst := make([]byte, len(src))
  cipherBlock.Encrypt(encrptDst, src)
  fmt.Println("DES --> ", encrptDst)
  plainDst := make([]byte, len(encrptDst))
  cipherBlock.Decrypt(plainDst, encrptDst)
  fmt.Println(plainDst)
}
