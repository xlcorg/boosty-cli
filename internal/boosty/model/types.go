package model

import "time"

type BoolValue bool

func (v BoolValue) String() string {
	return FormatBool(bool(v))
}

func FormatBool(value bool) string {
	if value {
		return "ДА"
	}
	return "НЕТ"
}

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}
