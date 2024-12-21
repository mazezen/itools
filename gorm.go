package itools

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"reflect"
	"time"
)

// NewGorm gorm 初始化mysql连接
// eg: root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local
func NewGorm(dsn string, args ...[]int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	if len(args) > 0 {
		if len(args[0]) >= 3 {
			return nil, errors.New("参数错误")
		}
		if len(args[0]) == 1 {
			// 设置数据库连接池最大连接数
			sqlDB.SetMaxOpenConns(args[0][0])
		}
		if len(args[0]) == 2 {
			// 连接池最大允许的空闲连接数
			sqlDB.SetMaxIdleConns(args[0][1])
		}
	} else {
		// 设置数据库连接池最大连接数 100
		sqlDB.SetMaxOpenConns(100)
		// 连接池最大允许的空闲连接数 20
		sqlDB.SetMaxIdleConns(20)
	}

	go func() {
		for {
			_ = sqlDB.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()

	return db, nil
}

// Paginate page, pageSize 传入参数
// s 分页默认配置
func Paginate(page interface{}, pageSize interface{}, s map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p := cast.ToInt(page)
		size := cast.ToInt(pageSize)
		if page == 0 {
			if s["page"] != nil {
				page = s["page"].(int)
			} else {
				page = 1
			}
		}

		if pageSize == 0 {
			if s["pageSize"] != nil {
				pageSize = s["pageSize"].(int)
			} else {
				pageSize = 10
			}
		}
		offset := (p - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

// FilterString 进行快速条件过滤
func FilterString(key, value, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		returnDB := db
		if value != "" {
			switch operator {
			case "like":
				returnDB = returnDB.Where(key+" Like ? ", "%"+value+"%")
			case "=", ">=", "<=", "<":
				returnDB = returnDB.Where(key+" "+operator+" "+"?", value)
			}
		}
		return returnDB
	}
}

// InOrNotInFilter where in 或者 where not in
func InOrNotInFilter(key string, value interface{}, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		returnDB := db
		if reflect.ValueOf(value).Len() != 0 {
			whereMap := map[string]interface{}{
				key: value,
			}
			switch operator {
			case "in":
				returnDB = returnDB.Where(whereMap)
			case "not in":
				returnDB = returnDB.Not(whereMap)
			}
		}
		return returnDB
	}
}
