package mysql

import (
	"bluebellAPI/settings"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// 定义全局的db对象
var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s", cfg.User,
		cfg.Password, cfg.Host, cfg.Password, cfg.DB,
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
