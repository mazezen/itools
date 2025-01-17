package itools

import (
	"fmt"
	"github.com/mazezen/itools/example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"testing"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/jiaxiao?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Printf("Connecting database failed: %v\n ", err.Error())
		os.Exit(1)
	}
}

// 测试分页
func TestPaginate(t *testing.T) {
	// 前端传递的参数
	page := 1
	pageSize := 10

	// 项目配置的默认分页
	s := make(map[string]interface{})
	s["page"] = 1
	s["pageSize"] = 10

	admins := make([]*models.Admin, 0)
	Db.Scopes(Paginate(page, pageSize, s)).Find(&admins)
	for _, v := range admins {
		fmt.Println(v)
	}
}

// 测试分页和带条件查询 "=", ">=", "<=", "<" , "like"
func TestFilterString(t *testing.T) {
	// 前端传递的参数
	page := 1
	pageSize := 10
	username := "test-demo-01"
	email := "111@aa.com"

	// 项目配置的默认分页
	s := make(map[string]interface{})
	s["page"] = 1
	s["pageSize"] = 10

	admins := make([]*models.Admin, 0)

	// SELECT username = `test-demo-01` AND email = `111@aa.com` LIMIT = 10
	// username 、 email为可选参数，非必填
	Db.Scopes(Paginate(page, pageSize, s)).
		Scopes(FilterString("username", username, "=")).
		Scopes(FilterString("email", email, "=")).Find(&admins)
	for _, v := range admins {
		fmt.Println(v)
	}
}

func TestInOrNotInFilter(t *testing.T) {
	// 前端传递的参数
	page := 1
	pageSize := 10

	// 项目配置的默认分页
	s := make(map[string]interface{})
	s["page"] = 1
	s["pageSize"] = 10

	admins := make([]*models.Admin, 0)
	usernames := []string{"test-demo-01", "test-demo-02"}

	// SELECT * FROM `admin` WHERE `username` IN ('test-demo-01','test-demo-02') LIMIT 10
	// username 非必填
	Db.Scopes(Paginate(page, pageSize, s)).
		Scopes(InOrNotInFilter("username", usernames, "in")).
		Find(&admins)
	for _, v := range admins {
		fmt.Println(v)
	}

	fmt.Println("======================================================")

	// SELECT * FROM `admin` WHERE `username` NOT IN ('test-demo-01','test-demo-02') LIMIT 10
	Db.Scopes(Paginate(page, pageSize, s)).
		Scopes(InOrNotInFilter("username", usernames, "not in")).
		Find(&admins)
	for _, v := range admins {
		fmt.Println(v)
	}
}
