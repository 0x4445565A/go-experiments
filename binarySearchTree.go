/**
 *
 * Implimenting Binary Search Tree with no documentation in "whiteboard" run.
 *
 * one compilaton error, mixed * and & on line 31
 *
 */
package main

import (
  "fmt"
)

type Tree struct {
  left  *Tree
  right *Tree
  value int
}

func (t *Tree) total() (total int) {
  total += t.value
  if t.right != nil {
    total += t.right.total()
  }
  if t.left != nil {
    total += t.left.total()
  }
  return total
}

func main() {
  t := Tree{value: 10}
  t.right = &Tree{value: 5}
  t.left = &Tree{value: 15}

  t2 := Tree{value: 7}
  t2.right = &Tree{value: 3}
  tLeft := *t.left
  tLeft.right = &t2

  fmt.Println(t.total())
}