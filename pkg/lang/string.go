package lang

import "regexp"

func IsInt(s string) bool {
	return regexp.MustCompile("^\\d+$").MatchString(s)
}

func IsFloat(s string) bool {
	return regexp.MustCompile("^\\d*[.]\\d*$").MatchString(s)
}
