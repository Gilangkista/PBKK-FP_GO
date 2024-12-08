package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);not null"`
	Slug  string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Songs []Song `gorm:"foreignKey:CategoryID"` // One-to-Many relationship
}
