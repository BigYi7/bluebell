package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数校验

	//c.ShouldBindJSON()
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseErr(c, CodeInvalidParam)
		return
	}
	//从c中取到当前大请求的用户id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseErr(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.createpost failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)
}

func GetPostDeatailHandler(c *gin.Context) {
	//1.获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseErr(c, CodeInvalidParam)
		return
	}
	//2.根据id取出帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.getpostbyid(pid) failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {

	//获取分页参数
	offset, limit := GetPageInfo(c)

	//获取数据
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, data)
}
