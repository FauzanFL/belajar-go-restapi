package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
