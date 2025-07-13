package main

import (
	"awesomego/hashing"
	"fmt"
)

func main() {
	var ring hashing.HashRing
	ring = hashing.NewConsistentHash(hashing.ConsistentHasher{}) // 🔁 swap to test bad hash

	ring.AddNode(hashing.NewNode("NodeA"))
	//ring.AddNode(hashing.NewNode("NodeB"))
	ring.AddNode(hashing.NewNode("NodeC"))

	ring.Set("apple", "red")
	ring.Set("banana", "yellow")
	ring.Set("cherry", "dark red")

	keys := []string{"apple", "banana", "cherry"}
	for _, k := range keys {
		val, node_id, ok := ring.Get(k)
		if ok {
			fmt.Printf("%s => %s and node => %s\n", k, val, node_id)
		} else {
			fmt.Printf("%s => not found\n", k)
		}
	}
}
