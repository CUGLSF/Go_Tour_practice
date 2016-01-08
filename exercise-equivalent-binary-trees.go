package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

type Sequence []int

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		go Walk(t.Left, ch)
		ch <- t.Value
		go Walk(t.Right, ch)
	}
}
// 实现sort.Interface的Len(), Less(i, j int) bool, 以及Swap(i, j int)
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	// chan_t1 and chan_t2 are both unbuffered channels
	chan_t1 := make(chan int)
	chan_t2 := make(chan int)
	s_t1 := make(Sequence, 0)
	s_t2 := make(Sequence, 0)
	go Walk(t1, chan_t1)
	go Walk(t2, chan_t2)
	for i := 0; i < 10; i++ {
		data_t1 := <-chan_t1
		data_t2 := <-chan_t2
		s_t1 = append(s_t1, data_t1)
		s_t2 = append(s_t2, data_t2)
	}
	sort.Sort(s_t1)
	//fmt.Println(s_t1)
	sort.Sort(s_t2)
	//fmt.Println(s_t2)
	for i := 0; i < 10; i++ {
		if s_t1[i] != s_t2[i] {
			return false
		}
	}
	return true
}

func main() {
	// c is a unbuffered channel
	c := make(chan int)
	// Test the Walk func
	go Walk(tree.New(1), c)
	for i := 0; i < 10; i++ {
		ii := <-c
		fmt.Println(ii)
	}
	// Test the Same func
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}