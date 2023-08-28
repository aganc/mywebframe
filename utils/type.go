package utils

import (
	"fmt"
	"strconv"
)

func ConvertInterface2Int64(value interface{}) (int64, error) {
	return strconv.ParseInt(fmt.Sprint(value), 10, 64)
}

func ConvertInterface2UInt64(value interface{}) (uint64, error) {
	return strconv.ParseUint(fmt.Sprint(value), 10, 64)
}

func ConvertInterface2Float32(value interface{}) (float32, error) {
	f, err := strconv.ParseFloat(fmt.Sprint(value), 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func ConvertInterface2UInt32(value interface{}) (uint32, error) {
	i, err := strconv.ParseUint(fmt.Sprint(value), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func ConvertInterface2UInt8(value interface{}) (uint8, error) {
	i, err := strconv.ParseUint(fmt.Sprint(value), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

func ConvertInterface2String(value interface{}) string {
	return fmt.Sprint(value)
}
