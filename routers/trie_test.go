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
	trie.Insert("/", nil)
	trie.Insert("/foo", 42)
}