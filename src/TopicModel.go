package src

type Topic struct {
	TopicId         int    `json:"id"`
	TopicTitle      string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"nefield=TopicTitle"`
	TopicIp         string `json:"ip" binding:"ipv4"`
}

type TopicClass struct {
	Id        int `gorm:"PRIMARY_KEY"`
	ClassName string
	ClassType string `gorm:"Column:class_type"`
}

type Topics struct {
	List []Topic `json:"topics" binding:"dive"`
}

type TopicQuery struct {
	Username string `form:"username" json:"username" binding:"required"`
	Page     int    `form:"page" json:"page"`
	Pagesize int    `form:"pagesize" json:"pagesize"`
}

func AddTopic(id int, title string) Topic {
	return Topic{TopicId: id, TopicTitle: title}
}
