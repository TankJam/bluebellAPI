package mysql

import (
	"bluebellAPI/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义全局的db对象
var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname?parseTime=true&loc=Local"
	// ?parseTime=true&loc=Local 这段不加的话，取时间就会报错!
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User,
		cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	// 获取连接后的db对象
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	// 设置最大连接数
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// 设置空闲连接的最大数量
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 关闭mysql连接
func Close() {
	_ = db.Close()
}
