package db

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func (d *DbOption) connectXorm() {
	engine, err := xorm.NewEngine(d.driver, d.dsn)
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(d.maxIdleConn)
	engine.SetMaxOpenConns(d.maxOpenConn)
	engine.ShowSQL(d.showSql)

	go func() {
		err := engine.Ping()
		if err != nil {
			log.Fatalf("db connect ping err: %v", err)
		}
		time.Sleep(1 * time.Hour)
	}()

	XDB = engine
}

func Transaction(engine *xorm.Engine, fs ...func(tx *xorm.Session) error) (err error) {
	session := engine.NewSession()
	if err = session.Begin(); err != nil {
		return
	}

	for _, f := range fs {
		if err = f(session); err != nil {
			if err = session.Rollback(); err != nil {
				return
			}
			if err = session.Close(); err != nil {
				return
			}
			return
		}
		if err = session.Commit(); err != nil {
			return
		}
		if err = session.Close(); err != nil {
			return
		}
	}
	return nil
}
