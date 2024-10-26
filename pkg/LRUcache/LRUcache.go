package main

import (
	"container/list"
	"sync"
)

/*потокобезопасныый кеш ,но возможно не тестированный*/
type LRUCache struct {
	capacity int
	m        map[interface{}]*list.Element
	list     *list.List
	mu       *sync.RWMutex
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[interface{}]*list.Element),
		list:     list.New(),
		mu:       &sync.RWMutex{},
	}
}

// Если ключ не найден, вернёт nil
func (this *LRUCache) Get(key int) interface{} {
	this.mu.RLock()         // Блокировка для чтения
	defer this.mu.RUnlock() // Разблокировка после чтения
	if elem, found := this.m[key]; found {
		this.list.MoveToFront(elem) // Переместить элемент в начало списка
		return elem.Value
	}
	return nil
}

func (this *LRUCache) Put(key int, value interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if elem, found := this.m[key]; found {
		// Если ключ уже существует, обновить значение и переместить в начало
		elem.Value = value
		this.list.MoveToFront(elem)
	} else {
		// Если ключ не существует, добавить новый элемент
		if this.list.Len() == this.capacity {
			// Если кэш полон, удалить наименее недавно использованный элемент
			backElem := this.list.Back()
			if backElem != nil {
				this.list.Remove(backElem)
				// Удаляем элемент из карты. Значение должно быть ключом
				for k, v := range this.m {
					if v == backElem {
						delete(this.m, k) // Удалить ключ из карты
						break
					}
				}
			}
		}
		// Добавляем новый элемент
		newElem := this.list.PushFront(value)
		this.m[key] = newElem // Сохраняем связь между ключом и элементом
	}
}

/*
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
 */
