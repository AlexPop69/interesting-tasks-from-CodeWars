package main

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	h := seconds / 3600
	m := seconds % 3600 / 60
	s := seconds % 3600 % 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
