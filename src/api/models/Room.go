package models

import "time"

type Room struct {
	Id        int    `gorm:"primary_key"`
	Title     string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Level     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
