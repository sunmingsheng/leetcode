package main

import "fmt"

//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
//
//它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lru-cache-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	cache := Constructor(2)
	cache.Put(1, 1);
	cache.Put(2, 2);
	cache.Get(1);       // 返回  1
	cache.Put(3, 3);    // 该操作会使得密钥 2 作废
	cache.Get(2);       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得密钥 1 作废
	cache.Get(1);       // 返回 -1 (未找到)
	cache.Get(3);       // 返回  3
	cache.Get(4);       // 返回  4
}

type Data struct {
	key int
	value int
	before *Data
	next *Data
}

type LRUCache struct {
	head *Data
	tail *Data
	m map[int]*Data
	size int
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		m : make(map[int]*Data),
		size: capacity,
	}
	c.head = &Data{
		key: -1,
		value:-1,
		before: nil,
		next: nil,
	}
	c.tail = &Data{
		key: -1,
		value:-1,
		before: nil,
		next: nil,
	}
	c.head.next = c.tail
	c.tail.before = c.head
	return c
}


func (this *LRUCache) Get(key int) int {
	if _,ok := this.m[key]; !ok {
		return -1
	} else {
		current := this.m[key]
		current.before.next = current.next
		current.next.before = current.before
		current.before = this.head
		current.next = this.head.next
		this.head.next = current
		return current.value
	}
}


func (c *LRUCache) Put(x int, y int)  {
	if _, ok := c.m[x]; ok {
		c.m[x].value = y
	} else {
		if len(c.m) >= c.size {
			delete(c.m, c.tail.before.key)
			last := c.tail.before
			last.before.next = c.tail
			c.tail.before = last.before
			last.next = nil
		}
		if x == 4 {
			fmt.Println(c.m)
			fmt.Println("--------------------")
		}
		c.m[x] = &Data{
			key: x,
			value: y,
			before: nil,
			next: nil,
		}
		old := c.head.next
		if x == 4 {
			fmt.Println(old)
		}
		c.head.next = c.m[x]
		old.before = c.m[x]
		c.m[x].next = old
		c.m[x].before = c.head
		if x == 4 {
			fmt.Println(c.m)
		}
	}
}
