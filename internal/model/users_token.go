package model

import "time"

type UserToken struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	UserID             uint      `json:"userId"`
	User               User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Token              string    `json:"token"`
	RefreshToken       string    `json:"refreshToken"`
	ExpiredTime        time.Time `json:"expiredTime"`
	RefreshExpiredTime time.Time `json:"refreshExpiredTime"`
	CreatedDate        time.Time `json:"createdDate" gorm:"autoCreateTime"`
	UpdateDate         time.Time `json:"updateDate" gorm:"autoUpdateTime"`
}
