package codec

import "encoding/xml"

type CDATA struct {
	Text string `xml:",cdata"`
}

// 消息基础结构
type BaseMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   CDATA
	MsgType      CDATA
}

type TextMsg struct {
	BaseMsg
	Content CDATA
}

type ImageMsg struct {
	BaseMsg
	PicUrl    CDATA
	MediaId   CDATA
	MsgId     string `xml:"MsgId"`
	MsgDataId string `xml:"MsgDataId"`
	Idx       string `xml:"Idx"`
}

type VoiceMsg struct {
	BaseMsg
	MediaId    CDATA
	Format     CDATA
	MsgId      string `xml:"MsgId"`
	MsgDataId  string `xml:"MsgDataId"`
	Idx        string `xml:"Idx"`
	MediaId16K CDATA
}
