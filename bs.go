package main

import (
  "fmt"
  "sort"
)

func bs(a []int, s int) int {
  sort.Ints(a)
  low := 0
  high := len(a) - 1
  i := 0
  if s > a[high] || s < a[low] {
    return -1
  }
  for low < high {
    i = ((high - low) / 2) + low
    if a[i] == s {
      return i
    }
    if a[i] < s {
      low = i + 1
    } else {
      high = i - 1
    }
  }
  return -1
}

func main() {
  fmt.Println(bs([]int{1, 5, 6, 7, 8, 12, 22, 42, 36}, 8))
  fmt.Println(bs([]int{1, 5, 6, 7, 8, 12, 22, 42, 36}, 5))
  fmt.Println(bs([]int{1, 5, 6, 7, 8, 12, 22, 42, 36}, 9))
  fmt.Println(bs([]int{1, 5, 6, 7, 8, 12, 22, 42, 36}, 0))
  fmt.Println(bs([]int{1, 5, 6, 7, 8, 12, 22, 42, 36}, 46))
}
