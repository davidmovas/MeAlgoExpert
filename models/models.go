package models

type LinkedList[T comparable] struct {
	Value T
	Next  *LinkedList[T]
}
