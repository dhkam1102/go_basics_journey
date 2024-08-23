// BINARY search, BFS, DFS
// TODO: need to prac on some graph types

// ****************************************************************
// Structure will be
// Pseudocode
// function in go
// ****************************************************************

package main

import (
	"fmt"
	"math"

	"github.com/gammazero/deque"
)

// FUNCTION BinarySearch(array, target):
//     left = 0
//     right = length of array - 1
//     WHILE left <= right:
//         mid = (left + right) / 2
//         IF array[mid] == target:
//             RETURN mid
//         ELSE IF array[mid] < target:
//             left = mid + 1
//         ELSE:
//             right = mid - 1
//     RETURN -1  // target not found

func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func buildTree() *Node {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Right = NewNode(6)
	return root
}

// FUNCTION BFS(graph, startNode):
//     queue = new Queue()
//     visited = new Set()
//     queue.enqueue(startNode)
//     visited.add(startNode)
//     WHILE queue is not empty:
//         currentNode = queue.dequeue()
//         PROCESS currentNode  // replace with the logic you want to apply
//         FOR each neighbor in graph[currentNode]:
//             IF neighbor is not in visited:
//                 queue.enqueue(neighbor)
//                 visited.add(neighbor)

func bfs(root *Node) []int {
	result := []int{}
	q := deque.New[*Node]()
	q.PushBack(root)

	for q.Len() > 0 {
		currentNode := q.PopFront()
		result = append(result, currentNode.Value)
		if currentNode.Left != nil {
			q.PushBack(currentNode.Left)
		}
		if currentNode.Right != nil {
			q.PushBack(currentNode.Right)
		}
	}
	return result
}

// FUNCTION DFS(graph, startNode):
//     stack = new Stack()
//     visited = new Set()

//     stack.push(startNode)
//     visited.add(startNode)

//     WHILE stack is not empty:
//         currentNode = stack.pop()
//         PROCESS currentNode  // replace with the logic you want to apply

//         FOR each neighbor in graph[currentNode]:
//             IF neighbor is not in visited:
//                 stack.push(neighbor)
//                 visited.add(neighbor)

func dfs(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{}
	visited := []int{}

	stack = append(stack, root)
	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited = append(visited, currentNode.Value)
		if currentNode.Right != nil {
			stack = append(stack, currentNode.Right)
		}

		if currentNode.Left != nil {
			stack = append(stack, currentNode.Left)
		}
	}

	return visited
}

// FUNCTION DFS(graph, currentNode, visited):
//     PROCESS currentNode  // replace with the logic you want to apply
//     visited.add(currentNode)
//     FOR each neighbor in graph[currentNode]:
//         IF neighbor is not in visited:
//             DFS(graph, neighbor, visited)

func dfsR(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(dfsR(root.Left)), float64(dfsR(root.Right))))
}

func main() {
	// create array for binary search
	// arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	// target := 7
	// result := binarySearch(arr, target)
	// fmt.Println(result)

	// create tree for bfs
	// root := buildTree()
	// result := bfs(root)
	// fmt.Println(result)

	// create tree for dfs
	root := buildTree()
	result := dfs(root)
	fmt.Println(result)

}
