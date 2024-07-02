package http

import (
	. "wechat-server/http"

	"github.com/gin-gonic/gin"
)

type WechatMsgService struct {
	router *gin.Engine
}

func NewWeChatMsgService(tk string) *WechatMsgService {
	wcms := &WechatMsgService{}

	token = tk

	// 设置路由
	wcms.router = gin.Default()
	wcms.router.GET("/", WechatGet)
	wcms.router.POST("/", WechatPost)

	return wcms
}

func (wcms *WechatMsgService) RegisterDistributorMap(m ContentHanlderMap) {
	msgDistributor.Map = m
}

func (wcms *WechatMsgService) Start() {
	wcms.router.Run(":80")
}
