package models

import "time"

type Attendance struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Status    bool      `json:"status" bson:"status"`
}
