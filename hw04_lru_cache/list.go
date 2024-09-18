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
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFront := &ListItem{
		Value: v,
	}

	if l.front == nil {
		l.back = newFront
	} else {
		newFront.Next = l.front
		l.front.Prev = newFront
	}

	l.front = newFront
	l.len++

	return newFront
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := &ListItem{
		Value: v,
	}

	if l.back == nil {
		l.front = newBack
	} else {
		newBack.Prev = l.back
		l.back.Next = newBack
	}

	l.back = newBack
	l.len++

	return newBack
}

func (l *list) Remove(i *ListItem) {
	next := i.Next
	prev := i.Prev
	switch {
	case next == nil && prev == nil: // Если елемент один
		l.front = nil
		l.back = nil
	case next == nil:
		l.back = l.back.Prev
		l.back.Next = nil
	case prev == nil:
		l.front = l.front.Next
		l.front.Prev = nil
	default:
		i.Next.Prev = prev
		i.Prev.Next = next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
