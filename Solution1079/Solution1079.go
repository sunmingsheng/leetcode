package main

import "fmt"

//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//
//示例 1：
//
//输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
//示例 2：
//
//输入："AAABBC"
//输出：188
// 
//
//提示：
//
//1 <= tiles.length <= 7
//tiles 由大写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/letter-tile-possibilities
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(numTilePossibilities("AAABBC"))
}

func numTilePossibilities(tiles string) int {
	m := make(map[string]int)
	runes := []rune(tiles)
	result := []rune{}
	back1079(runes, result, m)
	return len(m) - 1
}

func back1079(runes []rune, result []rune, m map[string]int) {
	if count, ok := m[string(result)]; ok {
		m[string(result)] = count + 1
	} else {
		m[string(result)] = 1
	}
	if len(runes) == 0 {
		return
	}
	for key, value := range runes {
		/*
		tempResult := make([]rune, len(result) + 1)
		copy(tempResult, result)
		tempResult = append(tempResult, value)
		 */
		temp := make([]rune, key)
		copy(temp, runes[:key])
		temp = append(temp, runes[key+1:]...)
		back1079(temp, append(result, value), m)
	}
}