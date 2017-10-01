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
trie := NewTrie()

trie.Insert([]byte("someKey"), []byte("someValue"))

value, ok := trie.Search([]byte("someKey"))
// value => []byte("someValue")
// ok => true
```

## Tests

```shell
$ go test -v
```
