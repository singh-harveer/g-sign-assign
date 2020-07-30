package trie

import (
	"unicode"
)

const (
	ALBHABET_SIZE = 10000
)

type TrieNode struct {
	childrens    [ALBHABET_SIZE]*TrieNode
	Char         string
	isWordEnd    bool
	Frequency    int
	MinHeapIndex int
}

type trie struct {
	Root *TrieNode
}

func NewTrie() *trie {
	return &trie{
		Root: &TrieNode{MinHeapIndex: -1},
	}
}

func isAlpha(v rune) bool {
	return unicode.IsLetter(v)
}

func (t *trie) Insert(word string) *TrieNode {
	wordRuneSlice := []rune(word)
	current := t.Root
	for _, value := range wordRuneSlice {

		// if !isAlpha(value) {
		// 	return nil
		// }
		index := value // - rune('a')
		if current.childrens[index] == nil {
			current.childrens[index] = &TrieNode{MinHeapIndex: -1, Char: string(value)}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
	(current.Frequency)++
	return current
}
