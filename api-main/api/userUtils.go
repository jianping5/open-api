package api

import (
	"fmt"
	"net/http"
	"api-main/utils"
	"github.com/gin-gonic/gin"
)

/*
	主要对userApi中常出现的内容进行封装
*/

//
// bindContextJson
//  @Description: 用来绑定 context 的 json 值，并做相关的错误处理
//  @param c
//  @param data 需要把 context 的 json 值绑定到什么变量上
//  @return bool 返回 true 为成功，false 为失败
//
func bindContextJson(c *gin.Context, data interface{}) bool {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("[api bindContextJson err] c.ShouldBindJSON : ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "缺少必填参数或不合法！"))
		return false
	}
	return true
}

