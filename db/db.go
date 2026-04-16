package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"xorm.io/xorm"
)

// DbOption
type DbOption struct {
	// ======================= common option ======================
	driver string                // 数据库驱动: MYSQL
	orm string                   // orm类型: xorm(XORM) | grom(GORM)

	// gorm: username:password@tcp(ip:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local
	// xorm: username:password@tcp(ip:port)/energy_pledge?charset=utf8mb4&parseTime=True&loc=Local
	dsn string                   // 数据源
	maxIdleConn int              // 最大空闲连接数
	maxOpenConn int              // 最大打开连接数

	// ======================= xorm options =======================
	showSql     bool   
}

var (
	XDB *xorm.Engine
	GDB *gorm.DB
)

type EngineOption func(option *DbOption)

func New(options ...EngineOption) *DbOption {
	op := &DbOption{
		driver: "mysql",
		maxIdleConn: 10,
		maxOpenConn: 20,
	}

	for _, option := range options {
		option(op)
	}

	if op.orm == "" {
		panic(fmt.Errorf("orm is now allow empty"))
	}

	if strings.ToUpper(op.orm) == "XORM" {
		op.connectXorm()
	} else if strings.ToUpper(op.orm) == "GORM" {
		op.connectGorm()
	} else {
		panic(fmt.Errorf("orm:[%s] is now allowd", op.orm))
	}

	return op
}

func WithDbDriver(driver string) EngineOption {
	return func(o *DbOption) {
		o.driver = driver
	}
}

func WithOrm(orm string) EngineOption {
	return func (o *DbOption)  {
		o.orm = orm
	}
}

func WithDsn(dsn string) EngineOption {
	return func (o *DbOption)  {
		o.dsn = dsn
	}
}


func WithMaxIdleConn(maxIdleConn int) EngineOption {
	return func(o *DbOption) {
		o.maxIdleConn = maxIdleConn
	}
}

func WithMaxOpenConn(maxOpenConn int) EngineOption {
	return func(o *DbOption) {
		o.maxOpenConn = maxOpenConn
	}
}

func WithShowSql(showSql bool) EngineOption {
	return func(o *DbOption) {
		o.showSql = showSql
	}
}
