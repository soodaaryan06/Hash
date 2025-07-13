package hashing

type Hasher interface {
	Hash(key string) int
}

type HashRing interface {
	AddNode(node *Node)
	Set(key, value string)
	Get(key string) (string, string, bool)
}
