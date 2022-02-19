package main

type Stacker interface {
	add(element int)
	delete()
	getFirst() int
	getSize() int
}

type Node struct {
	value    int
	nextNode *Node
}

type Stack struct {
	firstNode *Node
	size      int
}

func (s *Stack) add(element int) {
	s.firstNode = &Node{
		value:    element,
		nextNode: s.firstNode,
	}
	s.size++
}

func (s *Stack) delete(position int) {
	s.firstNode = s.firstNode.nextNode
	s.size--
}

func (s *Stack) getFirst() int {
	return s.firstNode.value
}

func (s *Stack) getSize() int {
	return s.size
}
