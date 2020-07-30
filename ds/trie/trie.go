package trie

import (
	"unicode"
)

const (
	ALBHABET_SIZE = 1000
)

//TrieNode -  to define trieNode structure
type TrieNode struct {
	childrens    [ALBHABET_SIZE]*TrieNode
	Char         string
	isWordEnd    bool
	Frequency    int
	MinHeapIndex int
}

// trie define trie structure
type trie struct {
	Root *TrieNode
}

/**
Trie constructor
**/
func NewTrie() *trie {
	return &trie{
		Root: &TrieNode{MinHeapIndex: -1},
	}
}

// isAlpha - to check rune type
func isAlpha(v rune) bool {
	return unicode.IsLetter(v)
}

/**
Insert to insert new word into trie
**/
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
