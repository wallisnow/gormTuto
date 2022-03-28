package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	. "topic.gin.com/src"
)

const (
	USERNAME = "username"
)

func main() {

	router := gin.Default()

	//msg := make(map[string]interface{})
	//msg["topics"] = Topic{1, "flower"}
	//===================================================Lesson 1======================================
	router.GET("/v1/topic/:topic_id", func(context *gin.Context) {
		//context.Writer.WriteString("hello")
		context.JSON(http.StatusOK, gin.H{"test": Topic{TopicId: 1, TopicTitle: "test"}})

		context.String(http.StatusOK, "Request ID is %s", context.Param("topic_id"))
	})

	router.GET("/v1/topics", func(context *gin.Context) {
		if context.Query(USERNAME) == "" {
			context.String(200, "get topics")
		} else {
			context.String(200, "%s get topics", context.Query(USERNAME))
		}
	})
	//===================================================Lesson 2======================================
	//当路由过多时， 分组便是一个很好的方式取代直接在url中给定路径， 例如v1, 那么如果我们有v2呢?
	//类似spring controller 中的根url
	v2 := router.Group("/v2/topics")
	{
		v2.GET("/:topic_id", func(context *gin.Context) {
			context.String(200, "call v2 topics : %s ", context.Param("topic_id"))
		})

		v2.GET("", func(context *gin.Context) {
			if context.Query(USERNAME) == "" {
				context.String(200, "no username")
			} else {
				context.String(200, "%s get topics", context.Query(USERNAME))
			}
		})
	}
	//===================================================Lesson 3======================================
	//显然， 业务代码不可以写在一起， 现在开始区分代码
	//1. 创建src文件夹， 并创建TopicDao, 将代码慢慢移动进去
	v3 := router.Group("/v3/topics")
	//2. 简单封装了  GetTopicDetail 成一个 HandlerFunc, 这里保证和HandlerFunc一至， 则不需要 GetTopicDetail() , 注意这里并没有带括号 也就是一个函数
	//可以参照HandlerFunc源码 ， 它就是一个参数为context 的函数类型
	//HandlerFunc defines the handler used by gin middleware as return value.
	//type HandlerFunc func(*Context)
	v3.GET("/:topic_id", GetTopicDetail)

	//为了不写重复代码， 使用中间件的方式来判断token， 以限制创建和删除用户
	v3.Use(HasToken())
	{
		v3.POST("", CreateTopic)
		v3.DELETE("/:topic_id", DeleteTopic)
	}
	//===================================================Lesson 4======================================
	//参数绑定
	//查看GetTopicList函数， 可以看到TopicQuery里面定义了form, 和json， 就达到了绑定的效果， 类似spring中直接传入一个json， 这个json可以直接映射成实体对象
	//验证器， binding
	//这里有很多验证， 例如required， min, max, ip 等, binding:"dive" 表示(topics是list, 那么它并不会验证topic, 加入dive则是必须的)次级也进行验证， 验证字段合法性
	v4 := router.Group("/v4/topics")
	{
		v4.GET("", GetTopicList)
		v4.POST("", CreateTopics)
	}
	//===================================================Lesson 5======================================
	//使用orm, 原生sql
	err := Show()
	if err != nil {
		fmt.Printf("error when connect db %v\n", err.Error())
		return
	}

	//===================================================Lesson 6=======================================
	//使用orm, 非原生sql
	con := NewConnection()
	con.Query()
	con.CloseDb()

	router.Run()
}
