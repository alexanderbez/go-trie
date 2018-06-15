# go-trie

[![GoDoc](https://godoc.org/github.com/alexanderbez/go-trie?status.svg)](https://godoc.org/github.com/alexanderbez/go-trie)
[![Build Status](https://travis-ci.org/alexanderbez/go-trie.svg?branch=master)](https://travis-ci.org/alexanderbez/go-trie)

A primitive implementation of the Trie data structure in Golang. The Trie data
structure allows for quick information retrieval, typically strings, and prefix
searching. It stores data as an associative array. Keys and their respective
values are stored as bytes.

Typically, values in a Trie are unnecessary, rather, they are not traditionally
needed. In such cases such as these, you can simply store some terminal or
constant value.

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

// Get all keys stored in the trie that contain a specific prefix
keyByPrefix := trie.GetPrefixKeys([]byte("somePrefix"))

// Get all values stored in the trie who's corresponding keys contain a
// specific prefix.
valuesByPrefix := trie.GetPrefixValues(prefix)
```

## Tests

```shell
$ go test -v
```

## Contributing

1. [Fork it](https://github.com/alexanderbez/go-trie/fork)
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

