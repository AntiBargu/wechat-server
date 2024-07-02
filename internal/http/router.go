package http

import (
	. "wechat-server/internal/codec"

	"github.com/gin-gonic/gin"
)

var token string

// 这里不会定义SetNext方法
// 参考自 C++ Core Guidelines C.131: Avoid trivial getters and setters
// 因为接口类的指针是invariant，故在派生类中实现传递职责的处理
type msgHandler interface {
	HandleRequest(*gin.Context)
}

// Go是静态多态，这里就不把GetHandler定义为MsgHandler接口类
// 下面的全局对象也可以设置为单例
// GET方法的处理对象
var getHandler = &Verifier{
	Next: nil,
}

var msgDistributor = &Distributor{
	Serial: &XMLSerialization{},
}

// POST方法的处理对象
var postHandler = &Verifier{
	Next: msgDistributor,
}

func WechatGet(c *gin.Context) {
	getHandler.HandleRequest(c)
}

func WechatPost(c *gin.Context) {
	postHandler.HandleRequest(c)
}
