package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.参数校验
	var p *models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是vaslidator.ve 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	}

	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.signup failed:%v\n", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseErr(c, CodeUserExist)
			return

		}
		ResponseErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	var p *models.ParamLogin
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是vaslidator.ve 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	}
	//2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseErr(c, CodeUSerNotExist)
			return
		}
		ResponseErr(c, CodeServerBusy)
		return
	}

	//3.返回相应
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserID,
		"user_name": user.Username,
		"token":     user.Token,
	})

}
