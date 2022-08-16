/*
Create: 2022/8/5
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"errors"
	"net"

	"github.com/JJApplication/fushin/server/uds"
	"github.com/JJApplication/fushin/utils/json"
)

type UDSClient struct {
	Addr        string
	conn        net.Conn
	MaxRecvSize int
}

const MaxReadSize = 4096

// Dial 连接指定的unix domain
// 在初始化后被调用连接
func (c *UDSClient) Dial() error {
	if c.conn == nil {
		conn, err := net.Dial("unix", c.Addr)
		if err != nil {
			return err
		}
		c.conn = conn
		return nil
	}
	return errors.New(ErrAlreadyDial)
}

// Send 发送请求 不接收服务端的返回数据
// 发送任意数据请使用SendRaw
func (c *UDSClient) Send(req uds.Req) error {
	if c.conn == nil {
		return errors.New(ErrDialClosed)
	}
	data, err := json.Json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(data)
	return err
}

// SendRaw 发送raw的数据
func (c *UDSClient) SendRaw(data string) error {
	if c.conn == nil {
		return errors.New(ErrDialClosed)
	}
	_, err := c.conn.Write([]byte(data))
	return err
}

// SendWithRes 发送数据并尝试从服务端接收数据
func (c *UDSClient) SendWithRes(req uds.Req) (uds.Res, error) {
	var res uds.Res
	if c.conn == nil {
		return uds.Res{}, errors.New(ErrDialClosed)
	}
	data, err := json.Json.Marshal(req)
	if err != nil {
		return uds.Res{}, err
	}
	_, err = c.conn.Write(data)
	var buf = make([]byte, maxSize(c.MaxRecvSize))
	count, err := c.conn.Read(buf)
	if err != nil {
		return uds.Res{}, err
	}
	if err = json.Json.Unmarshal(buf[:count], &res); err != nil {
		return uds.Res{}, err
	}

	return res, nil
}

// SendRawWithRes 发送raw数据并接收返回
func (c *UDSClient) SendRawWithRes(data string) ([]byte, error) {
	if c.conn == nil {
		return nil, errors.New(ErrDialClosed)
	}
	_, err := c.conn.Write([]byte(data))
	var buf = make([]byte, maxSize(c.MaxRecvSize))
	count, err := c.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:count], err
}

func (c *UDSClient) Close() error {
	return c.conn.Close()
}

func maxSize(size int) int {
	if size <= 0 {
		return MaxReadSize
	}
	return size
}
