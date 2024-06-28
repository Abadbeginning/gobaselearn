package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 屏障模式响应结构体
type BarrierRespone struct {
	Err    error
	Resp   string
	Status int
}

func doRequest(out chan<- BarrierRespone, url string) {
	res := BarrierRespone{
		// Err: nil,
		// Resp: "",
		// Status: 0,
	}

	client := http.Client{
		// 超时20s
		Timeout: time.Duration(20 * time.Second),
	}

	// 执行get请求
	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}

	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	// 将获取的结果放入通道
	out <- res
}

// 合并结果
func Barrier(urls []string) {
	requestNumber := len(urls)

	// in := make(chan BarrierRespone, requestNumber)
	in := make(chan BarrierRespone, requestNumber)
	response := make([]BarrierRespone, requestNumber)

	defer close(in)

	for _, url := range urls {
		go doRequest(in, url)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("error:", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}

	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Resp)
		}
	}

}

func main() {
	Barrier([]string{"https://www.baidu.com", "https://www.biying.com", "https://www.biying.com"})
}
