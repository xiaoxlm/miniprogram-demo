package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/router/test"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/router/wechat"
)

const (
	GroupPath = "/miniprogram/api"
)

func NewRooter(r *gin.Engine) {
	v1 := r.Group(GroupPath + "/v1")
	v1.GET("/test", test.Test)

	wechat.Router(v1)
}
