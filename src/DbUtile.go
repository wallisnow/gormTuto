package src

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Connection struct {
	db *gorm.DB
}

func Show() error {
	db, _ := gorm.Open("mysql", "citizix_user:An0thrS3crt@/citizix_db?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	rows, err := db.Raw("select topic_id,topic_title from topic").Rows()
	if err != nil {
		return err
	}
	for rows.Next() {
		var topic_id int
		var topic_title string
		rows.Scan(&topic_id, &topic_title)
		fmt.Println(topic_id, topic_title)
	}
	return nil
}

func NewConnection() *Connection {
	db, _ := gorm.Open("mysql", "citizix_user:An0thrS3crt@/citizix_db?charset=utf8&parseTime=True&loc=Local")
	return &Connection{
		db: db,
	}
}

func (this *Connection) Query() {
	this.db.LogMode(true)
	tc := &TopicClass{}
	this.db.First(tc)
	fmt.Println(tc)
}

func (this *Connection) CloseDb() {
	if this.db != nil {
		err := this.db.Close()
		if err != nil {
			return
		}
	}
}
