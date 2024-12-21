package itools

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// NewXrm  初始化mysql连接
func NewXrm(dsn string, showSql bool, args ...[]int) (*xorm.Engine, error) {
	x, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}
	x.ShowSQL(showSql)
	err = x.Ping()
	if err != nil {
		panic(err)
	}
	x.ShowSQL(showSql)

	if len(args) > 0 {
		if len(args[0]) >= 3 {
			return nil, errors.New("参数错误")
		}
		if len(args[0]) == 1 {
			// 设置数据库连接池最大连接数
			x.SetMaxOpenConns(args[0][0])
		}
		if len(args[0]) == 2 {
			// 连接池最大允许的空闲连接数
			x.SetMaxIdleConns(args[0][1])
		}
	} else {
		// 设置数据库连接池最大连接数 100
		x.SetMaxOpenConns(100)
		// 连接池最大允许的空闲连接数 20
		x.SetMaxIdleConns(20)
	}
	go func() {
		for {
			_ = x.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()

	return x, nil
}
