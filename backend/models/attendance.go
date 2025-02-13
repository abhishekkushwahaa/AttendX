package models

import "time"

type Attendance struct {
	Username  string    `bson:"username"`
	Timestamp time.Time `bson:"timestamp"`
	Method    string    `bson:"method"`
}
