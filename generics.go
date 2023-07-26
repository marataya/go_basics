package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func foo[T, U any](val1 T, val2 U) {
	fmt.Println(val1, val2)
}

func main() {
	cm := NewCustomMap[string, int]()
	cm.Insert("abc", 5)
	cm.Insert("d", 7)
	fmt.Println(*cm)
	cm1 := NewCustomMap[int, float64]()
	cm1.Insert(1, 9.99)
	cm1.Insert(2, 10.1)
	fmt.Println(*cm1)
	foo(1, 45.1)
	foo(2.5, 45.1)
	foo("asdf", 45.0)
}
