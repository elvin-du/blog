package utils

import "strings"

const (
	C_EXCERPT_TAG = "__EXCERPT__TAG__"
)

//template func
func Excerpt(in string) string {
	i := strings.Index(in, C_EXCERPT_TAG)
	if -1 != i {
		return in[:i]
	}
	return in
}

func YYYYMMDD(date string) string {
	feilds := strings.Fields(date)
	if len(feilds) >= 1 {
		return feilds[0]
	}
	return date
}
