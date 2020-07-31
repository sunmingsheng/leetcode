package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(isValid("(([]){})"))
}

type StackInterface interface {
	//栈顶弹出一个元素
	Pop() interface{}

	//向栈中压入一个元素
	Push(interface{})

	//获取栈顶的元素,但是不弹出
	Peek()

	//打印元素,从栈底到栈顶
	Print()

	//获取栈的长度
	Length() int

	//判断栈是不是为空
	Empty() bool
}


type Stack struct {
	//记录栈中元素的个数
	length int

	//数据存储
	data   []interface{}
}

//栈顶弹出一个元素
func (stack *Stack) Pop() interface{} {
	if stack.length <= 0 {
		return nil
	}
	returnData   := stack.data[stack.length - 1]
	stack.data    = append(stack.data[ : stack.length - 1],stack.data[stack.length : ]...)
	stack.length -= 1
	return returnData
}

//向栈中压入一个元素
func (stack *Stack) Push (data interface{}) {
	stack.data    = append(stack.data,data)
	stack.length += 1
}

//获取栈顶的元素,但是不弹出
func (stack *Stack) Peek() (data interface{}) {
	if stack.length <= 0 {
		return nil
	}
	return stack.data[stack.length - 1]
}

//打印元素,从栈底到栈顶
func (stack *Stack) Print() () {
	fmt.Println(stack.data)
}

//获取栈的长度
func (stack *Stack) Length() (length int) {
	length = stack.length
	return
}

//判断栈是不是为空
func (stack *Stack) Empty() (isEmpty bool) {
	if stack.length == 0 {
		isEmpty = false
	}
	isEmpty = true
	return
}


func isValid(content string) bool {
	//(,),{,},[,]
	newStack := Stack{}
	runes    := []rune(content)

	for i := 0; i < len(runes); i++ {
		tempRune := string(runes[i])
		if  tempRune == "(" || tempRune == "{" || tempRune == "[" {
			newStack.Push(tempRune)
		} else {
			popRune,_ := newStack.Peek().(string)
			if (tempRune == ")" && popRune == "(") || (tempRune == "}" && popRune == "{") || (tempRune == "]" && popRune == "[") {
				newStack.Pop()
			} else {
				return false
			}
		}
	}

	if newStack.length > 0 {
		return false
	}
	return true
}