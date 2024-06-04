package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("are you ok?")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("最终结果：%q\n", out.String())
}
