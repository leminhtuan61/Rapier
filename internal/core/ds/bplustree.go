package ds

type Item struct {
	Score  float64
	Member string
}

func (i *Item) CompareTo(other *Item) int {
	if i.Score < other.Score {
		return -1
	}
	if i.Score > other.Score {
		return 1
	}
	if i.Score > other.Score {
		return 1
	}
	if i.Member < other.Member {
		return -1
	}
	if i.Member > other.Member {
		return 1
	}
	return 0
}

type Node struct {
	Items    []*Item
	Children []*Node
	IsLeaf   bool
	Parent   *Node
	Next     *Node
}
type BPlusTree struct {
	Root   *Node
	Degree int
}

func NewBPlusTree(degree int) *BPlusTree {
	return &BPlusTree{
		Root:   &Node{IsLeaf: true},
		Degree: degree,
	}
}

func (t *BPlusTree) Score(member string) (float64, bool) {
	node := t.Root
	for !node.IsLeaf {
		node = node.Children[0]
	}
	for node != nil {
		for _, item := range node.Items {
			if item.Member == member {
				return item.Score, true
			}
		}
		node = node.Next
	}
	return 0, false
}
func (t *BPlusTree) Add(score float64, member string) int {
	item := &Item{Score: score, Member: member}
	node := t.Root
	for !node.IsLeaf {
		i := 0
		for i < len(node.Items) && score >= node.Items[i].Score {
			i++
		}
		node = node.Children[i]
	}

	for i, existingItem := range node.Items {
		if existingItem.Member == member {
			node.Items[i].Score = score
			return 1
		}
	}
	i := 0
	for i < len(node.Items) && score > node.Items[i].Score {
		i++
	}
	node.Items = append(node.Items, nil)
	copy(node.Items[i+1:], node.Items[i:])
	node.Items[i] = item
	if len(node.Items) > t.Degree-1 {
		t.splitNode(node)
	}
	return 1
}

func (t *BPlusTree) splitNode(node *Node) {
	// If the node is the root, we need to create a new root.
	if node.Parent == nil {
		t.splitRoot()
		return
	}

	// Split based on whether the node is a leaf or an internal node.
	if node.IsLeaf {
		t.splitLeaf(node)
	} else {
		t.splitInternal(node)
	}
}
func (t *BPlusTree) splitInternal(node *Node) {
	medianIndex := len(node.Items) / 2
	newInternal := &Node{IsLeaf: false, Parent: node.Parent}

	promotedItem := node.Items[medianIndex]
	newInternal.Items = append(newInternal.Items, node.Items[medianIndex+1:]...)
	newInternal.Children = append(newInternal.Children, node.Children[medianIndex+1:]...)
	node.Items = node.Items[:medianIndex]
	node.Children = node.Children[:medianIndex+1]

	for _, child := range newInternal.Children {
		child.Parent = newInternal
	}
	// Insert the promoted item and new child into the parent
	parent := node.Parent
	childIndex := 0
	for childIndex < len(parent.Children) && parent.Children[childIndex] != node {
		childIndex++
	}
	// Insert the promoted key and the new child.
	parent.Items = append(parent.Items[:childIndex], append([]*Item{promotedItem}, parent.Items[childIndex:]...)...)
	parent.Children = append(parent.Children[:childIndex+1], append([]*Node{newInternal}, parent.Children[childIndex+1:]...)...)

	// If the parent now overflows, split it too.
	if len(parent.Items) > t.Degree-1 {
		t.splitNode(parent)
	}

}
func (t *BPlusTree) splitLeaf(node *Node) {
	medianIndex := len(node.Items) / 2
	newLeaf := &Node{IsLeaf: true, Parent: node.Parent, Next: node.Next}
	newLeaf.Items = append(newLeaf.Items, node.Items[medianIndex:]...)
	node.Items = node.Items[:medianIndex]
	node.Next = newLeaf
	// Promote the first item of the new leaf to the parent
	parent := node.Parent
	promotedItem := newLeaf.Items[0]
	childIndex := 0
	for childIndex < len(parent.Children) && parent.Children[childIndex] != node {
		childIndex++
	}

	// Insert the promoted key and the new child node into the parent.
	parent.Items = append(parent.Items[:childIndex], append([]*Item{promotedItem}, parent.Items[childIndex:]...)...)
	parent.Children = append(parent.Children[:childIndex+1], append([]*Node{newLeaf}, parent.Children[childIndex+1:]...)...)

	// If the parent now overflows, split it too.
	if len(parent.Items) > t.Degree-1 {
		t.splitNode(parent)
	}
}
func (t *BPlusTree) splitRoot() {
	oldRoot := t.Root
	newRoot := &Node{IsLeaf: false}

	t.Root = newRoot
	oldRoot.Parent = newRoot
	newRoot.Children = append(newRoot.Children, oldRoot)
	if oldRoot.IsLeaf {
		t.splitLeaf(oldRoot)
	} else {
		t.splitInternal(oldRoot)
	}
}
func (t *BPlusTree) GetRank(member string) int {
	node := t.Root
	for !node.IsLeaf {
		node = node.Children[0] //always go to the leftmost child
	}
	rank := 0
	// Traverse the leaf nodes to find the member and calculate its rank
	for node != nil {
		for _, item := range node.Items {
			// Check if we have found the member
			if item.Member == member {
				return rank
			}
			rank++
		}
		node = node.Next
	}
	return -1 // Member not found
}
