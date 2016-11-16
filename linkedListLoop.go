/**
 *
 *
 */
 package main

 import (
   "fmt"
)

type Node struct {
  next *Node
  value string
}

/**
 *  Abuse nil struct types and maps to check for loops in linked list.
 */
func (n *Node) loops(checked map[*Node]*Node) bool {
  if checked[n] != nil {
    return true
  }
  if n.next == nil {
    return false
  }
  checked[n] = n
  return n.next.loops(checked)
}

func (n *Node) insert(data string) {
  if n.next == nil {
    n.next = &Node{value: data}
  } else {
    n.next.insert(data)
  }
}

func (n Node) unSafePrint() {
  fmt.Println(n.value)
  if n.next != nil {
    n.next.unSafePrint()
  }
}

func (n Node) safePrint() {
  if n.loops(make(map[*Node]*Node)) {
    fmt.Println("Inf loop detected")
  } else {
    n.unSafePrint()
  }
}

func (n *Node) last() *Node {
  if n.next == nil {
    return n
  } else {
    return n.next.last()
  }
}

func main() {
  node := Node{value: "andrew"}
  node.insert("bob")
  node.insert("carl")
  node.insert("derek")
  // Grab the last item's pointer and store it for later
  badActionDarek := node.last()
  node.insert("earl")
  node.insert("frank")
  node.insert("gary")
  node.insert("hank")
  // point last item's next pointer to the previously grabbed ptr.
  badAction := node.last()
  badAction.next = badActionDarek
  // Loop forever buahaha
  //node.unSafePrint()

  // Using safe loop instead
  node.safePrint()
  // Removing unsafe item
  badAction.next = nil
  // Printing
  node.safePrint()
}