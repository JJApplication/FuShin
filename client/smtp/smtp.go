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
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/gomail.v2"
)

// go的smtp

type SmtpClient struct {
	Sender     string   // 发件邮箱
	NickSender string   // 发件邮箱的别名 为空时与Sender一致
	PassWord   string   // 发件密码
	SmtpHost   string   // smtp提供商
	SmtpPort   int      // smtp端口
	To         []string // 主送
	Cc         []string // 抄送
	Bcc        []string // 暗送
}

// Send 纯文本
// subject 主题
// message 正文
// attach 附件
func (s *SmtpClient) Send(subject, message string, attach []string) error {
	if !s.check() {
		return errors.New(ErrCheck)
	}
	if s.To == nil || len(s.To) == 0 {
		return errors.New(ErrNoTo)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", s.getFrom())
	m.SetHeader("To", s.To...)
	if len(s.Cc) > 0 {
		m.SetHeader("Cc", s.Cc...)
	}
	if len(s.Bcc) > 0 {
		m.SetHeader("Bcc", s.Bcc...)
	}
	m.SetHeader("Subject", subject)
	m.SetBody(Text, message)
	s.setAttach(m, attach)
	dial := gomail.NewDialer(s.SmtpHost, s.SmtpPort, s.Sender, s.PassWord)
	return dial.DialAndSend(m)
}

// SendHtml html
// data 文件或者html文本
// isFile 为true时data以html文件的方式解析
func (s *SmtpClient) SendHtml(subject, data string, isFile bool, attach []string) error {
	if !s.check() {
		return errors.New(ErrCheck)
	}
	if s.To == nil || len(s.To) == 0 {
		return errors.New(ErrNoTo)
	}
	if data == "" {
		return errors.New(ErrNoHtml)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", s.getFrom())
	m.SetHeader("To", s.To...)
	m.SetHeader("Subject", subject)
	if isFile {
		content, err := getHtmlBody(data)
		if err != nil {
			return err
		}
		m.SetBody(Html, content)
	} else {
		m.SetBody(Html, data)
	}
	s.setAttach(m, attach)
	dial := gomail.NewDialer(s.SmtpHost, s.SmtpPort, s.Sender, s.PassWord)
	return dial.DialAndSend(m)
}

// SendContext 带上下文的发送
func (s *SmtpClient) SendContext(context context.Context, subject, message string, attach []string) error {
	if context == nil {
		return errors.New(ErrContextNil)
	}
	ch := make(chan error, 1)
	go func() {
		ch <- s.Send(subject, message, attach)
	}()
	select {
	case <-context.Done():
		return context.Err()
	case err := <-ch:
		return err
	}
}

// SendHtmlContext 带上下文的发送
func (s *SmtpClient) SendHtmlContext(context context.Context, subject, data string, isFile bool, attach []string) error {
	if context == nil {
		return errors.New(ErrContextNil)
	}
	ch := make(chan error, 1)
	go func() {
		ch <- s.SendHtml(subject, data, isFile, attach)
	}()
	select {
	case <-context.Done():
		return context.Err()
	case err := <-ch:
		return err
	}
}

// Try 尝试连接
func (s *SmtpClient) Try() error {
	dial := gomail.NewDialer(s.SmtpHost, s.SmtpPort, s.Sender, s.PassWord)
	closer, err := dial.Dial()
	if err != nil {
		return err
	}
	return closer.Close()
}

// 检查host
func (s *SmtpClient) check() bool {
	if s.SmtpPort <= 0 || s.Sender == "" || s.SmtpHost == "" {
		return false
	}
	return true
}

func (s *SmtpClient) getHost() string {
	return fmt.Sprintf("%s:%d", s.SmtpHost, s.SmtpPort)
}

func (s *SmtpClient) getFrom() string {
	if s.NickSender == "" {
		return s.Sender
	}
	return fmt.Sprintf("%s<%s>", s.NickSender, s.Sender)
}

func (s *SmtpClient) setAttach(m *gomail.Message, attach []string) {
	for _, f := range attach {
		if f != "" {
			m.Attach(f)
		}
	}
}

func getHtmlBody(f string) (string, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
