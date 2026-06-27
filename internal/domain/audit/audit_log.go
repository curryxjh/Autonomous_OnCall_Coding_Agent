package audit

import "time"

type AuditLog struct {
	ID           string
	RequestID    string
	UserID       string
	Method       string
	Path         string
	StatusCode   int
	LatencyMS    int64
	RequestBody  string
	ResponseBody string
	ErrorMsg     string
	Ctime        time.Time
}
