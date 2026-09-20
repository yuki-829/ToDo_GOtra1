package model

type Todo struct {
	ID   uint   `gorm:"primaryKey"`
	Task string `gorm:"not null"`
}
