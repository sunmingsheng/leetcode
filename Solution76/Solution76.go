package main

import (
	"fmt"
	"time"
)

//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-window-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//s1 := "wegdtzwabazduwwdysdetrrctotpcepalxdewzezbfewbabbseinxbqqplitpxtcwwhuyntbtzxwzyaufihclztckdwccpeyonumbpnuonsnnsjscrvpsqsftohvfnvtbphcgxyumqjzltspmphefzjypsvugqqjhzlnylhkdqmolggxvneaopadivzqnpzurmhpxqcaiqruwztroxtcnvhxqgndyozpcigzykbiaucyvwrjvknifufxducbkbsmlanllpunlyohwfsssiazeixhebipfcdqdrcqiwftutcrbxjthlulvttcvdtaiwqlnsdvqkrngvghupcbcwnaqiclnvnvtfihylcqwvderjllannflchdklqxidvbjdijrnbpkftbqgpttcagghkqucpcgmfrqqajdbynitrbzgwukyaqhmibpzfxmkoeaqnftnvegohfudbgbbyiqglhhqevcszdkokdbhjjvqqrvrxyvvgldtuljygmsircydhalrlgjeyfvxdstmfyhzjrxsfpcytabdcmwqvhuvmpssingpmnpvgmpletjzunewbamwiirwymqizwxlmojsbaehupiocnmenbcxjwujimthjtvvhenkettylcoppdveeycpuybekulvpgqzmgjrbdrmficwlxarxegrejvrejmvrfuenexojqdqyfmjeoacvjvzsrqycfuvmozzuypfpsvnzjxeazgvibubunzyuvugmvhguyojrlysvxwxxesfioiebidxdzfpumyon"
	//s2 := "ozgzyywxvtublcl"
	s1 := "ADOBECODEBANC"
	s2 := "ABC"
	fmt.Println(time.Now().UnixNano())
	fmt.Println(minWindow(s1, s2))
	fmt.Println(time.Now().UnixNano())
	//接近3秒的耗时
	//tcnvhxqgndyozpcigzykbiaucyvwrjvknifufxducbkbsmlanl
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	targetMap := make(map[rune]int)
	for _, value := range []rune(t) {
		if count, ok := targetMap[value]; ok {
			targetMap[value] = count + 1
		} else {
			targetMap[value] = 1
		}
	}
	fmt.Println(targetMap)
	minLength := 0
	results := ""
	sourceRunes := []rune(s)
    length := len(sourceRunes)
    targetLength := len([]rune(t))
	i := 0
	j := 0
    for i < length {
		j = i + targetLength - 1
    	//if j < i + targetLength - 1 {
		//	j = i + targetLength - 1
		//} else {
		//	j = j - 1
		//}
		fmt.Println(i, j)
		outFlag := false
    	for j < length {
			tempRunes := sourceRunes[i:j+1]
			tempMap := make(map[rune]int)
			for _, value := range tempRunes {
				if count, ok := tempMap[value]; ok {
					tempMap[value] = count + 1
				} else {
					tempMap[value] = 1
				}
			}
			//进行数据比对
			flag := true
			for key, value := range targetMap {
				if count, ok := tempMap[key]; !ok || (ok && count < value) {
					flag = false
					break
				} else {
					outFlag = true
				}
			}
			fmt.Println(tempMap, targetMap, flag, minLength)
			if flag && (minLength == 0 || len(tempRunes) < minLength)  {
				minLength = len(tempRunes)
				results = string(tempRunes)
				//寻找到后，意味着找到了以当前字符开头并满足条件的最小子串，跳出循环
				break
			}
			j ++
		}
		if !outFlag {
			break
		}
    	i ++
	}
	return results
}