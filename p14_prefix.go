package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for j := 0; ; j++ {
		stop := false
		for i := 0; i < len(strs); i++ {
			if j == len(strs[i]) {
				stop = true
				break
			}
			if strs[i][j] != strs[0][j] {
				stop = true
				break
			}
		}
		if stop {
			return strs[0][:j]
		}
	}
}

func longestCommonPrefix2(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	for i := 0; i < len(strs[0]); i++ {
		s := strs[0][i]
		flag := false
		for _, v := range strs {
			if i < len(v) {
				if v[i] != s {
					flag = true
				}
			} else {
				flag = true
				break
			}
		}

		if flag == true {
			break
		} else {
			prefix += string(s)
		}
	}

	return prefix
}

func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for j := 0; j < len(strs[0]); j++ {
		c := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == j || strs[i][j] != c {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}

// public String longestCommonPrefix(String[] strs) {
// 	if (strs == null || strs.length == 0) return "";
// 	for (int i = 0; i < strs[0].length() ; i++){
// 		char c = strs[0].charAt(i);
// 		for (int j = 1; j < strs.length; j ++) {
// 			if (i == strs[j].length() || strs[j].charAt(i) != c)
// 				return strs[0].substring(0, i);
// 		}
// 	}
// 	return strs[0];
// }
func main() {
	fmt.Println(longestCommonPrefix3([]string{"a", "ab"}))
}
