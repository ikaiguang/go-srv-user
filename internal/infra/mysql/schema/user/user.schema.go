// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	gorm "gorm.io/gorm"
	"time"
)

var _ = time.Time{}

// UserSchema User
var UserSchema User

// NewUser new schema
func NewUser() *User {
	return &User{}
}

// User ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户'
type User struct {
	Id              uint64     `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	UserUuid        string     `gorm:"column:user_uuid;unique;type:string;size:255;not null;default:'';comment:uuid" json:"user_uuid"`
	CreatedTime     time.Time  `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime     time.Time  `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	DeletedTime     *time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
	IsDeleted       bool       `gorm:"column:is_deleted;index;type:bool;not null;default:0;comment:是否已删除" json:"is_deleted"`
	UserEmail       string     `gorm:"column:user_email;type:string;size:255;not null;default:'';comment:邮箱" json:"user_email"`
	UserNickname    string     `gorm:"column:user_nickname;type:string;size:255;not null;default:'';comment:昵称" json:"user_nickname"`
	UserAvatar      string     `gorm:"column:user_avatar;type:string;size:255;not null;default:'';comment:头像" json:"user_avatar"`
	UserGender      string     `gorm:"column:user_gender;type:string;size:255;not null;default:'';comment:性别" json:"user_gender"`
	UserStatus      string     `gorm:"column:user_status;type:string;size:255;not null;default:'';comment:状态" json:"user_status"`
	ActiveBeginTime *time.Time `gorm:"column:active_begin_time;type:time;comment:激活开始时间" json:"active_begin_time"`
	ActiveEndTime   *time.Time `gorm:"column:active_end_time;type:time;comment:激活结束时间" json:"active_end_time"`
	DisableTime     *time.Time `gorm:"column:disable_time;type:time;comment:禁用时间" json:"disable_time"`
	BlacklistTime   *time.Time `gorm:"column:blacklist_time;type:time;comment:黑名单时间" json:"blacklist_time"`
	PasswordHash    string     `gorm:"column:password_hash;type:string;size:255;not null;default:'';comment:密码" json:"password_hash"`
}

// TableName table name
func (s *User) TableName() string {
	return "srv_user"
}

// CreateTableMigrator create table migrator
func (s *User) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewCreateTable(migrator, s)
}

// DropTableMigrator create table migrator
func (s *User) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewDropTable(migrator, s)
}
