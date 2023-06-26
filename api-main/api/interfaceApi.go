package api

import (
	"api-main/models"
	"api-main/utils"
	"context"
	"io/ioutil"
	"strings"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	openapisdkgo "github.com/jianping5/open-api-sdk-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddInterfaceInfo
//
//	@Description: 创建接口，InterfaceInfoAddRequest
//	@param c
func AddInterfaceInfo(c *gin.Context) {
	var interfaceInfo models.InterfaceInfo
	var interfaceInfoAddRequest models.InterfaceInfoAddRequest

	// 获取请求参数
	res := bindContextJson(c, &interfaceInfoAddRequest)
	if !res {
		return
	}
	// 获取当前用户
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)
	userId := user.Id

	// 初始化接口信息
	interfaceInfo = models.InterfaceInfo{
		UserId:         userId,
		Name:           interfaceInfoAddRequest.Name,
		Description:    interfaceInfoAddRequest.Description,
		Url:            interfaceInfoAddRequest.Url,
		RequestParams:  interfaceInfoAddRequest.RequestParams,
		RequestHeader:  interfaceInfoAddRequest.RequestHeader,
		ResponseHeader: interfaceInfoAddRequest.ResponseHeader,
		Status:         1,
		Method:         interfaceInfoAddRequest.Method,
		IsDelete:       0,
	}

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("interface_info")
	result, err := coll.InsertOne(context.TODO(), interfaceInfo)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(result))

}

// DeleteInterfaceInfo
//
//	@Description: 删除接口，DeleteRequest
//	@param c
func DeleteInterfaceInfo(c *gin.Context) {
	var deleteRequest models.DeleteRequest

	// 获取接口 id
	res := bindContextJson(c, &deleteRequest)
	if !res {
		return
	}
	interfaceId, err := primitive.ObjectIDFromHex(deleteRequest.Id)
	if err != nil {
		panic(err)
	}

	coll := DB.Database("open-api").Collection("interface_info")
	filter := bson.D{{"_id", interfaceId}}
	update := bson.D{{"$set", bson.D{{"is_delete", 1}}}}
	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK("删除成功"))
}

// UpdateInterfaceInfo
//
//	@Description: 修改接口，InterfaceInfoUpdateRequest
//	@param c
func UpdateInterfaceInfo(c *gin.Context) {
	var interfaceInfo models.InterfaceInfo
	var interfaceInfoUpdateRequest models.InterfaceInfoUpdateRequest

	// 获取请求参数
	res := bindContextJson(c, &interfaceInfoUpdateRequest)
	if !res {
		return
	}

	// 获取当前用户
	// session := sessions.Default(c)
	// user := session.Get(utils.USER_LOGIN_STATE).(models.User)
	// userId := user.Id

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId, err := primitive.ObjectIDFromHex(interfaceInfoUpdateRequest.Id)

	// 先查询旧接口信息
	filter := bson.D{{"_id", interfaceInfoId}}
	err = coll.FindOne(context.TODO(), filter).Decode(&interfaceInfo)

	// 更新（todo：待优化）
	interfaceInfo.Name = interfaceInfoUpdateRequest.Name
	interfaceInfo.Description = interfaceInfoUpdateRequest.Description
	interfaceInfo.Url = interfaceInfoUpdateRequest.Url
	interfaceInfo.RequestParams = interfaceInfoUpdateRequest.RequestParams
	interfaceInfo.RequestHeader = interfaceInfoUpdateRequest.RequestHeader
	interfaceInfo.ResponseHeader = interfaceInfoUpdateRequest.ResponseHeader
	interfaceInfo.Method = interfaceInfoUpdateRequest.Method
	if interfaceInfoUpdateRequest.Status != 0 {
		interfaceInfo.Status = interfaceInfoUpdateRequest.Status
	}

	// 更新
	filter = bson.D{{"_id", interfaceInfoId}}
	update := bson.D{{"$set", interfaceInfo}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(interfaceInfoUpdateRequest.Id))
}

// GetInterfaceInfoById
//
//	@Description: 根据 id 获取接口信息，id
//	@param c
func GetInterfaceInfoById(c *gin.Context) {
	// var idRequest models.IdRequest
	var interfaceInfo models.InterfaceInfo

	// 获取接口 id
	// res := bindContextJson(c, &idRequest)
	id, ok := c.GetQuery("id")
	if !ok {
		return
	}

	// 查询 MongoDB
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId, error := primitive.ObjectIDFromHex(id)
	if error != nil {
		panic(error)
	}

	filter := bson.D{{"_id", interfaceInfoId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&interfaceInfo)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(interfaceInfo))
}

