// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLoginLog = "login_log"

// LoginLog mapped from table <login_log>
type LoginLog struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LoginUsername string    `gorm:"column:login_username" json:"login_username"`
	Ipaddr        string    `gorm:"column:ipaddr" json:"ipaddr"`
	LoginLocation string    `gorm:"column:login_location" json:"login_location"`
	Browser       string    `gorm:"column:browser" json:"browser"`
	Os            string    `gorm:"column:os" json:"os"`
	Status        int32     `gorm:"column:status" json:"status"`
	Msg           string    `gorm:"column:msg" json:"msg"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName LoginLog's table name
func (*LoginLog) TableName() string {
	return TableNameLoginLog
}
