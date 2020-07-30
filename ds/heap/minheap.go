package heap

import (
	"fmt"

	"github.com/singh-harveer/g-sign-assign/ds/trie"
)

const MaxheapSize int = 20

type minHeapNode struct {
	Word      string
	Frequency int
	trNode    *trie.TrieNode
}
type minHeap struct {
	heapSlice   []*minHeapNode
	currentSize int
}

func NewMinheap() *minHeap {
	return &minHeap{
		heapSlice:   make([]*minHeapNode, MaxheapSize),
		currentSize: 0,
	}
}

func (heap *minHeap) Insert(currentTrieNode *trie.TrieNode, word string) {

	if currentTrieNode.MinHeapIndex != -1 {
		// update frequency
		heap.heapSlice[currentTrieNode.MinHeapIndex].Frequency = currentTrieNode.Frequency
		heap.MinHeapify(currentTrieNode.MinHeapIndex)

	} else if heap.currentSize < MaxheapSize {
		// if word is not present in heap and heap is not full
		node := &minHeapNode{
			Word:      word,
			Frequency: currentTrieNode.Frequency,
			trNode:    currentTrieNode,
		}
		heap.heapSlice[heap.currentSize] = node
		currentTrieNode.MinHeapIndex = heap.currentSize
		heap.currentSize++
		heap.BuildHeap()

	} else if heap.currentSize == MaxheapSize && currentTrieNode.Frequency > heap.heapSlice[0].Frequency {
		//when heapSlice is full and curent node still not inserted into heap
		// if currentTrieNode's frequency is greater then node at 0 index in minHeap
		// then remove min frequency node from heapSlice and insert currentTrieNode

		heap.heapSlice[0].trNode.MinHeapIndex = -1
		heap.heapSlice[0] = &minHeapNode{
			Frequency: currentTrieNode.Frequency,
			Word:      word,
			trNode:    currentTrieNode,
		}
		currentTrieNode.MinHeapIndex = 0
		heap.MinHeapify(0)
	}
	// now heapify it

}

func (heap *minHeap) MinHeapify(index int) {

	left := 2*index + 1
	right := 2*index + 2
	minimum := index

	if left < heap.currentSize && heap.heapSlice[left].Frequency < heap.heapSlice[index].Frequency {
		//make left node minmum node
		minimum = left
	} else if right < heap.currentSize && heap.heapSlice[right].Frequency < heap.heapSlice[index].Frequency {
		minimum = right
	}

	if minimum != index {
		heap.swap(index, minimum)
		// recursively call min heapify
		heap.MinHeapify(minimum)

	}

}
func (heap *minHeap) swap(idx1 int, idx2 int) {

	heap.heapSlice[idx1].trNode.MinHeapIndex = idx2
	heap.heapSlice[idx2].trNode.MinHeapIndex = idx1
	heap.heapSlice[idx1], heap.heapSlice[idx2] = heap.heapSlice[idx2], heap.heapSlice[idx1]
}

func (heap *minHeap) BuildHeap() {

	for i := ((heap.currentSize) / 2); i >= 0; i-- {
		heap.MinHeapify(i)
	}
}

func (heap *minHeap) Display() {

	for _, v := range heap.heapSlice {
		if v != nil {
			fmt.Println(v.Word, v.Frequency)
		}
	}
}
