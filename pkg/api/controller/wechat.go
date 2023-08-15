package controller

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/service"
)

type Wechat struct{}

func NewWechat() *Wechat {
	return &Wechat{}
}

func (*Wechat) GetAccessToken(ctx context.Context, code string) (*response.ResponseGetToken, error) {
	return service.GetAccessToken(ctx, code)
}
