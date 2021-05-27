package main

import (
	"container/list"
	"sync"
)

type Stack struct {
	l *list.List
	m sync.Mutex
}

func NewStack() *Stack {
	return &Stack{l: list.New()}
}

func (s *Stack) Clear() {
	s.m.Lock()
	defer s.m.Unlock()
	s.l = nil
}

func (s *Stack) Size() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.l.Len()
}

func (s *Stack) isEmpty() bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.l.Len() == 0
}

func (s *Stack) Push(v interface{}) {
	if v == nil {
		return
	}
	s.m.Lock()
	defer s.m.Unlock()
	s.l.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	s.m.Lock()
	defer s.m.Unlock()
	if elem := s.l.Front(); elem != nil {
		s.l.Remove(elem)
		return elem.Value
	}
	return nil
}

func (s *Stack) Top() interface{} {
	s.m.Lock()
	defer s.m.Unlock()
	if elem := s.l.Front(); elem != nil {
		return elem.Value
	}
	return nil
}
