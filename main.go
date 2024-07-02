package main

import (
	. "wechat-server/http"
	. "wechat-server/internal/http"
)

func main() {
	// 注册Token
	s := NewWeChatMsgService("your_token")

	// 注册自己的业务代码
	// 当前业务代码为Echo应答
	s.RegisterDistributorMap(&EchoMap{
		Map: map[string]ContentHandler{
			"text":  &EchoText{},
			"image": &EchoImage{},
			"voice": &EchoVoice{},
		},
	})

	// 启动服务
	s.Start()
}
