package model

import "time"

type Attendance struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"userId" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"unique;not null"`
	Type        string    `json:"type" gorm:"not null"`
	CheckIn     time.Time `json:"checkIn"`
	CheckOut    time.Time `json:"checkOut"`
	Status      string    `json:"status" gorm:"not null"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedByID uint      `json:"createdById"`
	UpdatedDate time.Time `json:"updateDate"`
	UpdatedByID uint      `json:"updateById"`
}
