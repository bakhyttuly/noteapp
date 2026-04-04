package models

type Note struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title" gorm:"not null"`
	Content    string   `json:"content" gorm:"not null"`
	CategoryID uint     `json:"category_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
