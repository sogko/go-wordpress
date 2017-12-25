package wordpress

import (
	"time"
)

type Time struct {
	time.Time
}

// 2017-12-25T09:54:42
const timeLayout = "2006-01-02T15:04:05"

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	t.Time, err = time.Parse(timeLayout, string(b))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(timeLayout)), nil
}
