package routerh

import "github.com/gin-gonic/gin"

func RunServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Run()
}
