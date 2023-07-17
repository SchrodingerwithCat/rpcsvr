package model

import (
	"time"
)

// model
type Student struct {
	Id             int32
	Name           string
	Email          string
	CollegeName    string
	CollegeAddress string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}