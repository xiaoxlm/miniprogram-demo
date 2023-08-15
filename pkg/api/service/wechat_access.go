package service

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/xiaoxlm/miniprogram-demo/cmd/miniprogram/global"
)

func GetAccessToken(ctx context.Context, code string) (*response.ResponseGetToken, error) {
	mini, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     global.Config.AppID,     // 小程序、公众号或者企业微信的appid
		Secret:    global.Config.AppSecret, // 商户号 appID
		HttpDebug: true,
		Debug:     false,
	})
	if err != nil {
		return nil, err
	}

	// code2session
	if _, err = mini.Auth.Session(ctx, code); err != nil {
		return nil, err
	}

	// 71_UpbrYP-Wkw-ECzI199QzSOiYlzy7FIWmJ9m3_CXKJUoP3J1gbJqDJ5e5JLFnEELAR9_wQugh6mF8zFy3XE5iXaiMdJnpkeCPBSi49MlzdqUOestS9-H6o6m1ZFUGWQeAAAVLF
	return mini.AccessToken.Refresh().GetToken(true)
}
