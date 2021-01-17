package merkledb

import "fmt"

// EmptyNode is a Node implementation that represents non existing Nodes
// it's used mainly when traversing the tree fetching a non existing key
// it pre-sets the conditions for insertion
type EmptyNode struct {
	parent Node
	key    Key
}

// NewEmptyNode returns a new EmptyNode
// the Node will have
//    Parent        -> typically the node that instantiates it
//    key           -> the address that is non existent in the tree
func NewEmptyNode(parent Node, key Key) Node {
	return &EmptyNode{
		parent: parent,
		key:    key,
	}
}

// GetChild should never be reached
func (e *EmptyNode) GetChild(key Key) (Node, error) { return nil, nil }

// GetNextNode returns itself
func (e *EmptyNode) GetNextNode(prefix Key, start Key, key Key) (Node, error) {
	return nil, nil
}

// Insert requests it's Parent to insert the k/v
func (e *EmptyNode) Insert(key Key, value []byte) error {
	return e.parent.Insert(key, value)
}

// Delete should never be called
func (e *EmptyNode) Delete(key Key) error { return nil }

// SetChild should never be called
func (e *EmptyNode) SetChild(node Node) error { return nil }

// SetParent should never be called
func (e *EmptyNode) SetParent(node Node) {}

// SetPersistence should never be called
func (e *EmptyNode) SetPersistence(p *Persistence) {}

// Value should never be called
func (e *EmptyNode) Value() []byte { return nil }

func (e *EmptyNode) Hash(key Key, hash []byte) error { return nil }

func (e *EmptyNode) GetHash() []byte { return nil }

func (e *EmptyNode) GetPreviousHash() []byte { return nil }

// Key holds the key of the to-be-inserted node
func (e *EmptyNode) Key() Key {
	return e.key
}

// Print should never be called
func (e *EmptyNode) Print() {
	fmt.Printf("ERROR: should never be called EmptyNode ID: %p - Parent: %p", e, e.parent)
}
