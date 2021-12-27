package convert

import (
	"encoding/json"
	"strconv"
)

func StrToInt(str string) (int, error) {
	res, err := strconv.Atoi(str)
	return res, err
}
func StrToInt64(str string) (int64, error) {
	res, err := strconv.ParseInt(str, 10, 64)
	return res, err
}

func IntToStr(num int) string {
	return strconv.Itoa(num)
}
func Int8ToStr(num int8) string {
	return strconv.Itoa(int(num))
}
func Int16ToStr(num int16) string {
	return strconv.Itoa(int(num))
}
func Int32ToStr(num int32) string {
	return strconv.Itoa(int(num))
}
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
func UintToStr(num uint) string {
	return strconv.Itoa(int(num))
}
func Uint8ToStr(num uint8) string {
	return strconv.Itoa(int(num))
}
func Uint16ToStr(num uint16) string {
	return strconv.Itoa(int(num))
}
func Uint32ToStr(num uint32) string {
	return strconv.Itoa(int(num))
}
func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10)
}
func Float64ToStr(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
func Float32ToStr(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', -1, 64)
}

// Str
// any type convert to string
func Str(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case string:
		key = value.(string)
	case float64:
		ft := value.(float64)
		key = Float64ToStr(ft)
	case float32:
		ft := value.(float32)
		key = Float32ToStr(ft)
	case int:
		it := value.(int)
		key = IntToStr(it)
	case uint:
		it := value.(uint)
		key = UintToStr(it)
	case int8:
		it := value.(int8)
		key = Int8ToStr(it)
	case uint8:
		it := value.(uint8)
		key = Uint8ToStr(it)
	case int16:
		it := value.(int16)
		key = Int16ToStr(it)
	case uint16:
		it := value.(uint16)
		key = Uint16ToStr(it)
	case int32:
		it := value.(int32)
		key = Int32ToStr(it)
	case uint32:
		it := value.(uint32)
		key = Uint32ToStr(it)
	case int64:
		it := value.(int64)
		key = Int64ToStr(it)
	case uint64:
		it := value.(uint64)
		key = Uint64ToStr(it)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
