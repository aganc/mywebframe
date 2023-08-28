/**************************************************************
 * Copyright (c) 2021 anxin.com, Inc. All Rights Reserved
 * User: zhangdongsheng<zhangdongsheng@anxin.com>
 * Date: 2021/9/5
 * Desc:
 **************************************************************/

package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//struct 自动把int64切片转成string
type Int64StringSlice []int64

func (slice Int64StringSlice) MarshalJSON() ([]byte, error) {
	values := make([]string, len(slice))
	for i, value := range []int64(slice) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}

	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (slice *Int64StringSlice) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []int64
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*slice = values
		return nil
	}
	*slice = make([]int64, len(values))
	for i, value := range values {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		(*slice)[i] = value
	}
	return nil
}

type IntStringSlice []int

func (slice IntStringSlice) MarshalJSON() ([]byte, error) {
	values := make([]string, len(slice))
	for i, value := range []int(slice) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}

	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (slice *IntStringSlice) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []int
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*slice = values
		return nil
	}
	*slice = make([]int, len(values))
	for i, value := range values {
		value, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		(*slice)[i] = value
	}
	return nil
}