// OnlineInterfaceInfo
//
//	@Description: 根据 id 上线指定接口 IdRequest
//	@param c
func OnlineInterfaceInfo(c *gin.Context) {
	// 获取接口 id
	var idRequest models.IdRequest
	res := bindContextJson(c, &idRequest)
	if !res {
		return
	}
	id := idRequest.Id

	// 更改对应接口的 status
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId, error := primitive.ObjectIDFromHex(id)
	if error != nil {
		panic(error)
	}

	filter := bson.D{{"_id", interfaceInfoId}}
	update := bson.D{{"$set", bson.D{{"status", 2}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(true, "上线成功"))
}

// OffLineInterfaceInfo
//
//	@Description: 根据 id 下线指定接口 IdRequest
//	@param c
func OffLineInterfaceInfo(c *gin.Context) {
	// 获取接口 id
	var idRequest models.IdRequest
	res := bindContextJson(c, &idRequest)
	if !res {
		return
	}
	id := idRequest.Id

	// 更改对应接口的 status
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", interfaceInfoId}}
	update := bson.D{{"$set", bson.D{{"status", 1}}}}
	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(true, "下线成功"))
}

// InvokeInterfaceInfo
//
//	@Description: 测试调用接口 InterfaceInfoInvokeRequest
//	@param c
func InvokeInterfaceInfo(c *gin.Context) {
	// 获取请求参数
	var interfaceInfo models.InterfaceInfo
	var interfaceInfoInvokeRequest models.InterfaceInfoInvokeRequest
	res := bindContextJson(c, &interfaceInfoInvokeRequest)
	if !res {
		return
	}

	// 查询该接口信息
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId, err := primitive.ObjectIDFromHex(interfaceInfoInvokeRequest.Id)
	filter := bson.D{{"_id", interfaceInfoId}, {"is_delete", 0}}
	err = coll.FindOne(context.TODO(), filter).Decode(&interfaceInfo)
	if err != nil {
		panic(err)
	}

	// 获取当前登录用户
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)

	// 统一走网关调用（SDK）

	// 请求参数
	body := interfaceInfoInvokeRequest.UserRequestParams

	client := &http.Client{}
	var req *http.Request

	// GET 请求
	if interfaceInfo.Method == "GET" {
		req, err = http.NewRequest("GET", interfaceInfo.Url, nil)
		if err != nil {
			return
		}
	}

	// POST 请求
	if interfaceInfo.Method == "POST" {
		req, err = http.NewRequest("POST", interfaceInfo.Url, strings.NewReader(body))
		if err != nil {
			return
		}
	}

	// 添加请求头（sdk）
	openApiClient := openapisdkgo.NewClient(user.AccessKey, user.SecretKey)
	openApiClient.AddHeaders(req, body)

	// 处理结果
	resp, err := client.Do(req)
	if err != nil {
		return
	}



	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(string(bytes)))
}

// ListInterfaceInfoByPage
//
//	@Description: 分页获取接口 InterfaceInfoQueryRequest
//	@param c
func ListInterfaceInfoByPage(c *gin.Context) {
	currentStr, _ := c.GetQuery("current")
	pageSizeStr, _ := c.GetQuery("pageSize")
	name, _ := c.GetQuery("Name")
	description, _ := c.GetQuery("Description")
	method, _ := c.GetQuery("Method")
	url, _ := c.GetQuery("Url")
	requestParams, _ := c.GetQuery("RequestParams")
	requestHeader, _ := c.GetQuery("RequestHeader")
	responseHeader, _ := c.GetQuery("ResponseHeader")
	statusStr, _ := c.GetQuery("Status")

	// 获取当前页和每页大小
	current, _ := strconv.ParseInt(currentStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	// 查询数据库
	coll := DB.Database("open-api").Collection("interface_info")
	var status int
	filter := bson.D{}
	if statusStr != "" {
		status, _ = strconv.Atoi(statusStr)
		filter = bson.D{{"is_delete", 0}, {"status", status}}
	} else {
		filter = bson.D{{"is_delete", 0}}
	}

	// 实现模糊查询

	if name != "" {
		filter = append(filter, bson.E{
			Key:   "name",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + name + ".*", Options: "i"}},
		})
	}

	if description != "" {
		filter = append(filter, bson.E{
			Key:   "description",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + description + ".*", Options: "i"}},
		})
	}

	if method != "" {
		filter = append(filter, bson.E{
			Key:   "method",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + method + ".*", Options: "i"}},
		})
	}

	if url != "" {
		filter = append(filter, bson.E{
			Key:   "url",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + url + ".*", Options: "i"}},
		})
	}

	if requestParams != "" {
		filter = append(filter, bson.E{
			Key:   "request_params",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + requestParams + ".*", Options: "i"}},
		})
	}

	if requestHeader != "" {
		filter = append(filter, bson.E{
			Key:   "request_header",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + requestHeader + ".*", Options: "i"}},
		})
	}

	if responseHeader != "" {
		filter = append(filter, bson.E{
			Key:   "response_header",
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + responseHeader + ".*", Options: "i"}},
		})
	}

	// 分页配置
	opts := options.Find().SetLimit(pageSize).SetSkip((current - 1) * pageSize)
	cursor, err := coll.Find(context.TODO(), filter, opts)

	// 存储接口信息列表
	var results []models.InterfaceInfo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(results))
}
