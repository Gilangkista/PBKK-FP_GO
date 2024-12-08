package domain

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);not null"`
	Slug  string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Songs []Song `gorm:"foreignKey:ArtistID"` // One-to-Many relationship
}
