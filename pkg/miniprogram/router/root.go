package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/pkg/miniprogram/router/test"
)

func NewRooter(r *gin.Engine) {

	v1 := r.Group("/miniprogram/api/v1")

	v1.GET("/test", test.Test)

}
