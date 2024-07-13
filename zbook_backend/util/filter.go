package util

func IsInt64InArray(value int64, array []int64) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
