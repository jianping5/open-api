package work

import (
	"api-main/api"
	"context"
	"log"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
)

func updateInvokeNum() {
	// 更新每个人的接口调用次数
	coll := api.DB.Database("open-api").Collection("user_interface_info")
	update := bson.D{{"$set", bson.D{
		{"left_num", uint(10)},
	}}}
	_, err := coll.UpdateMany(context.TODO(), bson.D{}, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("更新用户接口调用次数")

}

func TimedWork() {
	c := cron.New()
	// 添加要执行的定时任务，使用 cron 表达式指定执行时间
	c.AddFunc("0 0 0 * * *", updateInvokeNum) // 每天零点执行一次任务
	c.Start()

	// 可以在这里继续执行其他操作

	defer c.Stop()
	select {}
}
