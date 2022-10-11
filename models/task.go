package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID    string `json:"id"`
	Title string `json:"title"`
}
