package api

import (
	"api-main/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//
// UserRegister
//  @Description: 用户注册, 需要接受 json 内容:user_account、user_password、check_password
//  @param c
//
func UserRegister(c *gin.Context) {
	var user models.User
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	user = models.User{
		UserName: "Tom",
		UserAccount: "123456789",
	}

	coll := DB.Database("sample_user").Collection("user")
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)

	c.JSON(http.StatusOK, result)
}

/* func UserRegister(c *gin.Context) {
	// var user models.User
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	// user = models.User{
	// 	UserName: "Tom",
	// 	UserAccount: "123456789",
	// }

	coll := DB.Database("sample_user").Collection("user")
	idStr := "6491119fb6fbbea0dc97513c"
	id, err := primitive.ObjectIDFromHex(idStr)

	filter := bson.D{{"_id", id}}
	var result models.User
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	fmt.Println(result)

	c.JSON(http.StatusOK, result)
} */


//
// UserLogin
//  @Description: 用户登录, 需要接受 json 内容:user_account、user_password
//  @param c
//
func UserLogin(c *gin.Context) {
	var user models.User
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	user = models.User{
		UserName: "Tom",
		UserAccount: "123456789",
	}

	coll := DB.Database("sample_user").Collection("user")
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)

	c.JSON(http.StatusOK, result)
}


//
// UserLogout
//  @Description: 用户登出
//  @param c
//
func UserLogout(c *gin.Context) {

}