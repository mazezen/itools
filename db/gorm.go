package db

import (
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func (d *DbOption) connectGorm() {
	db, err := gorm.Open(mysql.Open(d.dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(d.maxIdleConn)
	sqlDB.SetMaxOpenConns(d.maxOpenConn)

	go func() {
		for {
			_ = sqlDB.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()

	GDB = db
}

// Paginate
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 || pageSize <= 0 {
			page = 1
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
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
