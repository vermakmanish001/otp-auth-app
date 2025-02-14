package models

import "time"

type User struct {
	Phone     string    `bson:"phone"`
	OTP       string    `bson:"otp"`
	ExpiresAt time.Time `bson:"expires_at"`
	DeviceID  string    `bson:"device_id"`
}
