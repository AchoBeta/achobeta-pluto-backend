package main

import (
	"tgwp/initalize"
	"tgwp/internal/gin/routerh"
	"tgwp/log/zlog"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	routerh.RunServer()
	zlog.Infof("程序运行完成！")

}
