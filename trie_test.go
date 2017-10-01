package trie

import (
	"bytes"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.root == nil {
		t.Errorf("Trie root is invalid (expected: %v) (got %v)", &trieNode{}, trie.root)
	}

	if trie.size != 1 {
		t.Errorf("Trie size is invalid (expected: %v) (got %v)", 1, trie.size)
	}
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()

	if trie.Size() != 1 {
		t.Errorf("Trie size is invalid. Expected: %v (got %v)", 1, trie.Size())
	}
}

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	size := 1

	kvPairs := map[string][]byte{
		"baby":  []byte{1, 2, 3, 4},
		"bad":   []byte{2, 1, 4, 6},
		"badly": []byte{4, 6, 1, 1},
		"bank":  []byte{7, 7, 4, 4},
		"box":   []byte{8, 1, 1, 9},
		"dad":   []byte{9, 0, 1, 1},
		"dance": []byte{6, 4, 2, 1},
	}

	for k, v := range kvPairs {
		trie.Insert([]byte(k), v)
		size++

		if trie.Size() != size {
			t.Errorf("Trie size is invalid. Expected: %v (got %v)", size, trie.Size())
		}
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()

	kvPairs := map[string][]byte{
		"baby":  []byte{1, 2, 3, 4},
		"bad":   []byte{2, 1, 4, 6},
		"badly": []byte{4, 6, 1, 1},
		"bank":  []byte{7, 7, 4, 4},
		"box":   []byte{8, 1, 1, 9},
		"dad":   []byte{9, 0, 1, 1},
		"dance": []byte{6, 4, 2, 1},
	}

	for k, v := range kvPairs {
		trie.Insert([]byte(k), v)
	}

	for k, eV := range kvPairs {
		v, ok := trie.Search([]byte(k))

		if !ok {
			t.Errorf("Unable to find key %v. Expected: %v (got %v)", k, true, ok)
		}

		if !bytes.Equal(v, eV) {
			t.Errorf("Invalid value for key %v. Expected: %v (got %v)", k, eV, v)
		}
	}

	k := []byte("idontexist")
	v, ok := trie.Search(k)
	if ok {
		t.Errorf("Unable to find key %v. Expected: %v (got %v)", k, false, ok)
	}

	if len(v) != 0 {
		t.Errorf("Invalid value for key %v. Expected: %v (got %v)", k, []byte{}, v)
func TestGetAllKeys(t *testing.T) {
	trie := NewTrie()

	trieKeys := trie.GetAllKeys()

	if len(trieKeys) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", 0, len(trieKeys))
	}

	kvPairs := map[string][]byte{
		"baby":  []byte{1, 2, 3, 4},
		"bad":   []byte{2, 1, 4, 6},
		"badly": []byte{4, 6, 1, 1},
		"bank":  []byte{7, 7, 4, 4},
		"box":   []byte{8, 1, 1, 9},
		"dad":   []byte{9, 0, 1, 1},
		"dance": []byte{6, 4, 2, 1},
		"zip":   []byte{0, 0, 1, 2},
	}

	for k, v := range kvPairs {
		trie.Insert([]byte(k), v)
	}

	trieKeys = trie.GetAllKeys()
	keysMap := make(map[string]bool)

	for _, k := range trieKeys {
		keysMap[string(k)] = true
	}

	if len(trieKeys) != len(kvPairs) {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", len(kvPairs), len(trieKeys))
	}

	for k := range kvPairs {
		if !keysMap[k] {
			t.Errorf("missing key from expected list of keys. expected: %v", k)
		}
	}
}
