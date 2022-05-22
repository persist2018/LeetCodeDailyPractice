package main

// https://leetcode.cn/problems/group-anagrams/
// Key point: mp:=make(map[[26]int][]string)
func groupAnagrams(strs []string) (ans [][]string) {
	mp := make(map[[26]int][]string)
	count := [26]int{}
	for _, str := range strs {
		for _, s := range str {
			count[s-'a']++
		}
		mp[count] = append(mp[count], str)
		count = [26]int{}
	}

	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
