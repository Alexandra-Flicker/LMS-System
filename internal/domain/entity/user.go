package entity

import "lms_system/internal/domain/common"

type UserContext struct {
	UserID uint            `json:"user_id"`
	Role   common.UserRole `json:"role"`
}
