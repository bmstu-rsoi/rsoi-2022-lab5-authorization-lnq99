package util

import (
	"strconv"
)

// ToInt convert string to int, if there is any err, return 0
func ToInt(n string) int {
	res, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}
