# go-trie

[![GoDoc](https://godoc.org/github.com/alexanderbez/go-trie?status.svg)](https://godoc.org/github.com/alexanderbez/go-trie)

A primitive implementation of the Trie data structure in Golang. The Trie data
structure allows for quick information retrieval, typically strings, and prefix
searching. It stores data as an associative array. Keys and their respective
values are stored as bytes.

Typically, values in a Trie are unnecessary, rather, they are not traditionally
needed. In such cases such as these, you can simply store some terminal or
constant value.

## Potential Additional Functionality

- Get all keys that match a prefix
- Get all values where each respective key matches a prefix

## Usage

```golang
// Construct an empty trie with a single root node
trie := NewTrie()

// Insert a key value pair into the trie
trie.Insert([]byte("someKey"), []byte("someValue"))

// Search for some value
value, ok := trie.Search([]byte("someKey"))

// Get all keys stored in the trie
keys := trie.GetAllKeys()

// Get all values stored in the trie
values := trie.GetAllValues()
```

## Tests

```shell
$ go test -v
```
