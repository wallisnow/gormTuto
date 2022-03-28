package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTopicDetail(context *gin.Context) {
	context.String(200, "v3 print Topic Deail %s", context.Param("topic_id"))
}

func CreateTopic(context *gin.Context) {
	topic := Topic{}
	err := context.BindJSON(&topic)
	if err != nil {
		context.String(400, "cannot create: %s", err.Error())
	}
	context.JSON(200, topic)
}

func CreateTopics(context *gin.Context) {
	topics := Topics{}
	err := context.BindJSON(&topics)
	if err != nil {
		context.String(400, "cannot create: %s", err.Error())
	}
	context.JSON(200, topics)
}

func DeleteTopic(context *gin.Context) {
	context.String(http.StatusOK, "delete topic")
}

func GetTopicList(context *gin.Context) {
	query := TopicQuery{}
	err := context.BindQuery(&query)
	if err != nil {
		context.String(400, "cannot get query %s", err.Error())
	} else {
		context.JSON(200, query)
	}
}
