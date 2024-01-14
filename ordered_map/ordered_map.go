package ordered_map

import (
	"fmt"
	"strings"
)

type OrderedMap struct {
	dict map[string]int
	list []string
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		dict: make(map[string]int),
		list: make([]string, 0),
	}
}

func (m *OrderedMap) Get(key string) int {
	v, ok := m.dict[key]
	if !ok {
		return 0
	}
	return v
}

func (m *OrderedMap) Set(key string, value int) {
	_, ok := m.dict[key]
	if !ok {
		for i, e := range m.list {
			if e == key {
				m.list[i] = key
				break
			}
		}
		m.dict[key] = value
		return
	}
	m.dict[key] = value
}

// Put append key to list, and put key-value to map
// remove existing key if already exists
func (m *OrderedMap) Put(key string, value int) {
	m.Remove(key)
	m.list = append(m.list, key)
	m.dict[key] = value
}

func (m *OrderedMap) Remove(key string) {
	_, ok := m.dict[key]
	if !ok {
		return
	}
	delete(m.dict, key)
	for i, e := range m.list {
		if e == key {
			m.list = append(m.list[:i], m.list[i+1:]...)
			break
		}
	}
}

func (m *OrderedMap) Keys() []string {
	return m.list
}

func (m *OrderedMap) Size() int {
	return len(m.list)
}

func (m *OrderedMap) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, key := range m.list {
		if i == len(m.list)-1 {
			sb.WriteString(fmt.Sprintf("%s: %d", key, m.dict[key]))
		} else {
			sb.WriteString(fmt.Sprintf("%s: %d, ", key, m.dict[key]))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
