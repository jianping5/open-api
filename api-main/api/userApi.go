package api

import (
	"api-main/models"
	"api-main/utils"
	"context"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/gin-contrib/sessions/cookie"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

var salt = "open-api"

// UserRegister
//
//	@Description: 用户注册, 需要接受 json 内容:userAccount、userPassword、checkPassword
//	@param c
func UserRegister(c *gin.Context) {
	var user models.User
	var userRegister models.UserRegister

	// todo: 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	res := bindContextJson(c, &userRegister)
	if !res {
		return
	}

	// 生成密钥对
	hash := md5.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(userRegister.UserAccount))
	rand.Seed(time.Now().UnixNano())
	accessKey := hex.EncodeToString(hash.Sum([]byte(strconv.Itoa(rand.Intn(9999)))))
	secretKey := hex.EncodeToString(hash.Sum([]byte(strconv.Itoa(rand.Intn(999999)))))

	// 初始化用户
	user = models.User{
		UserAccount:  userRegister.UserAccount,
		UserRole:     "admin",
		UserPassword: userRegister.UserPassword,
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		IsDelete:     0,
	}

	// 添加到数据库中
	coll := DB.Database("open-api").Collection("user")
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, utils.ResponseOK(result))
}

// UserLogin
//
//	@Description: 用户登录, 需要接受 json 内容:userAccount、userPassword
//	@param c
func UserLogin(c *gin.Context) {
	var user models.User
	var userLogin models.UserLogin

	// 获取用户登录信息
	res := bindContextJson(c, &userLogin)
	if !res {
		return
	}

	// 根据账号查找用户
	coll := DB.Database("open-api").Collection("user")
	filter := bson.D{{"user_account", userLogin.UserAccount}}
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	// 用户不存在
	if user == (models.User{}) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不存在"))
	}
	// 校验账号密码是否匹配
	if user.UserPassword != userLogin.UserPassword {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "密码错误"))
	}

	// todo: 保存登录态
	session := sessions.Default(c)
	session.Set(utils.USER_LOGIN_STATE, user)
	session.Save()

	c.JSON(http.StatusOK, utils.ResponseOK(user))
}

// UserLogout
//
//	@Description: 用户登出
//	@param c
func UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)

	c.JSON(http.StatusOK, utils.ResponseOK(user.Id))
}

// GetCurrentUser
//
//	@Description: 获取当前登录用户
//	@param c
func GetCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(utils.USER_LOGIN_STATE).(models.User)

	c.JSON(http.StatusOK, utils.ResponseOK(user))
}
