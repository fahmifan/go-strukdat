package binary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(10)

	assert.Equal(t, 5, tree.Root.Key)
	assert.Equal(t, 4, tree.Root.Left.Key)
	assert.Equal(t, 6, tree.Root.Right.Key)
	assert.Equal(t, 2, tree.Root.Left.Left.Key)
	assert.Equal(t, 10, tree.Root.Right.Right.Key)
}

func TestTree_Search(t *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(10)

	assert.False(t, Search(tree.Root, 9))
	assert.False(t, Search(tree.Root, 1))
	assert.True(t, Search(tree.Root, 5))
	assert.True(t, Search(tree.Root, 4))
	assert.True(t, Search(tree.Root, 10))
}

func TestTree_Traversal_InOrder(t *testing.T) {
	tree := Tree{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(10)

	out := ""
	TraversalInOrder(tree.Root, 0, func(data, lvl int) {
		out += fmt.Sprintf("%d-", data)
	})

	assert.Equal(t, "2-4-5-7-10-", out)
}
