package function

/*
* description: 判断切片中是否包含指定字符串
* created on: 2021/3/31 14:45
 */
func InSliceStr(member string, list []string) bool {
	for _, element := range list {
		if element == member {
			return true
		}
	}
	return false
}
