package _01910

import "fmt"

/**
https://leetcode-cn.com/problems/longest-palindrome/submissions/
寻找最长回文串长度（a A不是同一个字母）
遍历字符串，计算各个字母出现次数
最长回文串= 字母出现偶数次数+1（如果有的话）
 */

func FindLongestPalindromeString(param string) int {
	var letterNum map[byte]int = make(map[byte]int, 0)
	paramBytes := []byte(param)
	for i := 0; i < len(paramBytes); i++ {
		if num, ok := letterNum[paramBytes[i]]; ok {

			letterNum[paramBytes[i]] = num + 1
		} else {
			letterNum[paramBytes[i]] = 1
		}
		fmt.Println(string(paramBytes[i]),letterNum[paramBytes[i]])
	}

	longestPalindromeStrNum := 0
	hasSingle := false
	for _, num := range letterNum {
		longestPalindromeStrNum = +longestPalindromeStrNum + (num/2)*2
		if num%2 != 0 {
			hasSingle = true
		}
	}
	if hasSingle {
		return longestPalindromeStrNum + 1
	}
	return longestPalindromeStrNum

}
