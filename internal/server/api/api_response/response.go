package api_response

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) ResponseOK(ec errcode.ErrCode, data interface{}) {
	g.C.JSON(ec.HTTPCode, Response{
		Success: true,
		Code:    ec.Code,
		Msg:     ec.Desc,
		Data:    data,
	})
}

func (g *Gin) ResponseErr(ec errcode.ErrCode) {
	g.C.JSON(ec.HTTPCode, Response{
		Success: false,
		Code:    ec.Code,
		Msg:     ec.Desc,
		Data:    nil,
	})
}
