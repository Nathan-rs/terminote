package types

import "time"

type Task struct {
	ID           uint32
	Description  string
	CountyID     uint32
	Problem      string
	Solution     string
	DifficultyID uint32
	DateStart    time.Time
	DateEnd      *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type County struct {
	ID    uint32
	Name  string
	State string
}

type Difficulty struct {
	ID   uint32
	Name string
}
