package models

import "time"

type Todos struct {
	ID          uint      `gorm: "primaryKey"`
	Title       string    `gorm: "<-"`
	Description string    `gorm: "<-"`
	CreatedAt   time.Time `gorm: "autoCreateTime"`
	UpdatedTime time.Time `gorm: "autoUpdateTIme:milli"`
}
