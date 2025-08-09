package intern

import "sync"

type Table struct {
	mu  sync.RWMutex
	id  uint32
	s2i map[string]uint32
	i2s []string
}

func New() *Table {
	return &Table{s2i: make(map[string]uint32)}
}

func (t *Table) Intern(s string) uint32 {
	t.mu.RLock()
	if id, ok := t.s2i[s]; ok {
		t.mu.RUnlock()
		return id
	}
	t.mu.RUnlock()

	t.mu.Lock()
	if id, ok := t.s2i[s]; ok {
		t.mu.Unlock()
		return id
	}
	id := t.id
	t.id++
	t.s2i[s] = id
	t.i2s = append(t.i2s, s)
	t.mu.Unlock()
	return id
}

func (t *Table) Lookup(id uint32) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if int(id) < len(t.i2s) {
		return t.i2s[id]
	}
	return ""
}
