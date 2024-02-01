package model

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
