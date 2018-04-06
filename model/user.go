package model

import (
	"time"
)

type User struct {
	ID             uint32
	Email          string
	PassHash       []byte
	EmailConfirmed bool
	IsAdmin        bool
	LoginAllowed   bool
	InternalNotes  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

// Create User
// Update User
// Read User
// Destroy User

// Hash Password
// Validate Password
