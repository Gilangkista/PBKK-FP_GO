package domain

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null"`
	Slug        string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	Description *string `gorm:"type:text"`
	Songs       []Song  `gorm:"many2many:playlist_songs;"` // Many-to-Many relationship
}
