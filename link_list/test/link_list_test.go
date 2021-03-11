package test

import (
	"go_study/link_list"
	"testing"
)

func TestWeek8Job(t *testing.T) {
	tableA := link_list.NewLinkTable("A")
	tableA.Add("a")
	tableA.Add("b")
	tableA.Add("x")
	tableA.Add("y")
	tableA.Add("z")
	println(tableA.String())

	tableB := link_list.NewLinkTable("B")
	tableB.Add("d")
	tableB.Add("e")
	tableB.Add("f")
	tableB.Add("x")
	tableB.Add("y")
	tableB.Add("z")
	println(tableB.String())

	m := make(map[string]string, tableA.Count)

	iteratorA := tableA.GetIterator()
	//O(n)
	for iteratorA.HasNext() {
		item := iteratorA.Next()
		m[item] = item
	}

	var target string
	iteratorB := tableB.GetIterator()
	//O(m)
	for iteratorB.HasNext() {
		item := iteratorB.Next()
		if m[item] != ""  {
			target = item
			break
		}
	}

	if target == "" {
		println("链表A与链表B无合并")
	}else {
		println("链表A与链表B从元素[",target,"]开始合并")
	}
}