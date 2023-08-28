package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// StringSliceToJoinString 字符串切片通过连接符相连
func StringSliceToJoinString(slice []string, join string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", join, -1)
}

// StringSliceToInt64Slice 字符串切片转为 int64切片
func StringSliceToInt64Slice(p []string) (res []int64) {
	for _, v := range p {
		vint, _ := strconv.ParseInt(v, 10, 64)
		res = append(res, vint)
	}
	return res
}
