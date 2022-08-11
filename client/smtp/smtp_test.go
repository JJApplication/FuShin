/*
Create: 2022/8/10
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package smtp
package smtp

import (
	"context"
	"testing"
	"time"
)

func TestSmtpSend(t *testing.T) {
	smtp := SmtpClient{
		Sender:     "mail@163.com",
		NickSender: "mail",
		PassWord:   "pwd",
		SmtpHost:   "smtp.163.com",
		SmtpPort:   465,
		To:         []string{"liaorenj@gmail.com"},
	}
	err := smtp.Send("Theme", "Hello World", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("test smtp send text success")
}

func TestSmtpSendHtml(t *testing.T) {
	smtp := SmtpClient{
		Sender:     "mail@163.com",
		NickSender: "mail",
		PassWord:   "pwd",
		SmtpHost:   "smtp.163.com",
		SmtpPort:   465,
		To:         []string{"liaorenj@gmail.com"},
	}
	body := `<html><body><h1 style="color: red">Hello World</h1></body></html>`
	err := smtp.SendHtml("主题", body, false, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("test smtp send html success")
}

func TestSmtpSendContext(t *testing.T) {
	smtp := SmtpClient{
		Sender:     "mail@163.com",
		NickSender: "mail",
		PassWord:   "pwd",
		SmtpHost:   "smtp.163.com",
		SmtpPort:   465,
		To:         []string{"liaorenj@gmail.com"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := smtp.SendContext(ctx, "", "", nil)
	if err != nil {
		t.Error(err)
	}
	select {
	case <-ctx.Done():
		t.Log(ctx.Err())
		t.Log("test smtp with context success")
	default:
		t.Log("done")
	}
}

func TestSmtpSendAttach(t *testing.T) {
	smtp := SmtpClient{
		Sender:     "mail@163.com",
		NickSender: "mail",
		PassWord:   "pwd",
		SmtpHost:   "smtp.163.com",
		SmtpPort:   465,
		To:         []string{"liaorenj@gmail.com"},
	}

	err := smtp.Send("text with Attach", "see the attach file", []string{"test.zip"})
	if err != nil {
		t.Error(err)
	}
}
