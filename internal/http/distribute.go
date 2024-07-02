package http

import (
	. "wechat-server/http"
	. "wechat-server/internal/codec"

	"github.com/gin-gonic/gin"
)

type Distributor struct {
	Serial Serializer
	Map    ContentHanlderMap
}

func (d *Distributor) HandleRequest(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		// 空消息体
		return
	}

	base := &BaseMsg{}
	err = d.Serial.Unmarshal(body, base)
	if err != nil {
		// 解析消息体出错
		return
	}

	// 责任链的next是通过依赖注入的方式获取
	next := d.Map.GetHandler(base.MsgType.Text)
	if next != nil {
		next.Parser(body, c)
	} else {
		// 业务层没有实现该消息类型的Hanlder，流程结束
	}
}
