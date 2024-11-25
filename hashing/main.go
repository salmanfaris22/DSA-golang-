package main

import "fmt"

const ArraySize = 7

// HashTable will hold an arary
type HashTabler struct {
	array [ArraySize]*bucket
}

// buket linked list
type bucket struct {
	head *bucketNode
}

// bucker not struc
type bucketNode struct {
	key  string
	next *bucketNode
}

func initHashTabler() *HashTabler {
	result := &HashTabler{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func (h *HashTabler) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// func (h *HashTabler) Search(key string) bool {
// 	index := hash(key)
// }

// func (h *HashTabler) Delete(key string) {
// 	index := hash(key)
// }

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func main() {
	// testHasgTable := initHashTabler()
	// fmt.Println(testHasgTable)
	// testHasgTable.Insert("randy")
	testBuket := &bucket{}
	testBuket.insert("salman")
	testBuket.insert("salman")
	fmt.Println(testBuket.search("salman"))
	fmt.Println(testBuket.search("fa"))
}

// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exits")
	}

}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	preveNode := b.head
	for preveNode != nil {
		if preveNode.next.key == k {
			preveNode.next = preveNode.next.next
		}
		preveNode = preveNode.next
	}

}
