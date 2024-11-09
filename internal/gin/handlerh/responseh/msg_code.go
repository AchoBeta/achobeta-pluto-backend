package responseh

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type nilStruct struct{}

type Response struct {
	MsgCode MsgCode     `json:"msg_code"`
	Data    interface{} `json:"data"`
}

func response(c *gin.Context, msgcode MsgCode, data any) {
	c.JSON(http.StatusOK, Response{
		MsgCode: msgcode,
		Data:    data,
	})
}
func OkWithData(c *gin.Context, msgcode MsgCode, data any) {
	response(c, msgcode, data)
}

func OkWithMsg(c *gin.Context, msgcode MsgCode) {
	response(c, msgcode, nilStruct{})
}
func FailWithData(c *gin.Context, msgcode MsgCode, data any) {
	response(c, msgcode, data)
}

func FailWithMsg(c *gin.Context, msgcode MsgCode) {
	response(c, msgcode, nilStruct{})
}
