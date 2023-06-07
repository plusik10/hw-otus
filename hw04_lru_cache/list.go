package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  l.head,
		Prev:  nil,
	}

	if l.head != nil {
		l.head.Prev = newItem
	} else {
		l.tail = newItem
	}

	l.head = newItem
	l.length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.tail,
	}
	if l.tail != nil {
		l.tail.Next = newItem
	} else {
		l.head = newItem
	}
	l.tail = newItem
	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {

	switch {
	case l.length == 0:
		return
	case l.length == 1:
		l.head = nil
		l.tail = nil
		l.length--
		return
	case l.head == i:
		l.head = l.head.Next
		l.head.Prev = nil
		l.length--
		return
	case l.tail == i:
		l.tail = l.tail.Prev
		l.tail.Next = nil
		l.length--
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i)
	l.Remove(i)
}

func NewList() List {
	return new(list)
}
