package user

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxlm/miniprogram-demo/cmd/miniprogram/global"
	"net/http"
)

func Router(r *gin.RouterGroup) {
	r.GET("/code2Session", Code2Session)
}

func Code2Session(c *gin.Context) {
	mini, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     global.Config.AppID,     // 小程序、公众号或者企业微信的appid
		Secret:    global.Config.AppSecret, // 商户号 appID
		HttpDebug: true,
		Debug:     false,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	code := c.Query("code")

	rs, err := mini.Auth.Session(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rs)
	return
}
