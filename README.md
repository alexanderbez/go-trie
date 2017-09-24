# go-trie

A primitive implementation of the Trie data structure in Golang. The Trie data
structure allows for quick information retrieval, typically strings. Keys and
their respective values are stored as bytes.

Typically, values in a Trie are unnecessary or are not traditionally used. In
cases such as these, you can simply store a terminal value.

## Potential Additional Functionality

- Get all keys
- Get all values
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
