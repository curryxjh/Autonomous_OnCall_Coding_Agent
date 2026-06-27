package auth

import "time"

type Role struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	Ctime     time.Time `json:"ctime"`
	Uptime    time.Time `json:"uptime"`
}
