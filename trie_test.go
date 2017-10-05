package trie

import (
	"bytes"
	"testing"
)

// byteSliceEq returns true iff two slices of bytes contain the same elements.
func byteSliceEq(s1, s2 []Bytes) bool {
	if len(s1) != len(s2) {
		return false
	}

	result := true
	i, n := 0, len(s1)

	for i < n {
		j := 0
		e := false
		for j < n && !e {
			if bytes.Equal(s1[i], s2[j]) {
				e = true
			}
			j++
		}
		if !e {
			result = false
		}
		i++
	}

	return result
}

func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.root == nil {
		t.Errorf("trie root is invalid. expected: %v (got %v)", &trieNode{}, trie.root)
	}

	if trie.size != 1 {
		t.Errorf("trie size is invalid. expected: %v (got %v)", 1, trie.size)
	}
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()

	if trie.Size() != 1 {
		t.Errorf("trie size is invalid. expected: %v (got %v)", 1, trie.Size())
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
			t.Errorf("trie size is invalid. expected: %v (got %v)", size, trie.Size())
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
			t.Errorf("unable to find key %v. expected: %v (got %v)", k, true, ok)
		}

		if !bytes.Equal(v, eV) {
			t.Errorf("invalid value for key %v. expected: %v (got %v)", k, eV, v)
		}
	}

	k := []byte("idontexist")
	v, ok := trie.Search(k)
	if ok {
		t.Errorf("unable to find key %v. expected: %v (got %v)", k, false, ok)
	}

	if len(v) != 0 {
		t.Errorf("invalid value for key %v. expected: %v (got %v)", k, []byte{}, v)
	}
}

func TestGetAllValues(t *testing.T) {
	trie := NewTrie()

	trieVals := trie.GetAllValues()

	if len(trieVals) != 0 {
		t.Errorf("invalid length of values returned. expected: %v (got %v)", 0, len(trieVals))
	}

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

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 7,
			"expectedValues": []Bytes{
				[]byte{1, 2, 3, 4},
				[]byte{2, 1, 4, 6},
				[]byte{4, 6, 1, 1},
				[]byte{7, 7, 4, 4},
				[]byte{8, 1, 1, 9},
				[]byte{9, 0, 1, 1},
				[]byte{6, 4, 2, 1},
			},
		},
	}

	for _, tc := range testCases {
		trieVals = trie.GetAllValues()

		if len(trieVals) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of values returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieVals),
			)
		}

		if !byteSliceEq(trieVals, tc["expectedValues"].([]Bytes)) {
			t.Errorf("missing value from expected list of values. expected: %v (got %v)",
				tc["expectedKeys"].([]Bytes),
				trieVals,
			)
		}
	}

}

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

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 8,
			"expectedKeys": []Bytes{
				Bytes("baby"),
				Bytes("bad"),
				Bytes("badly"),
				Bytes("bank"),
				Bytes("box"),
				Bytes("dad"),
				Bytes("dance"),
				Bytes("zip"),
			},
		},
	}

	for _, tc := range testCases {
		trieKeys = trie.GetAllKeys()

		if len(trieKeys) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of keys returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieKeys),
			)
		}

		if !byteSliceEq(trieKeys, tc["expectedKeys"].([]Bytes)) {
			t.Errorf("missing key from expected list of keys. expected: %v (got %v)",
				tc["expectedKeys"].([]Bytes),
				trieKeys,
			)
		}
	}
}

func TestGetPrefixKeys(t *testing.T) {
	trie := NewTrie()

	prefix := []byte{}
	trieKeys := trie.GetPrefixKeys(prefix)

	if len(trieKeys) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", 0, len(trieKeys))
	}

	kvPairs := map[string]Bytes{
		"baby":  Bytes{1, 2, 3, 4},
		"bad":   Bytes{2, 1, 4, 6},
		"badly": Bytes{4, 6, 1, 1},
		"bank":  Bytes{7, 7, 4, 4},
		"box":   Bytes{8, 1, 1, 9},
		"dad":   Bytes{9, 0, 1, 1},
		"dance": Bytes{6, 4, 2, 1},
		"zip":   Bytes{0, 0, 1, 2},
	}

	for k, v := range kvPairs {
		trie.Insert(Bytes(k), v)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"prefix":      "z",
			"expectedLen": 1,
			"expectedKeys": []Bytes{
				Bytes("zip"),
			},
		},
		map[string]interface{}{
			"prefix":      "ba",
			"expectedLen": 4,
			"expectedKeys": []Bytes{
				Bytes("baby"),
				Bytes("bad"),
				Bytes("badly"),
				Bytes("bank"),
			},
		},
		map[string]interface{}{
			"prefix":      "bad",
			"expectedLen": 2,
			"expectedKeys": []Bytes{
				Bytes("bad"),
				Bytes("badly"),
			},
		},
	}

	for _, tc := range testCases {
		prefix = Bytes(tc["prefix"].(string))
		trieKeys = trie.GetPrefixKeys(prefix)

		if len(trieKeys) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of keys returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieKeys),
			)
		}

		if !byteSliceEq(trieKeys, tc["expectedKeys"].([]Bytes)) {
			t.Errorf("missing key from expected list of keys for prefix: %v. expected: %v (got %v)",
				prefix,
				tc["expectedKeys"].([]Bytes),
				trieKeys,
			)
		}
	}
}

func TestGetPrefixValues(t *testing.T) {
	trie := NewTrie()

	prefix := []byte{}
	trieValues := trie.GetPrefixValues(prefix)

	if len(trieValues) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", 0, len(trieValues))
	}

	kvPairs := map[string]Bytes{
		"baby":  Bytes{1, 2, 3, 4},
		"bad":   Bytes{2, 1, 4, 6},
		"badly": Bytes{4, 6, 1, 1},
		"bank":  Bytes{7, 7, 4, 4},
		"box":   Bytes{8, 1, 1, 9},
		"dad":   Bytes{9, 0, 1, 1},
		"dance": Bytes{6, 4, 2, 1},
		"zip":   Bytes{0, 0, 1, 2},
	}

	for k, v := range kvPairs {
		trie.Insert(Bytes(k), v)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"prefix":      "z",
			"expectedLen": 1,
			"expectedValues": []Bytes{
				Bytes{0, 0, 1, 2},
			},
		},
		map[string]interface{}{
			"prefix":      "ba",
			"expectedLen": 4,
			"expectedValues": []Bytes{
				Bytes{1, 2, 3, 4},
				Bytes{2, 1, 4, 6},
				Bytes{4, 6, 1, 1},
				Bytes{7, 7, 4, 4},
			},
		},
		map[string]interface{}{
			"prefix":      "bad",
			"expectedLen": 2,
			"expectedValues": []Bytes{
				Bytes{2, 1, 4, 6},
				Bytes{4, 6, 1, 1},
			},
		},
	}

	for _, tc := range testCases {
		prefix = Bytes(tc["prefix"].(string))
		trieValues = trie.GetPrefixValues(prefix)

		if len(trieValues) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of values returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieValues),
			)
		}

		if !byteSliceEq(trieValues, tc["expectedValues"].([]Bytes)) {
			t.Errorf("missing value from expected list of values for prefix: %v. expected: %v (got %v)",
				string(prefix),
				tc["expectedValues"].([]Bytes),
				trieValues,
			)
		}
	}
}
