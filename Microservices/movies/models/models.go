package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Genre       string `json:"genre"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Release     string `json:"release"`
	UserID      int    `json:"user_id"`
}
