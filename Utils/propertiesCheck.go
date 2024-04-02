package Utils

import (
	"regexp"
)

func EmptyString(str string) (isEmpty bool) {
	return str == ""
}

func NotEmailString(email string) (isNotEmail bool) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return !emailRegex.MatchString(email)
}
