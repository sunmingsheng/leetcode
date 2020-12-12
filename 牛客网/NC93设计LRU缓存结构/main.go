package main

import "fmt"

//设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
//set(key, value)：将记录(key, value)插入该结构
//get(key)：返回key对应的value值
//[要求]
//set和get方法的时间复杂度为O(1)
//某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
//当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
//若opt=1，接下来两个整数x, y，表示set(x, y)
//若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
//对于每个操作2，输出一个答案

func main() {
	//3 2 1
	//1 3 4
	fmt.Println(LRU([][]int{{1,1,1},{1,2,2},{1,3,3},{2,1},{1,4,4},{2,4}}, 3))
}

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU( operators [][]int ,  k int ) []int {
	res := []int{}
	c := New(k)
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			c.set(operators[i][1],operators[i][2])
		}
		if operators[i][0] == 2 {
			res = append(res, c.get(operators[i][1]))
		}
	}
	return res
}

type Data struct {
	key int
	value int
	before *Data
	next *Data
}

type Container struct {
	head *Data
	tail *Data
	m map[int]*Data
	size int
}

func New(k int) *Container{
	c := &Container{
		m : make(map[int]*Data),
		size: k,
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

func (c *Container) set(x,y int) {
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
		c.m[x] = &Data{
			key: x,
			value: y,
			before: nil,
			next: nil,
		}
		c.m[x].before = c.head
		c.m[x].next = c.head.next
		c.head.next.before = c.m[x]
		c.head.next = c.m[x]
	}
}

func (c *Container) get(x int) int{
	if _,ok := c.m[x]; !ok {
		return -1
	} else {
		current := c.m[x]
		current.before.next = current.next
		current.next.before = current.before
		c.m[x].before = c.head
		c.m[x].next = c.head.next
		c.head.next = c.m[x]
		return c.m[x].value
	}
}