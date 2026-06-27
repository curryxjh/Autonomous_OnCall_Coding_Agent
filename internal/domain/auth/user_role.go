package auth

import "time"

type UserRole struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	IsDeleted bool      `json:"is_deleted"`
	Ctime     time.Time `json:"ctime"`
	Uptime    time.Time `json:"uptime"`
}
