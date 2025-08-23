package model

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	IsActive    bool      `json:"isActive" gorm:"default:true"`
	CreatedDate time.Time `json:"createdDate" gorm:"autoCreateTime"`
	UpdateDate  time.Time `json:"updateDate" gorm:"autoUpdateTime"`
}