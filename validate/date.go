package validate

import "github.com/elancom/go-util/lang"

func IsDate(s string) bool {
	_, err := lang.ToDate(s)
	return err == nil
}

func IsDateTime(s string) bool {
	_, err := lang.ToDateTime(s)
	return err == nil
}

func IsDateOrTime(s string) bool {
	return IsDateTime(s) || IsDate(s)
}
