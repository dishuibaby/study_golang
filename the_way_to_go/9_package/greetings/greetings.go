// Package greetings greetings/greetings.go
package greetings

import "time"

func IsAM() bool {
	return time.Now().Hour() < 12
}

func IsAfternoon() bool {
	hour := time.Now().Hour()
	return hour >= 12 && hour < 17
}

func IsEvening() bool {
	return time.Now().Hour() >= 17
}
