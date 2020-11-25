package list_test

import (
	"testing"

	"github.com/lqqyt2423/algo-go/pkg/list"
	"github.com/stretchr/testify/assert"
)

func assertNilList(t *testing.T, dl *list.DoubleList) {
	assert.Equal(t, []int{}, dl.Inspect())
	assert.Equal(t, 0, dl.Size())
	assert.Nil(t, dl.Head())
	assert.Nil(t, dl.Tail())
}

func TestDoubleList_AddFirst(t *testing.T) {
	dl := list.NewDoubleList()
	assertNilList(t, dl)
	dl.AddFirst(&list.Node{Val: 0})
	assert.Equal(t, []int{0}, dl.Inspect())
	assert.Equal(t, 1, dl.Size())
	dl.AddFirst(&list.Node{Val: 2})
	assert.Equal(t, []int{2, 0}, dl.Inspect())
	assert.Equal(t, 2, dl.Size())
}

func TestDoubleList_AddLast(t *testing.T) {
	dl := list.NewDoubleList()
	assertNilList(t, dl)
	dl.AddLast(&list.Node{Val: 0})
	assert.Equal(t, []int{0}, dl.Inspect())
	assert.Equal(t, 1, dl.Size())
	dl.AddLast(&list.Node{Val: 2})
	assert.Equal(t, []int{0, 2}, dl.Inspect())
	assert.Equal(t, 2, dl.Size())
}

func TestDoubleList_RemoveFirst(t *testing.T) {
	dl := list.NewDoubleList()
	assertNilList(t, dl)
	dl.RemoveFirst()
	assert.Equal(t, []int{}, dl.Inspect())
	dl.AddLast(&list.Node{Val: 1})
	assert.Equal(t, []int{1}, dl.Inspect())
	dl.RemoveFirst()
	assertNilList(t, dl)
	dl.AddFirst(&list.Node{Val: 1})
	dl.AddFirst(&list.Node{Val: 2})
	dl.RemoveFirst()
	assert.Equal(t, []int{1}, dl.Inspect())
	dl.RemoveFirst()
	assert.Equal(t, []int{}, dl.Inspect())
}

func TestDoubleList_RemoveLast(t *testing.T) {
	dl := list.NewDoubleList()
	assertNilList(t, dl)
	dl.RemoveLast()
	assert.Equal(t, []int{}, dl.Inspect())
	dl.AddFirst(&list.Node{Val: 1})
	assert.Equal(t, []int{1}, dl.Inspect())
	dl.RemoveLast()
	assertNilList(t, dl)
	dl.AddFirst(&list.Node{Val: 1})
	dl.AddFirst(&list.Node{Val: 2})
	dl.RemoveLast()
	assert.Equal(t, []int{2}, dl.Inspect())
	dl.RemoveLast()
	assert.Equal(t, []int{}, dl.Inspect())
}

func TestDoubleList_Remove(t *testing.T) {
	dl := list.NewDoubleList()
	assertNilList(t, dl)
	dl.Remove(nil)
	assertNilList(t, dl)
	n1 := &list.Node{Val: 1}
	dl.AddLast(n1)
	dl.Remove(n1)
	assertNilList(t, dl)
	dl.AddLast(&list.Node{Val: 1})
	n2 := &list.Node{Val: 2}
	dl.AddLast(n2)
	dl.AddLast(&list.Node{Val: 3})
	assert.Equal(t, []int{1, 2, 3}, dl.Inspect())
	dl.Remove(n2)
	assert.Equal(t, []int{1, 3}, dl.Inspect())
	assert.Equal(t, 2, dl.Size())
}
