package inner

import (
	"api-main/api"
	"api-main/models"
	"api-main/pb"
	"context"
	"fmt"
	"log"

	// "fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	pb.UnimplementedUserInterfaceServiceServer
}

// 增加接口调用次数
func (s *Server) AddUserInterfaceInfo(ctx context.Context, req *pb.AddUserInterfaceInfoRequest) (*pb.AddUserInterfaceInfoResponse, error) {
	var userInterfaceInfo models.UserInterfaceInfo
	var oldUserInterfaceInfo models.UserInterfaceInfo

	// 获取当前用户 id 和接口 id
	userId, _ := primitive.ObjectIDFromHex(req.UserId)
	interfaceInfoId, _ := primitive.ObjectIDFromHex(req.InterfaceInfoId)

	// 初始化用户接口调用信息信息
	userInterfaceInfo = models.UserInterfaceInfo{
		UserId:          userId,
		InterfaceInfoId: interfaceInfoId,
		TotalNum:        1,
		LeftNum:         10,
	}

	// 添加到数据库中
	coll := api.DB.Database("open-api").Collection("user_interface_info")

	// 查询是否存在
	filter := bson.D{{"user_id", userId}, {"interface_info_id", interfaceInfoId}}
	_ = coll.FindOne(context.TODO(), filter).Decode(&oldUserInterfaceInfo)
	// 不存在，则将初始化的用户接口调用信息添加到数据库
	if oldUserInterfaceInfo == (models.UserInterfaceInfo{}) {
		_, err := coll.InsertOne(context.TODO(), userInterfaceInfo)
		if err != nil {
			log.Fatal(err)
			return &pb.AddUserInterfaceInfoResponse{
				Status: http.StatusOK,
				Error: "插入失败",
			}, nil
		}
	} else {
		// 若存在，则先判断是否还有调用次数
		if oldUserInterfaceInfo.LeftNum <= 0 {
			return &pb.AddUserInterfaceInfoResponse{
				Status: http.StatusOK,
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
			log.Fatal(err)
		}

	}

	return &pb.AddUserInterfaceInfoResponse{
		Status: http.StatusOK,
		Data:   true,
	}, nil

}

// 根据 accessKey 获取用户信息
func (s *Server) GetUserByAccessKey(ctx context.Context, req *pb.GetUserByAccessKeyRequest) (*pb.GetUserByAccessKeyResponse, error) {
	var user models.User
	accessKey := req.AccessKey

	// 查询数据库
	coll := api.DB.Database("open-api").Collection("user")

	filter := bson.D{{"access_key", accessKey}}
	_ = coll.FindOne(context.TODO(), filter).Decode(&user)

	// 封装响应 user
	userRes := &pb.User{
		UserId: user.Id.Hex(),
		SecretKey: user.SecretKey,
	}

	fmt.Println(user)

	// 若为空
	if user == (models.User{}) {
		return &pb.GetUserByAccessKeyResponse{
			Status: http.StatusOK,
			Data:   userRes,
			Error: "无权调用",
		}, nil
	}

	// 返回结果
	return &pb.GetUserByAccessKeyResponse{
		Status: http.StatusOK,
		Data:   userRes,
	}, nil

	// userRes := &pb.User{
	// 	UserId: "123",
	// 	SecretKey: "123",
	// }
	// return &pb.GetUserByAccessKeyResponse{
	// 	Status: http.StatusOK,
	// 	Data:   userRes,
	// }, nil
}

// 根据请求路径和请求方法类型获取接口信息
func (s *Server) GetInterfaceInfo(ctx context.Context, req *pb.GetInterfaceInfoRequest) (*pb.GetInterfaceInfoResponse, error) {
	var interfaceInfo models.InterfaceInfo
	url := req.Url
	method := req.Method

	fmt.Println(url)
	fmt.Println(method)

	// 查询数据库
	coll := api.DB.Database("open-api").Collection("interface_info")

	filter := bson.D{{"url", url}, {"method", method}, {"is_delete", 0}}
	_ = coll.FindOne(context.TODO(), filter).Decode(&interfaceInfo)

	fmt.Println(interfaceInfo)

	if interfaceInfo == (models.InterfaceInfo{}) {
		return &pb.GetInterfaceInfoResponse{
			Status: http.StatusOK,
			Error: "不存在",
		}, nil
	}

	return &pb.GetInterfaceInfoResponse{
		Status: http.StatusOK,
		Data:   interfaceInfo.Id.Hex(),
	}, nil

	// return &pb.GetInterfaceInfoResponse{
	// 	Status: http.StatusOK,
	// 	Data:   "123",
	// }, nil
}
