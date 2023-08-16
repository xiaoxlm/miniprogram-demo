package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/controller"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/response"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/router/middleware"
	"net/http"
)

func Router(r *gin.RouterGroup) {
	r.Group("/wechat").GET("/access-token", GetAccessToken)
}

func GetAccessToken(c *gin.Context) {
	code := c.Query("code")

	accessToken, err := controller.NewWechat().GetAccessToken(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	jwt, err := middleware.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error{Error: "generate jwt error." + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Login{
		AccessToken:     jwt,
		MiniAccessToken: accessToken.AccessToken,
		MiniExpiresIn:   accessToken.ExpiresIn,
	})
	return
}
