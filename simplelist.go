package simplelist

import (
	"sync"
)

// The smallest variable we can have - an empty struct
type emptyStruct struct{}

type SimpleList struct {
	items map[interface{}]emptyStruct
	sync.Mutex
}

// New returns an initialised SimpleList
func New() *SimpleList {
	return &SimpleList{
		items: make(map[interface{}]emptyStruct),
	}
}

// Add an item to the list
func (v *SimpleList) Add(item interface{}) {
	v.Lock()
	v.items[item] = emptyStruct{}
	v.Unlock()
}

// Remove an item from the list
func (v *SimpleList) Remove(item interface{}) {
	v.Lock()
	if _, ok := v.items[item]; ok {
		delete(v.items, item)
	}
	v.Unlock()
}

// Set whether an item is the list or not
func (v *SimpleList) Set(item interface{}, inlist bool) {
	if inlist {
		v.Add(item)
	} else {
		v.Remove(item)
	}
}

// Exists returns true if the item is in the list
func (v *SimpleList) Exists(item interface{}) bool {
	v.Lock()
	_, ok := v.items[item]
	v.Unlock()
	return ok
}
