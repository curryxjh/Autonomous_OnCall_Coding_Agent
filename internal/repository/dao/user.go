package dao

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	Username  string `gorm:"size:64;not null;uniqueIndex:idx_user_username;comment:用户名" json:"username"`
	Email     string `gorm:"size:128;not null;uniqueIndex:idx_user_email;comment:邮箱" json:"email"`
	Password  string `gorm:"size:255;not null;comment:密码" json:"-"`
	Ctime     int64  `gorm:"autoCreateTime:nano;not null;comment:创建时间" json:"ctime"`
	Uptime    int64  `gorm:"autoUpdateTime:nano;not null;comment:更新时间" json:"uptime"`
	IsDeleted bool   `gorm:"not null;default:false;index;comment:是否删除" json:"is_deleted"`
}

type UserDao interface {
}
