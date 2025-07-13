package hashing

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

type ConsistentHasher struct{}

func (c ConsistentHasher) Hash(key string) int {
	// Intentionally cause collisions
	h := sha1.Sum([]byte(key))
	return int((uint32(h[0])<<24 | uint32(h[1])<<16 | uint32(h[2])<<8 | uint32(h[3])))
}

type ConsistentHash struct {
	Hasher  Hasher
	hashes  []int
	nodeMap map[int]*Node
}

func NewConsistentHash(hasher Hasher) *ConsistentHash {
	return &ConsistentHash{
		Hasher:  hasher,
		nodeMap: make(map[int]*Node),
	}
}

func (r *ConsistentHash) AddNode(node *Node) {
	h := r.Hasher.Hash(node.ID)
	r.hashes = append(r.hashes, h)
	r.nodeMap[h] = node
	sort.Ints(r.hashes)
}

func (r *ConsistentHash) getNode(key string) *Node {
	if len(r.hashes) == 0 {
		return nil
	}
	h := r.Hasher.Hash(key)
	idx := sort.Search(len(r.hashes), func(i int) bool {
		return r.hashes[i] >= h
	})
	if idx == len(r.hashes) {
		idx = 0
	}
	return r.nodeMap[r.hashes[idx]]
}

func (r *ConsistentHash) Set(key, value string) {
	node := r.getNode(key)
	if node == nil {
		fmt.Println("No node found")
		return
	}
	node.Data[key] = value
}

func (r *ConsistentHash) Get(key string) (string, string, bool) {
	node := r.getNode(key)
	if node == nil {
		return "", "", false
	}
	val, ok := node.Data[key]
	return val, node.ID, ok
}
