package main

import "fmt"

type Programe struct {
	Name string
}

func (p Programe) Run() {
	fmt.Println("Run")
}

type SkillInterface interface {
	Run()
}

func main() {
  var pro Programe
  var a SkillInterface = pro
  a.Run()

}
