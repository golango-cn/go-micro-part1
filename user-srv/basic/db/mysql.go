package db

import (
	"database/sql"
	"go-micro-part1/user-srv/basic/config"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/micro/go-micro/v2/logger"
)

func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
