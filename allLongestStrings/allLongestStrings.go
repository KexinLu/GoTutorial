package allLongStrings

func allLongStrings(inputArray []string) []string {
	longlen := 0
	result := make([]string, 0)
	for _, str := range inputArray {
		if len(str) > longlen {
			result = make([]string, 0)
			longlen = len(str)
			result = append(result, str)
		} else if len(str) == longlen {
			result = append(result, str)
		}
	}
	return result
}
