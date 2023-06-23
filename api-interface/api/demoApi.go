package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IdRequest struct {
	Id uint `json:"id"`
}

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world")
}

func GetNameById(c *gin.Context) {
	var idRequest IdRequest

	res := bindContextJson(c, &idRequest)
	if !res {
		return
	}

	c.JSON(http.StatusOK, "Hello, " + strconv.Itoa(int(idRequest.Id)))
}

// bindContextJson
//
//	@Description: 用来绑定 context 的 json 值，并做相关的错误处理
//	@param c
//	@param data 需要把 context 的 json 值绑定到什么变量上
//	@return bool 返回 true 为成功，false 为失败
func bindContextJson(c *gin.Context, data interface{}) bool {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("[api bindContextJson err] c.ShouldBindJSON : ", err.Error())
		c.JSON(http.StatusForbidden, "缺少必填参数或不合法！")
		return false
	}
	return true
}
