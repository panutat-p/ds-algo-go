package ordered_map

import (
	"container/list"
	"fmt"
	"strings"
)

type entry struct {
	key   string
	value int
}

type LinkedHashMap struct {
	dict map[string]*list.Element
	list *list.List
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		dict: make(map[string]*list.Element),
		list: list.New(),
	}
}

func (m *LinkedHashMap) Get(key string) int {
	elem, ok := m.dict[key]
	if !ok {
		return 0
	}
	return elem.Value.(*entry).value
}

// Put append key to list, and put key-value to map
// remove existing key if already exists
func (m *LinkedHashMap) Put(key string, value int) {
	elem, ok := m.dict[key]
	if ok {
		m.list.Remove(elem)
	}
	newElem := m.list.PushBack(&entry{key, value})
	m.dict[key] = newElem
}

func (m *LinkedHashMap) Remove(key string) {
	elem, ok := m.dict[key]
	if !ok {
		return
	}
	m.list.Remove(elem)
	delete(m.dict, key)
}

func (m *LinkedHashMap) Keys() []string {
	keys := make([]string, 0, m.list.Len())
	for e := m.list.Front(); e != nil; e = e.Next() {
		key := e.Value.(*entry).key
		keys = append(keys, key)
	}
	return keys
}

func (m *LinkedHashMap) Size() int {
	return m.list.Len()
}

func (m *LinkedHashMap) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for e := m.list.Front(); e != nil; e = e.Next() {
		key := e.Value.(*entry).key
		value := e.Value.(*entry).value
		if e.Next() == nil {
			sb.WriteString(fmt.Sprintf("%s: %d", key, value))
		} else {
			sb.WriteString(fmt.Sprintf("%s: %d, ", key, value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
