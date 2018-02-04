package wordpress

import (
	"time"
)

var Location = time.UTC

type Time struct {
	time.Time
}

// 2017-12-25T09:54:42
const TimeLayout = "2006-01-02T15:04:05"

// "2017-09-24T13:28:06+00:00"
const TimeWithZoneLayout = "2006-01-02T15:04:05-07:00"

func (t *Time) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	zoneTime, err := time.Parse(TimeWithZoneLayout, string(b))
	if err != nil {
		noZoneTime, altErr := time.ParseInLocation(TimeLayout, string(b), Location)
		if altErr != nil {
			return err
		} else {
			zoneTime = noZoneTime
		}
	}
	t.Time = zoneTime
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(TimeLayout)), nil
}
