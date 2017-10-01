// Package trie contains a primitive implementation of the Trie data structure.
//
// Copyright 2017 Aleksandr Bezobchuk.
package trie

import (
	"sync"
)

// trieNode reflects a node that a trie is composed of. Each node can have a
// single child node, child. that acts as a pointer to a singly linked list to
// other potential children which are accessed via next. This allows for
// memory-efficiency in that we do not need to store the entire key alphabet
// as children for each node. A node may also have a value meaning it can be a
// terminating/search node.

// trieNode implements a node that the Trie is composed of. Each node contains
// a symbol that a key can be composed of unless the node is the root. The node
// has a collection of children that is represented as a hashmap. This allows
// for time and space complexity efficiency. The node may also contain a value
// that indicates a possible query result.
type trieNode struct {
	children map[byte]*trieNode
	symbol   byte
	value    []byte
	root     bool
}

// Trie implements a thread-safe search tree that stores byte key value pairs
// and allows for efficient queries.
type Trie struct {
	rw   sync.RWMutex
	root *trieNode
	size int
}

// NewTrie returns a new initialized empty Trie.
func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{root: true, children: make(map[byte]*trieNode)},
		size: 1,
	}
}

func newNode(symbol byte) *trieNode {
	return &trieNode{children: make(map[byte]*trieNode), symbol: symbol}
}

// Size returns the total number of nodes in the trie. The size includes the
// root node.
func (t *Trie) Size() int {
	t.rw.RLock()
	defer t.rw.RUnlock()
	return t.size
}

// Insert inserts a key value pair into the trie. If the key already exists,
// the value is updated. Insertion is performed by starting at the root
// and traversing the nodes all the way down until the key is exhausted. Once
// exhausted, the currNode pointer should be a pointer to the last symbol in
// the key and reflect the terminating node for that key value pair.
func (t *Trie) Insert(key, value []byte) {
	t.rw.Lock()
	defer t.rw.Unlock()

	currNode := t.root

	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			currNode.children[symbol] = newNode(symbol)
		}

		currNode = currNode.children[symbol]
	}

	// Only increment size if the key value pair is new, otherwise we consider
	// the operation as an update.
	if currNode.value == nil {
		t.size++
	}

	currNode.value = value
}

// Search attempts to search for a value in the trie given a key. If such a key
// exists, it's value is returned along with a boolean to reflect that the key
// exists. Otherwise, an empty value and false is returned.
func (t *Trie) Search(key []byte) ([]byte, bool) {
	t.rw.RLock()
	defer t.rw.RUnlock()

	currNode := t.root

	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			return nil, false
		}

		currNode = currNode.children[symbol]
	}

	return currNode.value, true
}

// GetAllKeys returns all the keys that exist in the trie. Keys are retrieved
// by performing a DFS on the trie where at each node we keep track of the
// current path (key) traversed thusfar and if that node has a value. If so,
// the full path (key) is appended to a list. After the trie search is
// exhausted, the final list is returned.
func (t *Trie) GetAllKeys() [][]byte {
	visited := make(map[*trieNode]bool)
	keys := [][]byte{}

	var dfsGetKeys func(n *trieNode, key []byte)
	dfsGetKeys = func(n *trieNode, key []byte) {
		if n != nil {
			pathKey := append(key, n.symbol)
			visited[n] = true

			if n.value != nil {
				fullKey := make([]byte, len(pathKey))
				// Copy the contents of the current path (key) to a new key so
				// future recursive calls will contain the correct bytes.
				copy(fullKey, pathKey)
				// Append the path (key) to the key list ignoring the first
				// byte which is the root symbol.
				keys = append(keys, fullKey[1:])
			}

			for _, child := range n.children {
				if _, ok := visited[child]; !ok {
					dfsGetKeys(child, pathKey)
				}
			}
		}
	}

	dfsGetKeys(t.root, []byte{})
	return keys
}
