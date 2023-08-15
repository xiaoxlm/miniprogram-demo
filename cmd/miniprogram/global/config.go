package global

import (
	"github.com/kunlun-qilian/conflogger"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.SetServiceName("app1", "..")
	confserver.ConfP(&Config)
	//confserver.AddCommand(Config.MasterDB.Commands()...)
}

var Config = struct {
	Logger  *conflogger.Log
	Server  *confserver.Server
	TestEnv string `env:""`
}{
	Server: &confserver.Server{
		Port:   80,
		Mode:   "debug",
		UseH2C: true,
	},
	TestEnv: "123",
}
