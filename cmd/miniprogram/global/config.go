package global

import (
	"github.com/kunlun-qilian/conflogger"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.SetServiceName("miniprogram-demo", "..")
	confserver.ConfP(&Config)
	//confserver.AddCommand(Config.MasterDB.Commands()...)
}

var Config = struct {
	Logger       *conflogger.Log
	Server       *confserver.Server
	AppID        string `env:""` // 小程序 app id
	AppSecret    string `env:""` // 小程序 app secret
	ClientID     string `env:""`
	ClientSecret string `env:""`
}{
	Server: &confserver.Server{
		Port:   80,
		Mode:   "release",
		UseH2C: true,
	},
	AppID:     "",
	AppSecret: "",

	ClientID:     "ops_user",
	ClientSecret: "vK84Cjem1K3WkCD3kKAJe9s3qSqF51aq",
}
