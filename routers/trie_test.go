package routers

import (
	"testing"
)

func TestCreateTrie(t *testing.T) {
	trie := CreateTrie()
	if trie.value != nil {
		t.Errorf("New trie's value is not nil: %v", trie.value)
	}
	if len(trie.children) != 0 {
		t.Errorf("New trie's children is not empty: %v", trie.children)
	}
}

func TestTrieInsert(t *testing.T) {
	trie := CreateTrie()

	trie.Insert("/", "pie")

	if len(trie.children) != 1 {
		t.Errorf("Inserted more than 1 child into trie")
	}

	if trie.children['/'].value != "pie" {
		t.Errorf("Value inserted at root %v != %v", trie.children[0].value, "pie")
	}

	trie.Insert("/abc", 42)

	if len(trie.children) != 1 {
		t.Errorf("Inserted a new children with the same rune")
	}

	node := trie

	for idx, r := range "/abc" {
		if child, ok := node.children[r]; !ok {
			t.Errorf("Didn't insert rune %v (%v) into trie", r, idx)
		} else {
			node = child
		}
	}

	if node.value != 42 {
		t.Errorf("Inserted wrong value at tail of trie: %v != 42", node.value)
	}

	trie.Insert("/bcd", 3.14)
	node = trie.children['/']

	if len(node.children) != 2 {
		t.Errorf("Expected 2 children, got %v", len(node.children))
	}
}