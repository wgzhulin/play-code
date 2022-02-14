package hashtable

import (
	"crypto/sha512"
	"github.com/play-code/src/basic/datastructure/basedata/ele"
)

const maxTableSize = 10

type HashTable struct {
	data [maxTableSize]*ele.Ele
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (h *HashTable) hash(key string) int {
	hashFunc := sha512.New()
	hash, _ := hashFunc.Write([]byte(key))
	return hash % maxTableSize
}

func (h *HashTable) Add(key string, ele *ele.Ele) {
	index := h.hash(key)
	h.data[index] = ele
}

func (h *HashTable) Exists(key string) bool {
	index := h.hash(key)
	return h.data[index] != nil
}

func (h *HashTable) Get(key string) *ele.Ele {
	index := h.hash(key)
	return h.data[index]
}

func (h *HashTable) Remove(key string) {
	index := h.hash(key)
	h.data[index] = nil
}
