package itools

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
	"xorm.io/xorm"
)

const DBDriverDefault = "mysql"

type DbOption struct {
	DbDriver    string `json:"db_driver"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	Charset     string `json:"charset"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
	ShowSql     bool   `json:"show_sql"`
}

type DbEngine struct {
	DbOption
}

type EngineOption func(option *DbOption)

func NewXrmEngine(options ...EngineOption) *DbEngine {
	d := &DbOption{}
	if d.DbDriver == "" {
		d.DbDriver = DBDriverDefault
	}
	if d.MaxIdleConn == 0 {
		d.MaxIdleConn = 10
	}
	if d.MaxOpenConn == 0 {
		d.MaxOpenConn = 20
	}
	for _, option := range options {
		option(d)
	}
	return &DbEngine{*d}
}

func WithXrmDbDriver(driver string) EngineOption {
	return func(o *DbOption) {
		o.DbDriver = driver
	}
}

func WithXrmHost(host string) EngineOption {
	return func(o *DbOption) {
		o.Host = host
	}
}

func WithXrmPort(port string) EngineOption {
	return func(o *DbOption) {
		o.Port = port
	}
}

func WithXrmUsername(username string) EngineOption {
	return func(o *DbOption) {
		o.Username = username
	}
}

func WithXrmPassword(password string) EngineOption {
	return func(o *DbOption) {
		o.Password = password
	}
}

func WithXrmDatabase(database string) EngineOption {
	return func(o *DbOption) {
		o.Database = database
	}
}

func WithXrmCharset(charset string) EngineOption {
	return func(o *DbOption) {
		o.Charset = charset
	}
}

func WithXrmMaxIdleConn(maxIdleConn int) EngineOption {
	return func(o *DbOption) {
		o.MaxIdleConn = maxIdleConn
	}
}

func WithXrmMaxOpenConn(maxOpenConn int) EngineOption {
	return func(o *DbOption) {
		o.MaxOpenConn = maxOpenConn
	}
}

func WithXrmShowSql(showSql bool) EngineOption {
	return func(o *DbOption) {
		o.ShowSql = showSql
	}
}

func (d *DbEngine) compact() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		d.DbOption.Username,
		d.DbOption.Password,
		d.DbOption.Host,
		d.DbOption.Port,
		d.DbOption.Database,
		d.DbOption.Charset)
}

func (d *DbEngine) Connect() {
	engine, err := xorm.NewEngine(d.DbOption.DbDriver, d.compact())
	if err != nil {
		os.Exit(-1)
	}
	engine.SetMaxIdleConns(d.DbOption.MaxIdleConn)
	engine.SetMaxOpenConns(d.DbOption.MaxOpenConn)
	engine.ShowSQL(d.DbOption.ShowSql)

	setDb(engine)
}

func (d *DbEngine) Close() {
	d.Close()
}

var Db *xorm.Engine

func setDb(_e *xorm.Engine) {
	Db = _e
	go func() {
		err := Db.Ping()
		if err != nil {
			log.Fatalf("db connect ping err: %v", err)
		}
		time.Sleep(1 * time.Hour)
	}()
}
