package user

import "time"

type User struct {
	// 用户ID
	ID string `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 邮箱
	Email string `json:"email"`
	// 密码
	Password string `json:"password"`
	// 创建时间
	Ctime time.Time `json:"ctime"`
	// 更新时间
	Uptime time.Time `json:"uptime"`
}
