package main

import (
	"container/list"
)

const base = 769

type MyHashSet struct {
	data []list.List
}

// define a hash function
func (s *MyHashSet) hash(key int) int {
	return key % base
}
func Constructor() MyHashSet {
	return MyHashSet{data: make([]list.List, base)}

}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(key)
	}
}
func (this *MyHashSet) Remove(key int) {

	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[h].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

func main() {

}
