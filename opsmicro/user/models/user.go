package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint `gorm:"primarykey" json:"id"`
	Username string `gorm:"type:varchar(50);not null;unique;uniqueIndex" validate:"required,min=3,max=50" json:"username"`
	CnName string `gorm:"type:varchar(100)" json:"cn_name"`
	Password string `gorm:"type:varchar(100);not null" validate:"required,min=8,max=50,pwdcomplex" json:"password"`
	Email string `validate:"omitempty,email" json:"email"`
	IsAdmin *bool `gorm:"not null;default:false" validate:"exists" json:"is_admin"`
	Role string `gorm:"type:varchar(20);not null" validate:"required,roleoptions" json:"role"`
	Phone string `gorm:"type:char(11);unique;not null;uniqueIndex" validate:"required,len=11,phonecheck" json:"phone"`
	IsActive *bool `gorm:"not null;default:true" validate:"exists" json:"is_active"`
	LastLogin *time.Time `json:"last_login"`
	CreatedTime time.Time `grom:"autoCreateTime" json:"create_time"`
	UpdatedTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"index" json:"delete_time"`
}

