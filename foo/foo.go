package main

import (
	"fmt"
)

// Definition for singly-linked list representing a number
type ListNode struct {
	Val  int
	Next *ListNode
}

func getNumber(l *ListNode) (n int) {
	if l == nil {
		return
	}
	n = l.Val + getNumber(l.Next)*10
	return
}

func putNumber(n int) (l *ListNode) {
	l = &ListNode{n % 10, nil}
	if r := n / 10; r != 0 {
		l.Next = putNumber(r)
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return putNumber(getNumber(l1) + getNumber(l2))
}

func twoSum(nums []int, target int) (res []int) {
	for x := 0; x < len(nums); x++ {
		for i := x + 1; i < len(nums); i++ {
			if nums[x]+nums[i] == target {
				res = []int{x, i}
				return
			}
		}
	}
	return
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(getNumber(l1))
	fmt.Println(getNumber(l2))
	ls := addTwoNumbers(l1, l2)
	fmt.Println(getNumber(ls))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 22))
}
