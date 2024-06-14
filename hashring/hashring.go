package hashring

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"sync"
)

// The node will be added to the Hash-Ring whenever a new system will join the servers
type Node struct {
	ID   string
	Addr string
}

type HashRing struct {
	Nodes       []Node
	SortedHash  []int
	NodeHashMap map[int]Node
	Replication int // the number of virtual nodes (or replicas) each physical node will have in the hash ring
	Mutex       sync.Mutex
}

func NewHashRing(replications int) *HashRing {
	return &HashRing{
		NodeHashMap: make(map[int]Node),
		Replication: replications,
	}
}

func (hr *HashRing) GenerateHashKey(key string) int {
	hash := sha256.Sum256([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	hashInt := 0
	fmt.Sscanf(hashString[:8], "%x", &hashInt)
	return hashInt
}

func (hr *HashRing) AddNewNodes(node Node) {
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	for i := 0; i < hr.Replication; i++ {
		hash := hr.GenerateHashKey(fmt.Sprintf("%s-%d", node.ID, i))
		hr.NodeHashMap[hash] = node
		hr.SortedHash = append(hr.SortedHash, hash)
	}
	sort.Ints(hr.SortedHash)
}

func (hr *HashRing) SearchKey(key int) int {
	idx := sort.Search(len(hr.SortedHash), func(i int) bool {
		return hr.SortedHash[i] >= key
	})
	if idx >= len(hr.SortedHash) {
		return 0
	}
	return idx
}

func (hr *HashRing) GetNode(key string) Node {
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	if len(hr.NodeHashMap) == 0 {
		return Node{}
	}

	hashKey := hr.GenerateHashKey(key)
	idx := hr.SearchKey(hashKey)

	return hr.NodeHashMap[hr.SortedHash[idx]]
}

// updateSortedKeys updates the sorted keys slice
func (hr *HashRing) updateSortedKeys() {
	hr.SortedHash = hr.SortedHash[:0]
	for hash := range hr.NodeHashMap {
		hr.SortedHash = append(hr.SortedHash, hash)
	}

	sort.Ints(hr.SortedHash)
}

// RemoveNode removes a node from the hash ring
func (hr *HashRing) RemoveNode(node Node) {
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	for i := 0; i < hr.Replication; i++ {
		hash := hr.GenerateHashKey(fmt.Sprintf("%s-%d", node.ID, i))
		delete(hr.NodeHashMap, hash)
	}

	hr.updateSortedKeys()
}
