package main

import "container/list"

type pair struct {
	k, v int
}

type MyHashMap struct {
	data []list.List
}

const base int = 769

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

func hash(key int) int {
	return key % base
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	h := hash(key)

	//find if this pair already exists
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if ev := e.Value.(pair); ev.k == key {
			e.Value = pair{key, value}
			return
		}
	}
	this.data[h].PushBack(pair{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := hash(key)

	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if ev := e.Value.(pair); ev.k == key {
			return ev.v
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if ev := e.Value.(pair); ev.k == key {
			this.data[h].Remove(e)
			return
		}
	}
	return

}
func main() {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
