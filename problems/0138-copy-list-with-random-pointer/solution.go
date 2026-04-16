package p0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func getClone(m *map[*Node]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}

	var clone, exists = (*m)[node]
	if !exists {
		clone = &Node{
			Val: node.Val,
		}
		(*m)[node] = clone
	}
	return clone
}

func copyRandomList(head *Node) *Node {
	var curr *Node = head
	m := make(map[*Node]*Node)
	for curr != nil {
		var clone = getClone(&m, curr)
		clone.Next = getClone(&m, curr.Next)
		clone.Random = getClone(&m, curr.Random)
		curr = curr.Next
	}

	return m[head]
}
