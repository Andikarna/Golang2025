package model

type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}