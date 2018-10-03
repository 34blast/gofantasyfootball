package util


/*
import (
	"fmt"
)
*/

func ReverseString(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
