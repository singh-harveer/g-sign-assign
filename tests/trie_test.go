package tests

import (
	"log"
	"testing"

	"github.com/singh-harveer/g-sign-assign/ds/trie"
)

func TestTrieInsert_01(t *testing.T) {
	trieObject := trie.NewTrie()
	testData := []string{"this", "is", "test", "case", "one", "data"}
	for _, v := range testData {
		trieObject.Insert(v)
	}
}

func TestTrieInsert_03(t *testing.T) {
	trieObject := trie.NewTrie()
	testData := []string{"this1@334455", "#", "*&#@#$%^&/////", "  "}
	for _, v := range testData {
		node := trieObject.Insert(v)
		if node != nil {
			// node is not nil which means test case failed
			log.Fatalf("Failed! expecting 'nil' but got %v\n", node)

		}
	}
}
