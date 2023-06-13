package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList_MoveToFront(t *testing.T) {
	t.Run("move element to front", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		item2 := l.PushBack(2)
		l.PushBack(3)
		l.MoveToFront(item2)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 2, l.Front().Value)
		require.Equal(t, 3, l.Back().Value)
	})

	t.Run("move head element", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		item := l.Front()
		l.MoveToFront(item)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 3, l.Front().Value)
		require.Equal(t, 1, l.Back().Value)
	})

	t.Run("move tail element", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		item := l.Back()
		l.MoveToFront(item)

		require.Equal(t, 3, l.Len())         // 3
		require.Equal(t, 1, l.Front().Value) // 1
		require.Equal(t, 2, l.Back().Value)  // 2
	})
}

func TestRemoveElement(t *testing.T) {
	t.Run("Remove single element", func(t *testing.T) {
		l := NewList()
		item := l.PushFront(1)
		l.Remove(item)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("remove tail element", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		item := l.Back()
		l.Remove(item)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 3, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("remove head element", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		item := l.Front()
		l.Remove(item)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 2, l.Front().Value)
		require.Equal(t, 1, l.Back().Value)
	})
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
