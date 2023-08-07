package main

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	index, flag := 0, -1
	ans := make([]string, numRows)
	for _, str := range s {
		ans[index] += string(str)
		if index == 0 || index == numRows-1 {
			flag = -flag
		}
		index += flag
	}
	res := ""
	for _, str1 := range ans {
		res += str1
	}
	return res
}
