package app

import "github.com/gin-gonic/gin"

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, &response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
