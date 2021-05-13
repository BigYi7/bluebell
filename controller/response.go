package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code":10001,//程序中错误码
	"msg":xx,//提示信息
	"data":{},//数据
}
*/

type Responsedata struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseErr(c *gin.Context, code ResCode) {
	rd := &Responsedata{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &Responsedata{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Responsedata{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
