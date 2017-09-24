# go-trie

A primitive implementation of the Trie data structure in Golang. The Trie data
structure allows for quick information retrieval, typically strings.

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
