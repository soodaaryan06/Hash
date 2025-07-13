package hashing

type Node struct {
	ID   string
	Data map[string]string
}

func NewNode(id string) *Node {
	return &Node{
		ID:   id,
		Data: make(map[string]string),
	}
}
