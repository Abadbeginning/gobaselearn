package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	input := "see you lala"
	md5_n := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5_n, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
    sha_512_256 := sha512.Sum512_256([]byte(input))
    hmac512 := hmac.New(sha512.New, []byte("secret"))
    hmac512.Write([]byte(input))
	fmt.Println("md5 --> ", md5.Sum(nil))
    fmt.Println("sha256 --> ", sha_256.Sum(nil))
    fmt.Println("sha512 --> ", sha_512.Sum(nil))
    fmt.Println("sha512_256 --> ", sha_512_256)
    fmt.Println("hmac512 --> ", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}
