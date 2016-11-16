/**
 *  This was an experiment to see how well I would
 *  impliment a hash list on a flight with no internet.
 *
 *  If you can learn anything from this feel free to...
 *  copy, share, edit, distribute, etc.
 */

package main

import (
  "fmt"
)

/**
 * Keep HashTable in a struct so we can leverage methods.
 */
type HashTable struct {
  table map[int]*Node
}

/**
 * Node struct makes managing pointers hella easier.
 */
type Node struct {
  next *Node
  name string
  email string
}

/**
 * Go through and print every value in linked list.
 */
func (n *Node) print() {
  fmt.Println(" ", n.name, ":", n.email)
  if n.next != nil {
    n.next.print()
  }
}

/**
 * Insert new value into linked list.
 */
func (n *Node) insert(nameData string, emailData string) {
  if n.next == nil {
    n.next = &Node{name: nameData, email: emailData}
  } else {
    n.next.insert(nameData, emailData)
  }
}

/**
 * Find name in linked list.
 */
func (n *Node) find(data string) string {
  if data == n.name {
    return n.email
  } else {
    if n.next != nil {
      return n.next.find(data)
    } else {
      return ""
    }
  }
}

/**
 * Hash input for a list size of max 15.
 */
func (t *HashTable) hash(s string) (hash int) {
  for n := range s {
    hash += int(n)
  }
  return hash % 15
}

/**
 * Insert item into hash table, chain if needed.
 */
func (t *HashTable) insert(nameData string, emailData string) {
  hash := t.hash(nameData)
  if t.table[hash] == nil {
    t.table[hash] = &Node{name: nameData, email: emailData}
  } else {
    t.table[hash].insert(nameData, emailData)
  }
}

/**
 * Find value in hash table.
 */
func (t HashTable) find(data string) string {
  hash := t.hash(data)
  if t.table[hash] == nil {
    return ""
  } else {
    return t.table[hash].find(data)
  }
}

/**
 * Go through the hash table and linked lists to print all items.
 */
func (t HashTable) print() {
  for k, v := range t.table {
    fmt.Println(k, "{")
    v.print()
    fmt.Println("}")
  }
}

func main() {
  // Initialize a hash table
  hash_table := HashTable{table: make(map[int]*Node)}

  // Insert some records
  hash_table.insert("franky", "frank@cooldude.com")
  hash_table.insert("kevz", "leethax0r@gmail.com")
  hash_table.insert("meranda", "ravz@kewlkatz.com")
  hash_table.insert("bobby", "bobby.jones@mega64.com")
  hash_table.insert("matt", "matt.attackz@yahoo.com")
  hash_table.insert("brandon", "derp@ibreak.systems")

  // Dump all records
  hash_table.print()

  // Find some records
  fmt.Println("Finding meranda:", hash_table.find("meranda"))
  fmt.Println("Finding matt:", hash_table.find("matt"))
  fmt.Println("Finding kevz:", hash_table.find("kevz"))
}