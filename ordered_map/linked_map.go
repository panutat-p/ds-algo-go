package ordered_map

import (
	"container/list"
	"fmt"
	"strings"
)

type LinkedMap struct {
	dict map[string]int
	list *list.List
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		dict: make(map[string]int),
		list: list.New(),
	}
}

func (m *LinkedMap) Get(key string) int {
	v, ok := m.dict[key]
	if !ok {
		return 0
	}
	return v
}

func (m *LinkedMap) Set(key string, value int) {
	_, ok := m.dict[key]
	if !ok {
		m.list.PushBack(key)
		m.dict[key] = value
		return
	}
	m.dict[key] = value
}

// Put append key to list, and put key-value to map
// remove existing key if already exists
func (m *LinkedMap) Put(key string, value int) {
	m.Remove(key)
	m.list.PushBack(key)
	m.dict[key] = value
}

func (m *LinkedMap) Remove(key string) {
	_, ok := m.dict[key]
	if !ok {
		return
	}
	delete(m.dict, key)
	for e := m.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			m.list.Remove(e)
			break
		}
	}
}

func (m *LinkedMap) Keys() []string {
	keys := make([]string, 0, m.list.Len())
	for e := m.list.Front(); e != nil; e = e.Next() {
		key := e.Value.(string)
		keys = append(keys, key)
	}
	return keys
}

func (m *LinkedMap) Size() int {
	return m.list.Len()
}

func (m *LinkedMap) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for e := m.list.Front(); e != nil; e = e.Next() {
		key := e.Value.(string)
		value := m.dict[key]
		if e.Next() == nil {
			sb.WriteString(fmt.Sprintf("%s: %d", key, value))
		} else {
			sb.WriteString(fmt.Sprintf("%s: %d, ", key, value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
