package adapter

import "strconv"

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}
