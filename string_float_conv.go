package general_goutils

import "strconv"

func ConvertStrToFloat64(str string) float64 {
	result, err := strconv.ParseFloat(str,  64)
	if err != nil {
		Logger.Error("failed to convert string to int64: " + err.Error())
		return 0
	}
	return result
}

func ConvertFloat64ToStr(floatValue float64, precision int) string {
	s := strconv.FormatFloat(floatValue, 'f',  precision, 64)
	return s
}