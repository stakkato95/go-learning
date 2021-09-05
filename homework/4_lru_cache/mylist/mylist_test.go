package mylist_test

import (
	"testing"

	"github.com/stakkato95/lru_cache/mylist"
	"github.com/stretchr/testify/assert"
)

func TestEmptyList(t *testing.T) {
	list := mylist.NewList()

	assert.Zero(t, list.Len())
	assert.Nil(t, list.Front())
	assert.Nil(t, list.Back())
	assert.False(t, list.Remove(nil))
	assert.False(t, list.MoveToFront(nil))
}

func TestSingleItemPush(t *testing.T) {
	t.Run("test PushFront", func(t *testing.T) {
		list := mylist.NewList()
		assert.NotNil(t, list.PushFront("a"))

		assert.Equal(t, list.Front(), list.Back())
		assert.Equal(t, 1, list.Len())

		assert.Nil(t, list.Front().Prev)
		assert.Nil(t, list.Front().Next)

		assert.Nil(t, list.Back().Next)
		assert.Nil(t, list.Back().Prev)
	})

	t.Run("test PushBack", func(t *testing.T) {
		list := mylist.NewList()
		assert.NotNil(t, list.PushBack("a"))

		assert.Equal(t, list.Front(), list.Back())
		assert.Equal(t, 1, list.Len())

		assert.Nil(t, list.Front().Prev)
		assert.Nil(t, list.Front().Next)

		assert.Nil(t, list.Back().Next)
		assert.Nil(t, list.Back().Prev)
	})
}

func TestMultiIyemPush(t *testing.T) {
	t.Run("test PushFront", func(t *testing.T) {
		list := mylist.NewList()
		assert.NotNil(t, list.PushFront("a"))
		assert.NotNil(t, list.PushFront("b"))
		assert.NotNil(t, list.PushFront("c"))
		assert.NotNil(t, list.PushFront("d"))

		assert.Equal(t, 4, list.Len())

		//check front
		assert.NotNil(t, list.Front())
		assert.Nil(t, list.Front().Prev)
		assert.NotNil(t, list.Front().Next)
		assert.Equal(t, "d", list.Front().Value)

		//check back
		assert.NotNil(t, list.Back())
		assert.NotNil(t, list.Back().Prev)
		assert.Nil(t, list.Back().Next)
		assert.Equal(t, "a", list.Back().Value)
	})

	t.Run("test PushBack", func(t *testing.T) {
		list := mylist.NewList()
		assert.NotNil(t, list.PushBack("a"))
		assert.NotNil(t, list.PushBack("b"))
		assert.NotNil(t, list.PushBack("c"))
		assert.NotNil(t, list.PushBack("d"))

		assert.Equal(t, 4, list.Len())

		//check front
		assert.NotNil(t, list.Front())
		assert.Nil(t, list.Front().Prev)
		assert.NotNil(t, list.Front().Next)
		assert.Equal(t, "a", list.Front().Value)

		//check back
		assert.NotNil(t, list.Back())
		assert.NotNil(t, list.Back().Prev)
		assert.Nil(t, list.Back().Next)
		assert.Equal(t, "d", list.Back().Value)
	})

	t.Run("test PushFront and PushBack", func(t *testing.T) {
		list := mylist.NewList()
		assert.NotNil(t, list.PushBack("c"))
		assert.NotNil(t, list.PushBack("d"))
		assert.NotNil(t, list.PushFront("b"))
		assert.NotNil(t, list.PushFront("a"))

		assert.Equal(t, 4, list.Len())

		//check front
		assert.NotNil(t, list.Front())
		assert.Nil(t, list.Front().Prev)
		assert.NotNil(t, list.Front().Next)
		assert.Equal(t, "a", list.Front().Value)

		//check 2nd
		second := list.Front().Next
		assert.NotNil(t, second)
		assert.Equal(t, list.Front(), second.Prev)
		assert.Equal(t, list.Front().Next, second)
		assert.Equal(t, "b", second.Value)

		//check 3rd
		third := second.Next
		assert.NotNil(t, third)
		assert.Equal(t, second, third.Prev)
		assert.Equal(t, second.Next, third)
		assert.Equal(t, "c", third.Value)

		//check back
		assert.NotNil(t, list.Back())
		assert.Equal(t, third, list.Back().Prev)
		assert.Nil(t, list.Back().Next)
		assert.Equal(t, "d", list.Back().Value)
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove first", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		newFront := list.Front().Next

		assert.True(t, list.Remove(list.Front()))
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, newFront, list.Front())
		assert.Equal(t, "b", list.Front().Value)
	})

	t.Run("remove last", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		newBack := list.Back().Prev

		assert.True(t, list.Remove(list.Back()))
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, newBack, list.Back())
		assert.Equal(t, "c", list.Back().Value)
	})

	t.Run("remove second", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		newSecond := list.Front().Next.Next

		assert.True(t, list.Remove(list.Front().Next))
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, newSecond, list.Front().Next)
		assert.Equal(t, "c", list.Front().Next.Value)
	})

	t.Run("remove second before last", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		newSecondLast := list.Back().Prev.Prev

		assert.True(t, list.Remove(list.Back().Prev))
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, newSecondLast, list.Back().Prev)
		assert.Equal(t, "b", list.Back().Prev.Value)
	})
}

