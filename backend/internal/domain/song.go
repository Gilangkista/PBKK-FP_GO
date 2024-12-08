package domain

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title      string     `gorm:"type:varchar(255);not null"`
	Slug       string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	ArtistID   uint       `gorm:"not null"`
	CategoryID uint       `gorm:"not null"`
	Artist     Artist     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key to Artist
	Category   Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key to Category
	Playlists  []Playlist `gorm:"many2many:playlist_songs;"`                      // Many-to-Many relationship
}
