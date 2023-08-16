package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/confserver"
	"github.com/xiaoxlm/miniprogram-demo/cmd/miniprogram/global"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/response"
	"net/http"
	"strings"
)

type AuthorizationHeader struct {
	Authorization string `name:"Authorization,omitempty" in:"header" `
}

// Authorization
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := getTokenFromHeader(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error{
				Error: err.Error(),
			})
			return
		}

		if err = CheckToken(accessToken); err != nil {
			ctx.JSON(http.StatusUnauthorized, response.Error{
				Error: err.Error(),
			})
			return
		}

		ctx.Next()
	}
}

func CheckToken(token string) error {
	claims, err := ParseJWT(token)
	if err != nil {
		return err
	}

	if claims.ClientID != global.Config.ClientID || claims.ClientSecret != global.Config.ClientSecret {
		return fmt.Errorf("非法用户")
	}

	return nil
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	header := AuthorizationHeader{}

	if err := confserver.Bind(ctx, &header); err != nil {
		return "", fmt.Errorf("authorization bind invalid")
	}

	seg := strings.Split(header.Authorization, "Bearer ")
	if len(seg) < 2 {
		return "", fmt.Errorf("bearer token parse invalid")
	}

	return seg[1], nil
}
