# utils
一个常用的go工具包，使用了go泛型，需要go1.18以上

# 使用方法

```go
package main

import (
    "fmt"

    "github.com/lewjian/utils/set"
)

func main() {
    s := set.NewSet[int]()
    s.Add(1, 2, 3, 1, 2, 3)
    fmt.Printf("%v", s.Values())
}
```

# roadmap
- set 
- array
- todo

# go doc

## set
```go
TYPES

type Set[T comparable] struct {
        // Has unexported fields.
}

func NewSet[T comparable]() *Set[T]
    NewSet init Set

func (s *Set[T]) Add(key ...T)
    Add a new key

func (s *Set[T]) Del(key T)
    Del key

func (s *Set[T]) Has(key T) bool
    Has returns has key

func (s *Set[T]) Values() []T
    Values returns unique keys
```
## array
```go
FUNCTIONS

func Diff[T comparable](src []T, others ...[]T) []T
    Diff returns elements in src, but not in others...

func Filter[T comparable](data []T, f func(i int) bool) []T
    Filter a slice use function f, element dropped when f get false

func InArray[T comparable](item T, data []T) (exist bool, index int)
    InArray check if data contains element item, index=-1 when not exists

func Intersect[T comparable](arrs ...[]T) []T
    Intersect return a slice contains values all arrs has

func Merge[T comparable](base []T, others ...[]T) []T
    Merge all slice into one

func Range[T comparable](data []T, f func(i int) (breakHere bool))
    Range same as for range, break when f got true

func Remove[T comparable](data []T, index int) []T
    Remove index element of data

func RemoveAll[T comparable](data []T, item T) []T
    RemoveAll remove all element item from data slice

func RemoveElement[T comparable](data []T, item T, count int) []T
    RemoveElement remove element of data, only remove first count elements if
    count < 0, remove all elements 删除切片中出现的前count次元素，如果count小于0，则删除所有切片中的element

func Reverse[T comparable](data []T)
    Reverse slice on the input data slice

func ReverseNew[T comparable](data []T) []T
    ReverseNew returns a new reverse slice

```
