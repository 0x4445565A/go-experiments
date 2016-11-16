package main

import (
  "fmt"
)

// FILO
type Stack struct {
  value int
  next  *Stack
}

func (s *Stack) add(i int) {
  if s.next == nil {
    s.next = &Stack{value: i}
  } else {
    s.next.add(i)
  }
}

func (s *Stack) pop() bool {
  if s.next == nil {
    return true
  } else {
    pop := s.next.pop()
    if pop {
      s.next = nil
    }
  }
  return false
}

func (s *Stack) print() {
  fmt.Println(s.value)
  if s.next != nil {
    s.next.print()
  }
}

// FIFO
type Queue struct {
  value int
  next  *Queue
}

func (q *Queue) add(i int) {
  if q.next == nil {
    q.next = &Queue{value: i}
  } else {
    q.next.add(i)
  }
}


// Is this hacky?  I don't know anymore.
// Could just pass by reference because the list is by reference.
func (q *Queue) pop() {
  q.value = q.next.value
  q.next = q.next.next
}

func (q *Queue) print() {
  fmt.Println(q.value)
  if q.next != nil {
    q.next.print()
  }
}

func main() {
  q := Queue{value: 1}
  q.add(2)
  q.add(3)
  q.add(4)
  q.print()
  q.pop()
  q.pop()
  q.print()
  fmt.Println("Stacks on stacks")
  s := Stack{value: 1}
  s.add(2)
  s.add(3)
  s.add(4)
  s.add(5)
  s.print()
  s.pop()
  s.pop()
  s.print()
}