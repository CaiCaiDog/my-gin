package service

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"my-gin/config"
	"fmt"
)

type MysqlConnect struct{}

var instance *MysqlConnect
var once sync.Once
var db *gorm.DB
var err error

/**
* 创建单例
**/
func GetIns() *MysqlConnect {
	once.Do(func() {
		instance = &MysqlConnect{}
	})
	return instance
}

/**
* 初始化数据库连接
**/
func (m *MysqlConnect) InitConnect() (issucc bool) {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", config.Database("username"), config.Database("password"), config.Database("host"), config.Database("port"), config.Database("database"), config.Database("charset")))
	if err != nil {
		return false
	}
	db.SingularTable(true)
	return true
}

/**
* 对外数据库连接对象
**/
func (m *MysqlConnect) GetMysqlDB() *gorm.DB {
	return db
}
