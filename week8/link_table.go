package week8

import "sync"

type LinkTable struct {
	sync.Mutex
	name string
	FirstNode *LinkNode
	LastNode  *LinkNode
	Count     int
}

func (link *LinkTable) String() string {
	iterator:= link.GetIterator()
	str := ""
	for iterator.HasNext() {
		if str == "" {
			str += "link "
			if link.name != "" {
				str += link.name
			}
			str += ": "
		}else {
			str += " -> "
		}
		str += iterator.Next()
	}
	return str
}

func NewLinkTable(str string) LinkTable {
	return LinkTable{name: str}
}

func (link *LinkTable) Add(str string) {
	link.Lock()
	defer link.Unlock()
	node := NewLinkNode(str)
	if link.LastNode != nil {
		link.LastNode.next = &node
	}
	link.LastNode = &node
	if link.FirstNode == nil {
		link.FirstNode = &node
	}
	link.Count = link.Count +1
}

func (link *LinkTable) GetIterator() LinkIterator {
	return NewLinkIterator(link)
}

type LinkNode struct {
	value string
	next *LinkNode
}

func NewLinkNode(value string) LinkNode {
	return LinkNode{value: value}
}

type LinkIterator struct {
	current LinkNode
	link    LinkTable
}

func NewLinkIterator(link *LinkTable) LinkIterator {
	return LinkIterator{link: *link}
}

func (iterator *LinkIterator) Next() string {
	if iterator.current == (LinkNode{}) {
		iterator.current = *(iterator.link.FirstNode)
		if iterator.current != (LinkNode{}) {
			return iterator.current.value
		}
	}
	value := (*iterator.current.next).value
	iterator.current = *iterator.current.next
	return value
}

func (iterator *LinkIterator) HasNext() bool {
	if iterator.current != (LinkNode{}) {
		return iterator.current.next != nil
	}else {
		return iterator.link.FirstNode != nil
	}
}

func (iterator *LinkIterator) CurrentItem() string {
	return iterator.current.value
}