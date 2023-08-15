package main

import (
	"context"
	"github.com/kunlun-qilian/confserver"
	"github.com/spf13/cobra"
	"github.com/xiaoxlm/miniprogram-demo/cmd/miniprogram/global"
	"github.com/xiaoxlm/miniprogram-demo/pkg/api/router"
)

func main() {
	confserver.Execute(func(cmd *cobra.Command, args []string) {
		router.NewRooter(global.Config.Server.Engine())
		if err := global.Config.Server.Serve(context.Background()); err != nil {
			panic(err)
		}
	})
}
