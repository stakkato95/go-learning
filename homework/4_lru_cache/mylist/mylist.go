package mylist

//generic and support types
type Value interface{}

type Item struct {
	Value Value
	Next  *Item
	Prev  *Item
}

type List interface {
	Len() int
	Front() *Item
	Back() *Item
	PushFront(v Value) *Item
	PushBack(v Value) *Item
	Remove(i *Item) bool
	MoveToFront(i *Item) bool
	Clear()
}

type list struct {
	length int
	first  *Item
	last   *Item
}

func NewList() List {
	return &list{}
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *Item {
	return l.first
}

func (l list) Back() *Item {
	return l.last
}

func (l *list) PushFront(v Value) *Item {
	if l.first == nil {
		l.initEpmtyList(v)
		return l.first
	}

	if l.length == 1 {
		l.first = &Item{Value: v, Next: l.last}
		l.last.Prev = l.first
		l.length++
		return l.first
	}

	oldFirst := l.first
	l.first = &Item{Value: v, Next: oldFirst}
	oldFirst.Prev = l.first
	l.length++
	return l.first
}

func (l *list) PushBack(v Value) *Item {
	if l.last == nil {
		l.initEpmtyList(v)
		return l.last
	}

	if l.length == 1 {
		l.last = &Item{Value: v, Prev: l.first}
		l.first.Next = l.last
		l.length++
		return l.last
	}

	oldLast := l.last
	l.last = &Item{Value: v, Prev: oldLast}
	oldLast.Next = l.last
	l.length++
	return l.last
}

func (l *list) Remove(i *Item) bool {
	if i == nil {
		return false
	}

	if l.length == 0 {
		return false
	}

	if i == l.first {
		newFirst := l.first.Next
		newFirst.Prev = nil
		l.first.Next = nil
		l.first = newFirst
		l.length--
		return true
	}
	if i == l.last {
		newLast := l.last.Prev
		newLast.Next = nil
		l.last.Prev = nil
		l.last = newLast
		l.length--
		return true
	}

	currentItem := l.first
	for {
		if currentItem == i {
			prev := i.Prev
			next := i.Next

			prev.Next = next
			next.Prev = prev

			i.Next = nil
			i.Prev = nil

			l.length--
			return true
		}

		if currentItem.Next == nil {
			return false
		}
		currentItem = currentItem.Next
	}
}

func (l *list) MoveToFront(i *Item) bool {
	if i == nil {
		return false
	}

	if l.length == 0 {
		return false
	}

	if l.length == 1 || i == l.first {
		return true
	}

	if l.length == 2 && i == l.last {
		newFirst := l.last
		newFirst.Prev = nil
		newFirst.Next = l.first

		l.first.Prev = l.last
		l.first.Next = nil

		l.last = l.first
		l.first = newFirst
		return true
	}

	if ok := l.Remove(i); !ok {
		return false
	}
	l.PushFront(i.Value)

	return true
}

func (l *list) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

func (l *list) initEpmtyList(v Value) {
	l.first = &Item{Value: v}
	l.last = l.first
	l.length++
}
