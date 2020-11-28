package loadbalancer

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type HashBalance struct {
	mux      sync.RWMutex
	hash     Hash
	replicas int
	keys     UInt32Slice
	hashMap  map[uint32]string
}

func NewHashBalance(replicas int, fn Hash) *HashBalance {
	m := &HashBalance{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[uint32]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (h *HashBalance) IsEmpty() bool {
	return len(h.keys) == 0
}

func (h *HashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	h.mux.Lock()
	defer h.mux.Unlock()
	for i := 0; i < h.replicas; i++ {
		hash := h.hash([]byte(strconv.Itoa(i) + addr))
		h.keys = append(h.keys, hash)
		h.hashMap[hash] = addr
	}
	sort.Sort(h.keys)
	return nil
}

func (h *HashBalance) Get(key string) (string, error) {
	if h.IsEmpty() {
		return "", errors.New("node is empty")
	}
	hash := h.hash([]byte(key))
	idx := sort.Search(len(h.keys), func(i int) bool {
		return h.keys[i] >= hash
	})
	if idx == len(h.keys) {
		idx = 0
	}
	h.mux.RLock()
	defer h.mux.Unlock()
	return h.hashMap[h.keys[idx]], nil
}
