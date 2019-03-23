package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Map struct {
	replicas int
	hash     func([]byte) uint32
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, hash func([]byte) uint32) *Map {
	return &Map{
		replicas: replicas,
		hash:     hash,
		hashMap:  make(map[int]string),
	}
}
func (m *Map) Add(keys ...string) {
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {

			v := int(m.hash([]byte(key + strconv.Itoa(i))))
			m.hashMap[v] = key
			println(v)
			m.keys = append(m.keys, v)
		}
	}
	sort.Ints(m.keys)

}

func (m *Map) Get(value string) (node string) {

	vhash := int(m.hash([]byte(value)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= vhash
	})
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]

}
