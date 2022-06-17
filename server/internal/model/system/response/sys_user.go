package response

import (
	"github.com/tmnhs/fginx/server/internal/model/system"
)

type UserResponse struct {
	User *system.User `json:"user"`
}

type LoginResponse struct {
	User      system.User `json:"user"`
	TokenID   string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}

type UserStatis struct {
	Status           int   `json:"status"`
	LastTimestamp    int64 `json:"lastTimeStamp"`
	FinishOrderCount int   `json:"finishOrderCount"`
}
