package models

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"type:varchar(255);not null"`
	Author string `gorm:"type:varchar(255);not null"`
}
