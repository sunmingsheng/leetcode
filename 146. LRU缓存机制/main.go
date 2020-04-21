package main

import "fmt"

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
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
//链接：https://leetcode-cn.com/problems/lru-cache
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	l := Constructor(1)
	fmt.Println(l.Get(1))
	l.Put(1,11)
	fmt.Println(l.Get(1))
	l.Put(1,22)
	fmt.Println(l.Get(1))
	l.Put(2,33)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
}

type Data struct {
	val    int
	before *Data
	next   *Data
}

type LRUCache struct {
	m        map[int]*Data //map
	capacity int           //容量
	head     *Data         //头指针
	tail     *Data         //尾指针
}

func Constructor(capacity int) LRUCache {
	head := &Data{
		val:    0,
		before: nil,
		next:   nil,
	}
	tail := &Data{
		val:    0,
		before: nil,
		next:   nil,
	}
	head.next = tail
	tail.before = head
	return LRUCache{
		m:        make(map[int]*Data),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.m[key]; ok {
		value.before.next = value.next
		value.next.before = value.before
		old := this.head.next
		this.head.next = value
		value.before = this.head
		value.next = old
		old.before = value
		return value.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.m[key]; ok {
		val.val = value
	} else {
		if len(this.m) >= this.capacity {
			//删除最后一个元素
			d := this.tail.before
			d.before.next = this.tail
			this.tail.before = d.before
			delete(this.m, d.val)
		}
		data := &Data{
			val:    value,
			before: nil,
			next:   nil,
		}
		this.m[key] = data
		old := this.head.next
		this.head.next = data
		old.before = data
		data.next = old
		data.before = this.head
	}
}
