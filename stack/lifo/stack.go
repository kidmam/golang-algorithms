package main

import "fmt"

// StackItem 구조체입니다. 값과 다음 StackItem을 가지고 있습니다.
type StackItem struct {
	item interface{}
	next *StackItem
}

// Stack 구조체는 StackItem과 Stack의 깊이를 가지고 있습니다.
type Stack struct {
	sp    *StackItem // Stack에서 가장 위에 있는 아이템을 저정합니다.
	depth uint64     // Stack에 아이템이 몇개 있는지를 저장합니다.
}

// Stack을 생성합니다.
func New() *Stack {
	var stack *Stack = new(Stack)

	stack.depth = 0
	return stack
}

// Stack에 아이템을 집어 넣습니다.
func (stack *Stack) Push(item interface{}) {
	stack.sp = &StackItem{item: item, next: stack.sp}
	stack.depth++
}

// Stack에서 가장 위에 있는 아이템을 제거합니다.
// 가장 위에있는 아이템을 제거하고 그 다음 아이템을 위로 올립니다.
func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}

	return nil
}

// Stack의 가장 위에 있는 아이템을 확인합니다.
// 아이템을 제거되지 않고 읽기만 합니다.
func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.item
	}

	return nil
}

func main() {
	var stack *Stack = New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	// 가장 나중에 들어간 데이터가 가장 먼저 나옵니다. (LIFO)
	for i := 5; i > 0; i-- {
		item := stack.Pop()

		fmt.Println(item)
	}
}
