package model

import "time"

type Book struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Title    string    `gorm:"size:255; not null" json:"title" binding:"required"`
	Author   string    `gorm:"size:100; not null" json:"author" binding:"required"`
	Price    uint      `gorm:"default:0" json:"price" binding:"gte=0"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
}
