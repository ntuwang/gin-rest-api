package v1

import (
	"fmt"
	"gin-rest-api/controller"
	"gin-rest-api/service"
	"gin-rest-api/util"
	"github.com/gin-gonic/gin"
)

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := controller.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	authService := service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Failed(controller.UnAuthorized, err.Error())
		return
	}

	if !isExist {
		appG.Failed(controller.UnAuthorized, "未找到用户")
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Failed(controller.UnAuthorized, fmt.Sprintf("Token生成失败, %s", err.Error()))
		return
	}

	appG.Success(map[string]string{
		"token": token,
	})
}
