package main

type LinkedLister interface {
	add(element int, position int)
	get(position int) int
	delete(position int)
	getSize() int
}

type Node struct {
	value    int
	nextNode *Node
}

type LinkedList struct {
	firstNode *Node
	size      int
}

func (l *LinkedList) add(element int, position int) {
	currentPosition := 0
	currentNode := l.firstNode

	if position == 0 {
		l.firstNode = &Node{
			value:    element,
			nextNode: currentNode.nextNode,
		}
		l.size++
		return
	}

	for currentPosition != position-1 {
		currentNode = currentNode.nextNode
		currentPosition++
	}

	currentNode.nextNode = &Node{
		value:    element,
		nextNode: currentNode.nextNode,
	}
}

func (l *LinkedList) get(position int) int {
	currentPosition := 0
	currentNode := l.firstNode

	for position != currentPosition {
		currentNode = currentNode.nextNode
		currentPosition++
	}

	return currentNode.value
}

func (l *LinkedList) delete(position int) {
	currentPosition := 0
	currentNode := l.firstNode

	if position == 0 {
		l.firstNode = l.firstNode.nextNode
		l.size--
		return
	}

	for currentPosition != position-1 {
		currentNode = currentNode.nextNode
		currentPosition++
	}

	currentNode.nextNode = currentNode.nextNode.nextNode
	l.size--
	return
}

func (l *LinkedList) getSize() int {
	return l.size
}
