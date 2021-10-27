package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	e.From = "kim <494040246@qq.com>"
	e.To = []string{"494040246@qq.com"}
	e.Subject = "发现惊天大秘密！"
	e.Text = []byte("Hello World")
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "494040246@qq.com", "klbmkizobueacbbe", "smtp.qq.com"))
	if err != nil {
		panic(err)
	}

}
