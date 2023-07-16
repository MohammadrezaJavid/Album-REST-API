package utils

import "strconv"

func StrToInt(str string, base int, bitSize int) int64 {
	i, err := strconv.ParseInt(str, base, bitSize)
	if err != nil {
		panic(err)
	}
	return i

}
