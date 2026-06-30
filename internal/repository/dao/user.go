package dao

import (
	"context"
	"errors"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = errors.New("邮箱冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	Username  string         `gorm:"size:64;not null;uniqueIndex:idx_user_username;comment:用户名" json:"username"`
	Email     sql.NullString `gorm:"size:128;not null;uniqueIndex:idx_user_email;comment:邮箱" json:"email"`
	Password  string         `gorm:"size:255;not null;comment:密码" json:"-"`
	Ctime     int64          `gorm:"autoCreateTime:nano;not null;comment:创建时间" json:"ctime"`
	Uptime    int64          `gorm:"autoUpdateTime:nano;not null;comment:更新时间" json:"uptime"`
	IsDeleted bool           `gorm:"not null;default:false;index;comment:是否删除" json:"is_deleted"`
}

type UserDao interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	Insert(ctx context.Context, u User) error
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{db: db}
}

func (dao *GormUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?").First(&u).Error
	return u, err
}

func (dao *GormUserDao) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return u, err
}

func (dao *GormUserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Uptime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictErrNo {
			// 邮箱冲突
			return ErrUserDuplicate
		}
	}
	return err
}
