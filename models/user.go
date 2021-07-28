package models


// User 结构体
type User struct {
	UserID   int64 `db:"user_id"`  // bigint
	Username string `db:"username"`
	Password string `db:"password"`
}
