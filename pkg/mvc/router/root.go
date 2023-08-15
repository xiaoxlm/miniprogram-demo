package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/pkg/mvc/router/test"
	"github.com/xiaoxlm/miniprogram-demo/pkg/mvc/router/user"
)

const (
	GROUP_PATH = "/miniprogram-demo/api"
)

func NewRooter(r *gin.Engine) {
	v1 := r.Group(GROUP_PATH + "/v1")
	v1.GET("/test", test.Test)

	user.Router(v1)
}
