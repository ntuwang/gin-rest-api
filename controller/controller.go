package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success      = 200
	Failed       = 500
	ParamError   = 400
	NotFound     = 404
	UnAuthorized = 401
)

var codeMsg = map[int]string{
	Success:      "正常",
	Failed:       "系统异常",
	ParamError:   "参数错误",
	NotFound:     "记录不存在",
	UnAuthorized: "鉴权失败",
}

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceId interface{} `json:"trace_id"`
}

// BindAndValid binds and validates data
func (g *Gin) BindAndValid(form interface{}) int {
	err := g.C.Bind(form)
	if err != nil {
		return ParamError
	}

	return Success
}

// Success setting gin.JSON
func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    Success,
		Msg:     codeMsg[Success],
		Data:    data,
		TraceId: g.C.GetString("trace_id"),
	})
}

// Failed setting gin.JSON
func (g *Gin) Failed(code int, msg string) {
	errMsg := codeMsg[code] + ": " + msg
	g.C.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    code,
		Msg:     errMsg,
		Data:    nil,
		TraceId: g.C.GetString("trace_id"),
	})
}

// Index
// @Summary root
// @Produce  json
// @Success 200 {object} Response
// @Router / [get]
func Index(c *gin.Context) {
	appG := Gin{C: c}
	appG.Success("Gin-rest-api Web")
}
