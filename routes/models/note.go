package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/sys/cpu"
)

type Note struct {
	_         cpu.CacheLinePad
	ID        uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title     string     `json:"title" gorm:"not null"`
	Detail    string     `json:"detail"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}
