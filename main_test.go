package main

import (
	"testing"
	"time"
)

func TestTimeToEmoji(t *testing.T) {
	for _, c := range []struct {
		in   time.Time
		want string
	}{
		{time.Date(2023, 9, 8, 12, 30, 0, 0, time.UTC), ":clock1230:"},
		{time.Date(2023, 9, 8, 12, 45, 0, 0, time.UTC), ":clock1230:"},
		{time.Date(2023, 9, 8, 12, 59, 0, 0, time.UTC), ":clock1:"},
		{time.Date(2023, 9, 8, 12, 35, 0, 0, time.UTC), ":clock1230:"},
		{time.Date(2023, 9, 8, 12, 15, 0, 0, time.UTC), ":clock12:"},
		{time.Date(2023, 9, 8, 12, 0, 0, 0, time.UTC), ":clock12:"},
		{time.Date(2023, 9, 8, 12, 20, 0, 0, time.UTC), ":clock1230:"},
		{time.Date(2023, 9, 8, 12, 29, 0, 0, time.UTC), ":clock1230:"},
		{time.Date(2023, 9, 8, 13, 50, 0, 0, time.UTC), ":clock2:"},
		{time.Date(2023, 9, 8, 15, 30, 0, 0, time.UTC), ":clock330:"},
		{time.Date(2023, 9, 8, 18, 20, 0, 0, time.UTC), ":clock630:"},
		{time.Date(2023, 9, 8, 2, 29, 0, 0, time.UTC), ":clock230:"},
	} {
		t.Run(c.in.String(), func(t *testing.T) {
			got := timeToEmoji(c.in)
			if c.want != got {
				t.Fatalf("want %s, got %s", c.want, got)
			}
		})
	}
}
