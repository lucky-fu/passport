package middler

import (
	"fmt"

	"github.com/passport/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql ...
var Mysql *mysqlStorage

type mysqlStorage struct {
	masterClient *gorm.DB
	slaveClinet  *gorm.DB
}

// InitMysql ...
func InitMysql() {
	Mysql = &mysqlStorage{
		masterClient: getDatabaseClient(Config.Mysql.Master),
		slaveClinet:  getDatabaseClient(Config.Mysql.Slave),
	}
}

func getDatabaseClient(dbConfig define.MysqlConfig) *gorm.DB {
	var (
		err    error
		client *gorm.DB
		dns    string
	)
	// "root:123@tcp(host:port)/test?charset=utf8"
	dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, "Asia%2FShanghai")
	if client, err = gorm.Open(mysql.Open(dns)); nil != err {
		panic("初始化数据库连接失败, 原因 : " + err.Error())
	}
	realDB, err := client.DB()
	if err != nil {
		panic(fmt.Errorf("connect read database %s failed, %v", dbConfig.Host, err.Error()))
	}

	realDB.SetMaxIdleConns(dbConfig.Max / 2)
	realDB.SetMaxOpenConns(dbConfig.Max)
	fmt.Println("初始化mysql成功")
	return client
}

// GetMasterClient ...
func (ms *mysqlStorage) GetMasterClient() *gorm.DB {
	return ms.masterClient
}

// GetSlaveClient ...
func (ms *mysqlStorage) GetSlaveClient() *gorm.DB {
	return ms.slaveClinet
}
