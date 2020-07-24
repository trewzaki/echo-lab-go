package models

import "time"

// StudentScoreName >> Note: Rename {Name} to your nickname
type StudentScoreName struct {
	ID        uint32     `gorm:"AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Subject   string     `gorm:"type:varchar(100);not null" json:"subject"`
	Score     float64    `json:"score"`
	CreatedAt time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}
