package main

import "fmt"

// Node 구조체는 이진트리의 노드를 구성합니다.
type Node struct {
	Key   int
	Left  *Node // 왼쪽 노드는 현재 노드보다 작은 수를 갖습니다.
	Right *Node // 오른쪽 노드는 현재 노드보다 큰 수를 갖습니다.
}

// 이진 탐색 트리에서 값을 검색합니다.
func (n *Node) Search(searchKey int) bool {
	// 노드 구조체 n이 nil이라면 아루 값도 없는 경우 입니다.
	if n == nil {
		return false
	}

	// 찾으려는 값이 노드의 key보다 크다면 오른쪽 노드에서 다시 찾습니다.
	// 그렇지 않다면 (찾으려는 값이 노드의 key보다 작다면) 왼쪽 노드에서 다시 찾습니다.
	if n.Key < searchKey {
		return n.Right.Search(searchKey)
	} else if n.Key > searchKey {
		return n.Left.Search(searchKey)
	}

	// 위 조건에 포함하지 않는다면 n.key는 찾으려는 값입니다.
	// n.Key == key
	return true
}

// 이진 탐색 트리에 값을 추가합니다.
func (n *Node) Insert(addKey int) {

	// 추가하려는 값이 노드의 key보다 크다면 오른쪽 노드에서 다시 추가 조건을 확인합니다.
	// 그렇지 않다면 (추가하려는 값이 노드의 key보다 작다면) 왼쪽 노드에서 다시 추가 조건을 확인합니다.
	// 오른쪽 또는 왼쪽 노드가 nil이라면 새로운 노드를 추가합니다.
	if n.Key < addKey {
		if n.Right == nil {
			n.Right = &Node{Key: addKey}
		} else {
			n.Right.Insert(addKey)
		}
	} else if n.Key > addKey {
		if n.Left == nil {
			n.Left = &Node{Key: addKey}
		} else {
			n.Left.Insert(addKey)
		}
	}
	// 만약 추가하려는 노드가 있을 경우 아무런 동작을 하지 않습니다.
	// n.Key == key
}

// 이진 탐색 트리에서 가장 작은값을 찾습니다.
// 트리의 왼쪽 끝에 있는 값이 가장 작은 값입니다.
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// 이진 탐색 트리에서 가장 큰값을 찾습니다.
// 트리의 오른쪽 끝에 있는 값이 가장 큰 값입니다.
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}

// 노드를 삭제합니다.
func (n *Node) Delete(removeKey int) *Node {

	// 삭제하려는 값이 노드 key보다 크다면 오른쪽 노드에서 삭제 작업을 진행합니다.
	// 만약 삭제하려는 값이 노드 key보다 작다면 왼쪽 노드에서 삭제 작업을 진행합니다.
	if n.Key < removeKey {
		n.Right = n.Right.Delete(removeKey)
	} else if n.Key > removeKey {
		n.Left = n.Left.Delete(removeKey)
		// 삭제하려는 key를 찾았을 경우
		// 대상 노드의 왼쪽 노드가 없을 경우 반대쪽인 오른쪽 노드를 리턴합니다.
		// 만약 대상 노드의 오른쪽이 없을 경우 반대쪽인 왼쪽 노드를 리턴합니다.
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// 먄악 대상 노드의 왼쪽과 오른쪽 노드가 모두 있을 경우
		// 오른쪽 노드에서 가장 작은 값을 확인합니다.
		min := n.Right.Min()

		// 대상 key를 오른쪽에서 찾은 가장 작은값으려 변경하고
		n.Key = min

		// 오른쪽에 있는 가장 작은 값을 삭제합니다.
		n.Right = n.Right.Delete(min)
	}

	return n
}

func main() {

	tree := &Node{6, nil, nil}
	tree.Insert(5)
	tree.Insert(8)
	fmt.Println(tree.Search(8))

	tree.Insert(12)
	tree.Insert(3)
	tree.Delete(5)
	fmt.Println(tree.Search(5))
}
