package src

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var Conn *Connection

type Connection struct {
	Db  *gorm.DB
	err error
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

func GetConnection() *Connection {
	if Conn == nil {
		Conn = newConnection()
	}
	return Conn
}

func newConnection() *Connection {
	db, err := gorm.Open("mysql", "citizix_user:An0thrS3crt@/citizix_db?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return &Connection{
		Db:  db,
		err: err,
	}
}

func (this *Connection) PrintFirst(model interface{}) {
	this.Db.First(model)
	fmt.Println(model)
}

func (this *Connection) Create(model interface{}) int64 {
	return this.Db.Create(model).RowsAffected
}

func (this *Connection) CloseDb() {
	if this.Db != nil {
		err := this.Db.Close()
		if err != nil {
			return
		}
	}
}
