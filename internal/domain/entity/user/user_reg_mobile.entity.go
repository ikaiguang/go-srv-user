// Package entities
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package entities

import (
	"time"
)

var _ = time.Time{}

// UserRegMobile ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户-注册-手机号'
type UserRegMobile struct {
	Id          uint64    `gorm:"column:id;primaryKey" json:"id"`          // ID
	UserMobile  string    `gorm:"column:user_mobile" json:"user_mobile"`   // 用户手机号码
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"` // 最后修改时间
	UserUuid    string    `gorm:"column:user_uuid" json:"user_uuid"`       // uuid
}

// TableName table name
func (s *UserRegMobile) TableName() string {
	return "srv_user_reg_mobile"
}
