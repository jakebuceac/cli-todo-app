package helpers

import (
	"time"

	"github.com/mergestat/timediff"
)

func CalculateTimeDifference(timestamp string) (string, error) {
	time, err := time.Parse("2006-01-02T15:04:05-07:00", timestamp)

	if err != nil {
		return "", err
	}

	return timediff.TimeDiff(time), nil
}
