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
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
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

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int, toClose bool) {
	if t == nil {
		return;
	}



	Walk(t.Left, ch, false)
	ch <- t.Value
	Walk(t.Right, ch, false)

	if toClose {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)

	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		fmt.Println(v1, ok1, v2, ok2)
		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else {
			return !(ok1 || ok2)
		}
	}
}

func main() {
	fmt.Println(Same(New(1), New(2)))
}
