package utils

import (
	"fmt"
	"strconv"
)

//FormatFloat64 decimalPlaces 保留小数位数
func FormatFloat64(number float64, decimalPlaces int) float64 {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(decimalPlaces)+"f", number), 64)
	return f
}

type FormatPercentResult struct {
	D float64
	S string
}

// FormatPercent 计算百分比 90% pDecimalPlaces 计算出来的百分比需要保留的小数位数
func FormatPercent(number, total float64, pDecimalPlaces int) FormatPercentResult {
	result := FormatPercentResult{
		D: 0,
		S: "0%",
	}
	if total == 0 || number == 0 {
		return result
	}
	if number >= total {
		result.D = 100
		result.S = "100%"
		return result
	}
	if pDecimalPlaces < 0 {
		pDecimalPlaces = 0
	}
	// 计算
	result.D = FormatFloat64((number/total)*100, pDecimalPlaces)
	result.S = fmt.Sprintf("%v%v", result.D, "%")
	return result
}

func FormatPercentString(total float64, val float64) string {
	if total == 0 {
		return "0.00"
	}
	return fmt.Sprintf("%.2f", (val/total)*100)
}
