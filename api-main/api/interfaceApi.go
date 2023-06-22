package api

import (
	"api-main/models"
	"api-main/utils"
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	var InterfaceInfoAddRequest models.InterfaceInfoAddRequest

	// 获取请求参数
	res := bindContextJson(c, &InterfaceInfoAddRequest)
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
		Name:           InterfaceInfoAddRequest.Name,
		Description:    InterfaceInfoAddRequest.Description,
		Url:            InterfaceInfoAddRequest.Url,
		RequestParams:  InterfaceInfoAddRequest.RequestParams,
		RequestHeader:  InterfaceInfoAddRequest.RequestHeader,
		ResponseHeader: InterfaceInfoAddRequest.ResponseHeader,
		Status:         0,
		Method:         InterfaceInfoAddRequest.Method,
		IsDelete:       0,
	}

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("interface_info")
	result, err := coll.InsertOne(context.TODO(), interfaceInfo)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOk(interfaceInfo.Id))

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
	interfaceId := primitive.ObjectIDFromHex(deleteRequest.Id)

	coll := DB.Database("open-api").Collection("interface_info")
	filter := bson.D{{"_id", interfaceId}}
	update := bson.D{{"$set", bson.D{{"is_delete", 1}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
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
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)
	userId := user.Id

	// 初始化接口信息
	interfaceInfo = models.InterfaceInfo{
		UserId:         userId,
		Name:           interfaceInfoUpdateRequest.Name,
		Description:    interfaceInfoUpdateRequest.Description,
		Url:            interfaceInfoUpdateRequest.Url,
		RequestParams:  interfaceInfoUpdateRequest.RequestParams,
		RequestHeader:  interfaceInfoUpdateRequest.RequestHeader,
		ResponseHeader: interfaceInfoUpdateRequest.ResponseHeader,
		Status:         interfaceInfoUpdateRequest.Status,
		Method:         interfaceInfoUpdateRequest.Method,
		IsDelete:       0,
	}

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId := primitive.ObjectIDFromHex(interfaceInfoUpdateRequest.Id)
	filter := bson.D{{"_id", interfaceInfoId}}
	update := bson.D{bson.Marshal(interfaceInfo)}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(interfaceInfo.Id))
}

// GetInterfaceInfoById
//
//	@Description: 根据 id 获取接口信息，id
//	@param c
func GetInterfaceInfoById(c *gin.Context) {
	var interfaceInfo models.InterfaceInfo

	// 获取接口 id
	id := c.Param("id")

	// 查询 MongoDB
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId := primitive.ObjectIDFromHex(id)
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
	res := bindContextJson(&idRequest)
	if !res {
		return
	}
	id := idRequest.Id

	// 更改对应接口的 status
	coll := DB.Database("open-api").Collection("interface_info")
	interfaceInfoId := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", interfaceInfoId}}
	update := bson.D{{"status", 1}}
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
	interfaceInfoId := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", interfaceInfoId}}
	update := bson.D{{"status", 0}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
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
	var interfaceInfoInvokeRequest models.InterfaceInfoInvokeRequest
	res := bindContextJson(c, &InterfaceInfoInvokeRequest)
	if !res {
		return
	}

	// todo：调用接口，需要开发 SDK
}

// ListInterfaceInfoByPage
//
//	@Description: 分页获取接口 InterfaceInfoQueryRequest
//	@param c
func ListInterfaceInfoByPage(c *gin.Context) {
	// 获取请求参数
	var interfaceInfoQueryRequest models.InterfaceInfoQueryRequest
	res := bindContextJson(c, &interfaceInfoQueryRequest)
	if !res {
		return
	}

	// 获取当前页和每页大小
	current := interfaceInfoQueryRequest.Current
	pageSize := interfaceInfoQueryRequest.PageSize

	// 查询数据库
	coll := DB.Database("open-api").Collection("interface_info")
	opts := options.Find().SetLimit(pageSize).SetSkip((current-1)*pageSize)
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)

	// 存储接口信息列表
	var results []models.InterfaceInfo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOk(results))
}