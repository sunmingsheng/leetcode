package main

import "fmt"

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//示例:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Trie struct {
	val byte //数组值
	nexts map[byte]*Trie //指向接下来元素的指针集合
	isEnd bool
}

func main() {
	obj := Constructor();
	obj.Insert("abcdf");
	fmt.Println(obj.Search("abcdf"))
	fmt.Println(obj.StartsWith("abc"))
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		nexts: make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	bytes := []byte(word)
	if len(bytes) <= 0 {
		return
	}
	for key, value := range bytes {
		//如果数据存在
		if _, ok := this.nexts[value]; !ok {
			//如果数据不存在，创建
			this.nexts[value] = &Trie{
				nexts: make(map[byte]*Trie),
			}
		}
		if key == len(bytes) - 1 {
			this.nexts[value].isEnd = true
		}
		this = this.nexts[value]
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	bytes := []byte(word)
	if len(bytes) <= 0 {
		return true
	}
	for key, value := range bytes {
		//如果数据存在
		if _, ok := this.nexts[value]; !ok {
			return false
		}
		if key == len(bytes) - 1 {
			if this.nexts[value].isEnd == true {
				return true
			} else {
				return false
			}
		}
		this = this.nexts[value]
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	bytes := []byte(prefix)
	if len(bytes) <= 0 {
		return true
	}
	for _, value := range bytes {
		//如果数据存在
		if _, ok := this.nexts[value]; !ok {
			return false
		}
		this = this.nexts[value]
	}
	return true
}