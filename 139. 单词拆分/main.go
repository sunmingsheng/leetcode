package main

import "fmt"

//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/word-break
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	length := len(wordDict)
	for i := 0; i < length; i++ {
		for j := 0; j < length - 1; j++ {
			if len(wordDict[j]) < len(wordDict[j+1]) {
				wordDict[j], wordDict[j+1] = wordDict[j+1], wordDict[j]
			}
		}
	}
	m := make(map[string]struct{})
	return wordBreak_(s, wordDict, m)
}

func wordBreak_(s string, wordDict []string, m map[string]struct{}) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	for _, value := range wordDict {
		wordLen := len(value)
		if wordLen > length {
			continue
		}
		if _, ok := m[s[wordLen:]]; ok {
			continue
		}
		if s[:wordLen] == value {
			res := wordBreak_(s[wordLen:], wordDict, m)
			if res {
				return true
			} else {
				m[s[wordLen:]] = struct{}{}
			}
		}
	}
	return false
}


