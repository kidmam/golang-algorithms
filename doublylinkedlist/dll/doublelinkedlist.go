package main

import "sync"
import "fmt"

// List 구조체는 아이템의 묶음입니다.
type List struct {
	head   *Item        // 제일 앞에 있는 아이템입니다.
	last   *Item        // 가장 뒤에 있는 아이템입니다.
	len    int          // List 구조체의 전체 길이입니다.
	locker sync.RWMutex // 데이터(Item)을 등록, 삭제할때 데이터 안정성을 위해 lock을 걸어줍니다.
}

// Item은 저장 공간을 의미합니다.
type Item struct {
	Val  interface{} // 저장할 데이터입니다. go에서 interface{}는 어떤 타입의 데이터든 받을 수 있습니다.
	prev *Item       // 이전 아이템을 지정합니다.
	next *Item       // 다음 아이템을 지정합니다.
	list *List       // 아이템(Item)이 속한 List를 지정합니다.
}

// 새로운 List를 생성합니다.
func New() *List {
	list := &List{} // 빈 List 구조체의 메모리 주소를 대입합니다.
	list.len = 0    // List의 길이는 0입니다.
	return list
}

// List에 Item을 저정합니다.
func Insert(value interface{}, list *List) *List {

	// Item 구조체에 value, 지정한 List의 head, last, 그리고 현재 list 주소를 저장합니다.
	// newItem 변수에 지정한 값을 저정한 Item 변수의 주소를 대입합니다.
	newItem := &Item{value, list.last, list.head, list}

	// 입력을 위해 list에 Lock을 걸어줍니다. 그리고 함수가 종료될때 Unlock해 줍니다.
	list.locker.Lock()
	defer list.locker.Unlock()

	// 만약 list의 head가 nil이라면 빈 List였으며 처음으로 Item을 등록하는 것입니다.
	// list의 head와 last에 입력하는 Item을 지정합니다.
	if list.head == nil {
		list.head = newItem
		list.last = newItem
	} else {
		list.last.next = newItem
		list.head.prev = newItem
		list.head = newItem
	}

	list.len++

	return list
}

// List의 가장 앞에있는 Item을 불러옵니다.
func (list *List) First() *Item {
	return list.head
}

// List의 가장 뒤에있는 Item을 불러옵니다.
func (list *List) Last() *Item {
	return list.last
}

// 현제 Item의 이전 Item을 불러 옵니다.
func (item *Item) Prev() *Item {
	return item.prev
}

// 현재 Item의 다음 Item을 불러옵니다.
func (item *Item) Next() *Item {
	return item.next
}

// List에서 저정한 Item이 있는지 확인합니다.
func Has(value interface{}, list *List) bool {
	if list.head == nil {
		return false
	}
	first := list.First()

	for {
		if first.Val == value {
			return true
		} else {
			if first.next != nil {
				first = first.next
			} else {
				return false
			}
		}
	}
}

// List에서 지정한 Item을 삭제합니다.
// RLock/RUnlock을 사용해 읽기에 대한 보장을 받게합니다. 내가 읽을때 다른 사람은 쓰지 못하게 합니다.
func Remove(value interface{}, list *List) *List {

	list.locker.RLock()
	if list.head == nil {
		return list
	}
	list.locker.RUnlock()

	list.locker.RLock()
	first := list.First()
	last := list.Last()
	list.locker.RUnlock()

	list.locker.Lock()
	defer list.locker.Unlock()

	for {
		if last.next == nil {
			return list
		}

		if first.Val == value {
			first.next.prev = first.prev
			first.prev.next = first.next
			first.prev = nil
			first.next = nil
			first.Val = nil
			first.list = nil
			list.len--
			return list
		} else {
			first = first.next
		}
	}
}

// List의 길이를 확인합니다.
func Length(list *List) int {
	return list.len
}

func main() {
	list := New()
	list = Insert(1, list)
	list = Insert(2, list)
	list = Insert(3, list)
	list = Insert(10, list)
	list = Insert(103, list)
	list = Insert(56, list)
	has := Has(103, list)
	fmt.Println("103있냐? ", has)
	fmt.Println("현재길이: ", Length(list))
	list = Remove(103, list)
	first := list.First()
	fmt.Println("1번 데이터:", first.Val)
	fmt.Println("다음 데이터: ", first.next.Val)
	len := Length(list)
	fmt.Println("현재길이: ", len)
}
