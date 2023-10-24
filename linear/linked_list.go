package linear

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedist *LinkedList) AddToHead(property int) {
	node := Node{}
	node.property = property

	if node.nextNode != nil {
		node.nextNode = linkedist.headNode
	}

	linkedist.headNode = &node
}
