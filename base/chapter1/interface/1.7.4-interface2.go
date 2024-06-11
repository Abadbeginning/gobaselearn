package main

import "fmt"

type Message interface {
  sending()
}

type Email struct {
  From string
  To string
  Subject string
  Body string
}

func (e *Email) sending() {
  fmt.Printf("发送email %s < %s > \n", e.From, e.To)
}


type SMS struct {
  From string
  To string
  Body string
}


func (s *SMS) sending() {
  fmt.Printf("发送sms %s < %s > \n", s.From, s.To)
}

func sendMes(m Message) {
  m.sending()
}

func main() {
  email := Email{"张三", "李四", "邮件标题", "邮件内容"}
  sms := SMS{"小张", "小四", "短信内容"}
  sendMes(&email)
  sendMes(&sms)
}
