package model

import "time"

type IO struct {
	ID            string    `json:"id"`
	OperationType string    `json:"operationType"`
	Category      string    `json:"category"`
	Name          string    `json:"name"`
	Value         float64   `json:"value"`
	Date          time.Time `json:"time,omitempty"`
}
