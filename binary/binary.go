package binary

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

// Insert − Inserts an element in a tree/create a tree.
func (t *Tree) Insert(key int) {
	node := &Node{Key: key}
	if t.Root == nil {
		t.Root = node
		return
	}

	cursor := t.Root
	for {
		if key < cursor.Key {
			// if no child, insert the node as the cursor child
			if cursor.Left == nil {
				cursor.Left = node
				break
			}
			// go left
			cursor = cursor.Left
			continue
		}

		if cursor.Right == nil {
			// if no child, add the node as the cursor child
			cursor.Right = node
			break
		}
		// go right
		cursor = cursor.Right
	}
}

// Search searches a key in a tree.
func Search(root *Node, key int) bool {
	cursor := root
	for {
		if cursor == nil {
			return false
		}

		if key == cursor.Key {
			return true
		}

		// go left
		if key < cursor.Key {
			cursor = cursor.Left
			continue
		}

		// go right
		cursor = cursor.Right
	}
}

// Preorder Traversal − Traverses a tree in a pre-order manner.

// TraversalInOrder traverses a tree in an in-order manner.
func TraversalInOrder(root *Node, lvl int, cb func(data, lvl int)) {
	if root == nil {
		return
	}
	TraversalInOrder(root.Left, lvl+1, cb)
	cb(root.Key, lvl+1)
	TraversalInOrder(root.Right, lvl+1, cb)
}

// Postorder Traversal − Traverses a tree in a post-order manner
