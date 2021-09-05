package stack_test

import (
	"testing"

	"github.com/stakkato95/example_int_stack/stack"
	"github.com/stretchr/testify/assert"
)

func TestPushWhenCapasityIsZero(t *testing.T) {
	s := stack.NewStack(0)
	err := s.Push(1)
	assert.NotNil(t, err)
}

func TestPopWhenEmptyStask(t *testing.T) {
	s := stack.NewStack(10)

	item, ok := s.Pop()

	assert.Zero(t, item)
	assert.False(t, ok)
}

func TestPeekWhenEmptyStask(t *testing.T) {
	s := stack.NewStack(10)

	item, ok := s.Peek()

	assert.Zero(t, item)
	assert.False(t, ok)
}

func TestPush(t *testing.T) {
	data := testData()
	s := stack.NewStack(len(data))

	for _, item := range data {
		err := s.Push(item)
		assert.Nil(t, err)
	}

	assert.Equal(t, len(data), s.Size())
}

func TestPushAndNormalPop(t *testing.T) {
	data := testData()
	dataLen := len(data)
	s := stack.NewStack(dataLen)

	for _, item := range data {
		s.Push(item)
	}

	assert.Equal(t, dataLen, s.Size())

	for i := 0; i < dataLen; i++ {
		item, ok := s.Pop()
		assert.Equal(t, data[dataLen-i-1], item)
		assert.True(t, ok)
	}

	assert.Zero(t, s.Size())

	item, ok := s.Pop()
	assert.Zero(t, item)
	assert.False(t, ok)

	item, ok = s.Peek()
	assert.Zero(t, item)
	assert.False(t, ok)
}

func TestPushAndRandomPop(t *testing.T) {
	data := testData()
	dataLen := len(data)
	s := stack.NewStack(dataLen)

	t.Run("push 1, remove 1", func(t *testing.T) {
		assert.Nil(t, s.Push(data[0]))
		item, ok := s.Peek()
		assert.True(t, ok)
		assert.Equal(t, data[0], item)

		item, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, data[0], item)

		item, ok = s.Peek()
		assert.False(t, ok)
		assert.Zero(t, item)
	})

	t.Run("push 3, remove 1, push 1, remove 3", func(t *testing.T) {
		//push 3
		assert.Nil(t, s.Push(data[0]))
		assert.Nil(t, s.Push(data[1]))
		assert.Nil(t, s.Push(data[2]))

		//remove 1
		item, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, data[2], item)

		//push 1
		newItem := 100500
		assert.Nil(t, s.Push(newItem))
		item, ok = s.Peek()
		assert.True(t, ok)
		assert.Equal(t, newItem, item)

		//remove 3
		item, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, newItem, item)

		item, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, data[1], item)

		item, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, data[0], item)
	})

	t.Run("empty stack", func(t *testing.T) {
		item, ok := s.Pop()
		assert.False(t, ok)
		assert.Zero(t, item)

		item, ok = s.Peek()
		assert.False(t, ok)
		assert.Zero(t, item)

		assert.Zero(t, s.Size())
	})
}

func TestErrorOnPushWhenMaxCapcityReached(t *testing.T) {
	s := stack.NewStack(3)

	for i := 0; i < 3; i++ {
		assert.Nil(t, s.Push(i))
	}

	assert.NotNil(t, s.Push(1))
}

func testData() []int {
	return []int{1, 3, 5, 7, 9}
}
