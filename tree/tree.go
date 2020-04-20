package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(n, k int) *Tree {
	var insert func(*Tree, int) *Tree
	insert = func(t *Tree, v int) *Tree {
		if t == nil {
			return &Tree{nil, v, nil}
		}
		if v < t.Value {
			t.Left = insert(t.Left, v)
		} else {
			t.Right = insert(t.Right, v)
		}
		return t
	}
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// New0 returns a new, one-left-leaf binary tree holding the values k, 2k, ..., 10k.
func New0(n, k int) *Tree {
	var insert func(*Tree, int) *Tree
	insert = func(t *Tree, v int) *Tree {
		if t == nil {
			return &Tree{nil, v, nil}
		}
		if v < t.Value {
			t.Left = insert(t.Left, v)
		} else {
			return &Tree{t, v, nil}
		}
		return t
	}
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	var walk func(*Tree, chan int)
	walk = func(t *Tree, ch chan int) {
		if t == nil {
			return
		}
		if t.Left != nil {
			walk(t.Left, ch)
		} else if t.Right != nil {
			walk(t.Right, ch)
		}
		ch <- t.Value
	}
	walk(t, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *Tree, n int) bool {
	ch1 := make(chan int, n)
	Walk(t1, ch1)
	ch2 := make(chan int, n)
	Walk(t2, ch2)
	for i := 0; i < n; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	n := 16
	k := 10
	tr1 := New0(n, k)
	fmt.Println(tr1)
	tr2 := New(n, k)
	fmt.Println(tr2)
	if Same(tr1, tr1, n) {
		fmt.Println("Trees are same")
	} else {
		fmt.Println("Trees differ")
	}
}
