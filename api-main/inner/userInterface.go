package inner

import (
	"context"
	"api-main/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) AddUserInterfaceInfo(ctx context.Context, req *pb.AddUserInterfaceInfoRequest) (*pb.AddUserInterfaceInfoResponse, error) {
	var userInterfaceInfo models.UserInterfaceInfo
	var oldUserInterfaceInfo models.UserInterfaceInfo

	// 获取当前用户 id 和接口 id
	userId, _ := primitive.ObjectIDFromHex(req.userId)
	interfaceInfoId, _ := primitive.ObjectIDFromHex(req.interfaceInfoId)

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
			return &pb.AddUserInterfaceInfoResponse{
				Data: false,
			}, nil
		}

		// 否则，则更新
		update := bson.D{{"$set", bson.D{
			{"total_num", oldUserInterfaceInfo.TotalNum + 1},
			{"left_num", oldUserInterfaceInfo.LeftNum - 1},
		}}}

		_, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}

	}

	return &pb.addUserInterfaceInfoResponse{
		Data: true,
	}, nil

}