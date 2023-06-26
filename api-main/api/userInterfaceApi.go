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

// AddUserInterfaceInfo
//
//	@Description: 添加用户接口关系，UserInterfaceInfoAddRequest
//	@param c
func AddUserInterfaceInfo(c *gin.Context, interfaceInfoId primitive.ObjectID) bool {
	var userInterfaceInfo models.UserInterfaceInfo
	var oldUserInterfaceInfo models.UserInterfaceInfo
	// 获取当前用户
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)
	userId := user.Id

	// 初始化用户接口调用信息信息
	userInterfaceInfo = models.UserInterfaceInfo{
		UserId:          userId,
		InterfaceInfoId: interfaceInfoId,
		TotalNum:        1,
		LeftNum:         10,
	}

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("user_interface_info")
	
	// 查询是否存在
	filter := bson.D{{"user_id", userId}, {"interface_info_id", interfaceInfoId}}
	err := coll.FindOne(context.TODO(), filter).Decode(&oldUserInterfaceInfo)
	// 不存在，则将初始化的用户接口调用信息添加到数据库
	if oldUserInterfaceInfo == (models.UserInterfaceInfo{}) {
		_, err = coll.InsertOne(context.TODO(), userInterfaceInfo)
		if err != nil {
			panic(err)
		}
	} else {
		// 若存在，则先判断是否还有调用次数
		if oldUserInterfaceInfo.LeftNum <= 0 {
			return false
		}

		// 否则，则更新
		update := bson.D{{"$set", bson.D{
			{"total_num", oldUserInterfaceInfo.TotalNum+1},
			{"left_num", oldUserInterfaceInfo.LeftNum-1},
		}}}

		_, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		
	}

	return true

}

// ListTopInvokeInterfaceInfo
//
//	@Description: 获取调用次数最高的接口
//	@param c
func ListTopInvokeInterfaceInfo(c *gin.Context) {
	var interfaceInfoDTOList []models.InterfaceInfoDTO
	var interfaceInfoList []models.InterfaceInfo
	var userInterfaceInfoList []models.UserInterfaceInfo
	var idList []primitive.ObjectID
	var idTotalNumMap map[primitive.ObjectID]uint

	// 查找 total_num 前 3 的接口 id（user_interface_info）
	coll := DB.Database("open-api").Collection("user_interface_info")
	opts := options.Find().SetSort(bson.D{{"total_num", -1}})
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)

	if err = cursor.All(context.TODO(), &userInterfaceInfoList); err != nil {
		panic(err)
	}

	// 存储接口 id 和 对应调用总数的 map
	idTotalNumMap = make(map[primitive.ObjectID]uint, 0)

	// 获得其 id 组成的数组，并填充 map
	for _, item := range userInterfaceInfoList {
		idList = append(idList, item.InterfaceInfoId)
		idTotalNumMap[item.InterfaceInfoId] +=  item.TotalNum
	}

	// 然后根据 idList 查找对应的接口名称（interface_info）
	coll = DB.Database("open-api").Collection("interface_info")

	filter := bson.D{{"is_delete", 0}}

	filter = append(filter, bson.E{
		Key: "_id",
		Value: bson.M{
			"$in": idList,
		},
	})

	cursor, err = coll.Find(context.TODO(), filter)

	if err = cursor.All(context.TODO(), &interfaceInfoList); err != nil {
		panic(err)
	}

	// 最后遍历 接口数组，将数值填充到 【接口信息 DTO】 内
	// 调用总数就拿当前的接口 Id 从之前存储的 idTotalNumMap 中获取
	// TODO：暂时计算 top 2
	count := 0
	for _, item := range interfaceInfoList {
		if count >= 2 {
			break
		}
		interfaceInfoDTOList = append(interfaceInfoDTOList, models.InterfaceInfoDTO{
			Name: item.Name,
			TotalNum: idTotalNumMap[item.Id],
		})
		count += 1
	}


	// 最后赋值给 interfaceInfoDTOList 并返回
	c.JSON(http.StatusOK, utils.ResponseOK(interfaceInfoDTOList))
}
