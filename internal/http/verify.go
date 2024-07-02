package http

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// 微信服务器的接入验证器
type Verifier struct {
	// 用于验证业务的Token信息
	Next msgHandler
}

func (v *Verifier) HandleRequest(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")

	// 1.将token、timestamp、nonce三个参数进行字典序排序
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)

	// 2.将三个参数字符串拼接成一个字符串进行sha1加密
	str := strings.Join(strs, "")
	h := sha1.New()
	h.Write([]byte(str))

	// 3.开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
	if fmt.Sprintf("%x", h.Sum(nil)) == signature {
		if v.Next != nil {
			v.Next.HandleRequest(c)
		} else {
			// 流程结束
			return
		}
	} else {
		// 签名验证错误，接入失败
		return
	}
}
