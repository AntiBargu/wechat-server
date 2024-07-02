package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	. "wechat-server/internal/codec"

	. "wechat-server/http"

	"github.com/gin-gonic/gin"
)

type EchoText struct {
}

func (*EchoText) Parser(body []byte, c *gin.Context) {
	msg := &TextMsg{}
	err := xml.Unmarshal(body, msg)
	if err != nil {
		// 解析消息体出错
		return
	}

	// msg.CreateTime.Text = strconv.FormatInt(time.Now().Unix(), 10)
	msg.FromUserName, msg.ToUserName = msg.ToUserName, msg.FromUserName
	c.XML(http.StatusOK, msg)
}

type EchoImage struct {
}

func (*EchoImage) Parser(body []byte, c *gin.Context) {
	msg := &ImageMsg{}
	err := xml.Unmarshal(body, msg)
	if err != nil {
		// 解析消息体出错
		fmt.Println("消息解析出错", err)
		return
	}

	msg.FromUserName, msg.ToUserName = msg.ToUserName, msg.FromUserName
	c.XML(http.StatusOK, msg)
}

type EchoVoice struct {
}

func (*EchoVoice) Parser(body []byte, c *gin.Context) {
	msg := &VoiceMsg{}
	err := xml.Unmarshal(body, msg)
	if err != nil {
		// 解析消息体出错
		fmt.Println("消息解析出错", err)
		return
	}

	msg.FromUserName, msg.ToUserName = msg.ToUserName, msg.FromUserName
	c.XML(http.StatusOK, msg)
}

type EchoMap struct {
	Map map[string]ContentHandler
}

func (em *EchoMap) GetHandler(content string) ContentHandler {
	return em.Map[content]
}

func (em *EchoMap) SetHandler(content string, contentHandler ContentHandler) {
	em.Map[content] = contentHandler
}
