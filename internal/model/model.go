package model

import "time"

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ConditionsT map[string]any
type Predicates map[string][]any

type Price int