func TestMoveToFrond(t *testing.T) {
	t.Run("move nil", func(t *testing.T) {
		list := mylist.NewList()
		assert.False(t, list.MoveToFront(nil))
	})

	t.Run("empty list", func(t *testing.T) {
		list := mylist.NewList()
		assert.False(t, list.MoveToFront(&mylist.Item{}))
	})

	t.Run("list with only one item", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		front := list.Front()

		assert.True(t, list.MoveToFront(list.Front()))
		assert.Equal(t, front, list.Front())
		assert.Equal(t, front, list.Back())
	})

	t.Run("move first to front", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		front := list.Front()
		back := list.Back()

		assert.True(t, list.MoveToFront(list.Front()))
		assert.Equal(t, front, list.Front())
		assert.Equal(t, back, list.Back())
		assert.Equal(t, 2, list.Len())
	})

	t.Run("move last to front in a list with 2 item", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		front := list.Front()
		back := list.Back()

		assert.True(t, list.MoveToFront(list.Back()))
		assert.Equal(t, back, list.Front())
		assert.Equal(t, front, list.Back())
		assert.Equal(t, 2, list.Len())
	})

	t.Run("move second item to front", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		assert.True(t, list.MoveToFront(list.Front().Next))
		assert.Equal(t, "b", list.Front().Value)
		assert.Equal(t, "a", list.Front().Next.Value)
		assert.Equal(t, 4, list.Len())
	})

	t.Run("move second last item to front", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")

		assert.True(t, list.MoveToFront(list.Back().Prev))
		assert.Equal(t, "c", list.Front().Value)
		assert.Equal(t, "a", list.Front().Next.Value)
		assert.Equal(t, "b", list.Front().Next.Next.Value)
		assert.Equal(t, "d", list.Back().Value)
		assert.Equal(t, 4, list.Len())
	})

	t.Run("move middle item to front", func(t *testing.T) {
		list := mylist.NewList()
		list.PushBack("a")
		list.PushBack("b")
		list.PushBack("c")
		list.PushBack("d")
		list.PushBack("e")

		assert.True(t, list.MoveToFront(list.Front().Next.Next))
		assert.Equal(t, "c", list.Front().Value)
		assert.Equal(t, "a", list.Front().Next.Value)
		assert.Equal(t, "b", list.Front().Next.Next.Value)
		assert.Equal(t, "d", list.Front().Next.Next.Next.Value)
		assert.Equal(t, "e", list.Back().Value)
		assert.Equal(t, 5, list.Len())
	})
}

func TestClear(t *testing.T) {
	list := mylist.NewList()
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("b")

	assert.Equal(t, 3, list.Len())

	list.Clear()

	assert.Zero(t, list.Len())
}
