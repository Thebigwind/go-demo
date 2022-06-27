// Task: Implement a struct named 'RangeList'
//A pair of integers define a range, for example: [1, 5). This range includes integers: 1, 2, 3, and 4.
//A range list is an aggregate of these ranges: [1, 5), [10, 11), [100, 201)
//NOTE: Feel free to add any extra member variables/functions you like.
package main

import (
	"errors"
	"fmt"
)

func main() {

	//rl := RangeList{}
	rl := NewRangeList()
	rl.Add([2]int{1, 5})
	rl.Print()
	// Should display: [1, 5)

	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)

	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 11})
	rl.Print()
	// Should display: [1, 8) [11, 21)

	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)

	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)
}

//type RangeList struct {
//	//TODO: implement
//
//}
//
//func (rangeList *RangeList) Add(rangeElement [2]int) error {
//	//TODO: implement
//	return nil
//}
//
//func (rangeList *RangeList) Remove(rangeElement [2]int) error {
//	//TODO: implement
//	return nil
//}
//func (rangeList *RangeList) Print() error {
//	//TODO: implement
//	return nil
//}
//
//

type RangeList interface {
	// Add an rangeElement to the list.
	Add(rangeElement [2]int) error
	// Remove an rangeElement from the list.
	Remove(rangeElement [2]int) error
	// Print the list of rangeElements.
	Print() error
}

// RangeLinkedList struct
// RangeLinkedList is an implementation of RangeList interface.
type RangeLinkedList struct {
	// head of rangeNode.
	head *rangeNode
}
type rangeNode struct {
	// RangeElement
	rangeElement [2]int
	// Next node
	next *rangeNode
}

//Add TimeComplexity:O(n) SpaceComplexity:O(1)
func (rangeList *RangeLinkedList) Add(rangeElement [2]int) error {
	// Check if the rangeElement is valid.
	if rangeElement[0] >= rangeElement[1] {
		return errors.New("range require [0]<[1],Illegal range")
	}
	// traverse from head to tail.
	var node, pre *rangeNode
	nextRangeElement := [2]int{0, 0}
	for node, pre = rangeList.head.next, rangeList.head; node != nil && rangeElement[0] < rangeElement[1]; node, pre = node.next, pre.next {
		if rangeElement[0] >= node.rangeElement[1] {
			continue
		}
		//divide the range Element into two parts:
		//one part is before node.range Element[1]. (rangeElement)
		//and the other part is after node.range Element[1]. (nextRangeElement)
		nextRangeElement[0] = node.rangeElement[1]
		nextRangeElement[1] = rangeElement[1]
		//we only need to consider one part in this cycle.
		//and the other part well be considered in the next cycle.
		if rangeElement[0] < node.rangeElement[0] {
			if rangeElement[1] < node.rangeElement[0] {
				node = &rangeNode{rangeElement, node}
				pre.next = node
			} else {
				node.rangeElement[0] = rangeElement[0]
			}
			// determine whether node and pre can be merged
			if pre != rangeList.head && node.rangeElement[0] <= pre.rangeElement[1] {
				pre.rangeElement[1] = node.rangeElement[1]
				pre.next = node.next
				node = pre.next
			}
		}
		rangeElement = nextRangeElement
	}
	//if the rangeElement is still valid,add it to the tail.
	if rangeElement[0] < rangeElement[1] {
		// determine whether node and pre can be merged
		if pre != rangeList.head && pre.rangeElement[1] == rangeElement[0] {
			pre.rangeElement[1] = rangeElement[1]
		} else {
			node = &rangeNode{rangeElement, node}
			pre.next = node
		}
	}
	return nil
}

//Remove TimeComplexity:O(n) SpaceComplexity:O(1)
func (rangeList *RangeLinkedList) Remove(rangeElement [2]int) error {
	// Check if the rangeElement is valid.
	if rangeElement[0] >= rangeElement[1] {
		return errors.New("range require [0]<[1],Illegal range")
	}
	// traverse from head to tail.
	var node, pre *rangeNode
	nextRangeElement := [2]int{0, 0}
	for node, pre = rangeList.head.next, rangeList.head; node != nil && rangeElement[0] < rangeElement[1]; {
		if rangeElement[1] < node.rangeElement[0] {
			return nil
		}
		if rangeElement[0] >= node.rangeElement[1] {
			pre, node = pre.next, node.next
			continue
		}
		//same as 'Add'.Divide the range Element into two parts
		nextRangeElement[0] = node.rangeElement[1]
		nextRangeElement[1] = rangeElement[1]
		if rangeElement[0] > node.rangeElement[0] {
			// insert node1 if the left interval is not empty after node-rangeElement
			node1 := &rangeNode{[2]int{node.rangeElement[0], rangeElement[0]}, node.next}
			pre.next = node1
			pre = pre.next
		}
		if rangeElement[1] < node.rangeElement[1] {
			// insert node2 if the right interval is not empty after node-rangeElement
			node2 := &rangeNode{[2]int{rangeElement[1], node.rangeElement[1]}, node.next}
			pre.next = node2
			pre = pre.next
		}
		//delete node
		pre.next = node.next
		node = pre.next
		rangeElement = nextRangeElement
	}
	return nil
}

//Print TimeComplexity:O(n) SpaceComplexity:O(1)
func (rangeList *RangeLinkedList) Print() error {
	// traverse from head to tail.
	var node *rangeNode
	for node = rangeList.head.next; node != nil; node = node.next {
		fmt.Printf("[%d,%d) ", node.rangeElement[0], node.rangeElement[1])
	}
	fmt.Println()
	return nil
}

// NewRangeList returns a new RangeLinkedList object of type RangeList.
// TimeComplexity:O(1) SpaceComplexity:O(1)
func NewRangeList() RangeList {
	return &RangeLinkedList{
		head: &rangeNode{
			rangeElement: [2]int{-1, -1},
			next:         nil,
		},
	}
}
