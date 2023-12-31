package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func Len(head *Node) int {
	count := 0
	current := head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	var prev *Node

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

func PushBack(head *Node, data int) *Node {
	newNode := &Node{data: data, next: nil}

	if head == nil {
		head = newNode
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}

	return head
}

func PushFront(head *Node, data int) *Node {
	newNode := &Node{data: data, next: head}
	head = newNode

	return head
}

func PrintList(head *Node) {
	current := head

	for current != nil {
		fmt.Print(current.data)
		current = current.next
		if current != nil {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {

	var head *Node

	for i := 1; i <= 5; i++ {
		// head = PushBack(head, rand.Intn(10))
		head = PushFront(head, i)
	}

	PrintList(head)

}
