package main

import (
	"fmt"
	"time"
)

// 放茶叶
func putInTea() <-chan string {
  vegetables := make(chan string)
  go func() {
    time.Sleep(5 * time.Second)
    vegetables <- "茶叶茶叶放茶杯"
  }()
  return vegetables
}

// 烧水
func boilWater() <-chan string {
  water := make(chan string)
  go func() {
    time.Sleep(5 * time.Second)
    water <- "水水水烧水"
  }()

  return water
}

func main() {
  teaCh := putInTea()
  waterCh := boilWater()
  time.Sleep(2 * time.Second)
  
  fmt.Println("2s休息")

  tea := <- teaCh
  water := <- waterCh
  fmt.Println(tea, water)
  
}
