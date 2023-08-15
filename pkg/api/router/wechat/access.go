package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/controller"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/response"
	"net/http"
)

func Router(r *gin.RouterGroup) {
	r.Group("/wechat").GET("/access-token", GetAccessToken)
}

func GetAccessToken(c *gin.Context) {
	code := c.Query("code")

	accessToken, err := controller.NewWechat().GetAccessToken(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error{Error: err})
		return
	}

	c.JSON(http.StatusOK, accessToken)
	return
}
