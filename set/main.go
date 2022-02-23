package main

type Seter interface {
	add(e int)
	delete(e int)
	exists(e int) bool
	getSize() int
}

type Node struct {
	value    int
	nextNode *Node
}

type Set struct {
	size      int
	firstNode *Node
}

func (s *Set) exists(e int) bool {
	currentNode := s.firstNode
	for {
		if currentNode.value == e {
			return true
		} else if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	return false
}

func (s Set) add(e int) {
	if !s.exists(e) {
		s.firstNode = &Node{
			value:    e,
			nextNode: s.firstNode,
		}
		s.size++
	}
}

func (s Set) delete(e int) {
	if s.firstNode.value == e {
		s.firstNode = s.firstNode.nextNode
		s.size--
		return
	}
	currentNode := s.firstNode

	for currentNode.nextNode != nil {
		if currentNode.nextNode.value == e {
			currentNode.nextNode = currentNode.nextNode.nextNode
			s.size--
			return
		}
		currentNode = currentNode.nextNode
	}
}

func (s Set) getSize() int {
	return s.size
}

var s Seter = &Set{}
