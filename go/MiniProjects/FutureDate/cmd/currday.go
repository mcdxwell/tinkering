package cmd

import (
	"time"
)

func GetCurrDate() time.Time {
	current_date := time.Now().Local()
	return current_date
}
