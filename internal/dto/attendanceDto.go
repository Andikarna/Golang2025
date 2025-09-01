package dto

import "time"

type AttendanceListResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      string    `json:"userId" gorm:"unique;not null"`
	Date        time.Time `json:"date" gorm:"unique;not null"`
	Type        string    `json:"type" gorm:"not null"`
	CheckIn     time.Time `json:"checkIn"`
	CheckOut    time.Time `json:"checkOut"`
	Status      string    `json:"status" gorm:"not null"`
	Description string    `json:"description"`
}
