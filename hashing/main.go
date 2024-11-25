package main

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

func initHashTabler() {
	hashTable := &HashTabler{}
	for i := range hashTable.array {
		hashTable.array[i] =
	}
}

func main() {

}
