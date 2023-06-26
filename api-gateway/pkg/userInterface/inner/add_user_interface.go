package inner

import (
	// "context"
	"api-gateway/pkg/userInterface/pb"
	"api-gateway/pkg/utils"
	"context"
	"fmt"
	"strconv"
	"time"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	openapisdkgo "github.com/jianping5/open-api-sdk-go"
)

var INTERFACE_URL = "http://localhost:3000"


func AddUserInterfaceInfo(ctx *gin.Context, c pb.UserInterfaceServiceClient) {
	// 1. 获取用户 id
	// 获取 accessKey
	accessKey := ctx.Request.Header.Get("accessKey")

	fmt.Println(accessKey)

	// 发送 rpc 请求，获取对应的 secretKey
	res1, err1 := c.GetUserByAccessKey(context.Background(), &pb.GetUserByAccessKeyRequest{
		AccessKey: accessKey,
	})
	if err1 != nil {
		ctx.AbortWithError(http.StatusBadGateway, err1)
		return 
	}

	// 无权限
	if res1.Error != "" {
		HandleNoAuth(ctx)
		return
	}

	// 提取 secretKey 和 用户 id
	secretKey := res1.Data.SecretKey
	userId := res1.Data.UserId


	// 2. 获取接口 id
	// 获取 url 和 method
	url := INTERFACE_URL + ctx.Request.URL.Path
	method := ctx.Request.Method
	// 发送 rpc 请求，获取对应的接口 id
	res2, err2 := c.GetInterfaceInfo(context.Background(), &pb.GetInterfaceInfoRequest{
		Url: url,
		Method: method,
	})
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadGateway, err2)
		return 
	}
	
	// 提取 interfaceInfoId
	interfaceInfoId := res2.Data
	fmt.Println(interfaceInfoId)

	// 接口不存在
	if res2.Error != "" {
		HandleNoAuth(ctx)
		return
	}
	// userId := "64942fe59bbbfd2c4a848175"
	// interfaceInfoId := "6495626e13abff90306d18a9"


	// 3. 用户鉴权
	nonceStr := ctx.Request.Header.Get("nonce")
	timestampStr := ctx.Request.Header.Get("timestamp")
	signStr := ctx.Request.Header.Get("sign")
	bodyStr := ctx.Request.Header.Get("body")

	fmt.Println(accessKey, nonceStr, timestampStr, signStr, bodyStr)

	// 随机数不满足要求
	nonce, _ := strconv.ParseInt(nonceStr, 10, 64) 
	fmt.Println(nonce)
	if nonce > int64(10000) {
		HandleNoAuth(ctx)
		return
	}

	// 时间和当前时间不能超过 5 分钟
	// 获取当前时间戳 和 请求参数中的时间戳进行比较
	currentTimestamp := time.Now().Unix()
	timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
	fmt.Println(currentTimestamp)
	if currentTimestamp - timestamp >= int64(60*5) {
		HandleNoAuth(ctx)
		return
	}

	fmt.Println(currentTimestamp)

	// 比较签名是否一致
	sign := openapisdkgo.GenSign(bodyStr, secretKey)

	fmt.Println(sign)
	if sign != signStr {
		HandleNoAuth(ctx)
		return
	}

	// 4. 增加接口调用次数
	res3, err3 := c.AddUserInterfaceInfo(context.Background(), &pb.AddUserInterfaceInfoRequest{
		UserId: userId,
		InterfaceInfoId: interfaceInfoId,
	})

	if err3 != nil {
		ctx.AbortWithError(http.StatusBadGateway, err3)
		return 
	}

	// 接口调用次数不足
	if !res3.Data {
		HandleNoCount(ctx)
		return
	}
}

func HandleNoAuth(c *gin.Context) {
	var errorCode utils.ErrorCode
	errorCode.Code = 400101
	errorCode.Message = "无权限"
	c.JSON(http.StatusOK, utils.ResponseError(errorCode, "无权限"), )
	c.Abort()
}

func HandleNoCount(c *gin.Context) {
	var errorCode utils.ErrorCode
	errorCode.Code = 400101
	errorCode.Message = "该接口剩余调用次数不足"
	c.JSON(http.StatusOK, utils.ResponseError(errorCode, "该接口剩余调用次数不足"), )
	c.Abort()
}