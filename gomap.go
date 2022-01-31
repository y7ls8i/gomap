// Package gomap implements a map that is goroutine-safe.
package gomap

import "sync"

// GoMap contains functionality to store and retrieve data from the map.
type GoMap interface {
	Clear()
	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)
	All() map[string]interface{}
}

type gomap struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewGoMap creates a new goroutine-safe map.
func NewGoMap() GoMap {
	m := &gomap{}
	m.data = map[string]interface{}{}
	return m
}

// Clear empties the map.
func (m *gomap) Clear() {
	m.Lock()
	m.data = map[string]interface{}{}
	m.Unlock()
}

// Get returns the value based on the key. It returns nil if not found.
func (m *gomap) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	data, ok := m.data[key]
	if !ok {
		data = nil
	}

	return data
}

// Set replaces the value with the new one.
func (m *gomap) Set(key string, value interface{}) {
	m.Lock()
	m.data[key] = value
	m.Unlock()
}

// Remove deletes the key & value from the map.
func (m *gomap) Remove(key string) {
	m.Lock()
	delete(m.data, key)
	m.Unlock()
}

// All returns a normal map.
func (m *gomap) All() map[string]interface{} {
	m.Lock()
	defer m.Unlock()

	copyData := map[string]interface{}{}
	for key, value := range m.data {
		copyData[key] = value
	}

	return copyData
}
