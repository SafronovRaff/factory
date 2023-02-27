package main

import (
	"fmt"
	"time"
)

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type InMemoryCache struct {
	//Code here
	expireIn time.Duration
	cache    map[string]CacheEntry
}

// NewInMemoryCache - конструктор
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		cache:    make(map[string]CacheEntry),
	}
}

// Set - сохраняет новую запись в кеше вместе с текущим временем
func (c *InMemoryCache) Set(key string, value interface{}) {
	//Устанавливаем новое значение в кэше с текущей меткой времени
	c.cache[key] = CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
}

// Get - возвращает значение из кеша если оно есть или время действия не истекло
func (c *InMemoryCache) Get(key string) interface{} {
	// Получаем значение из кеша и проверяем, не истек ли срок действия
	entry, ok := c.cache[key]
	if ok && time.Since(entry.settledAt) <= c.expireIn { // Since возвращает время, прошедшее с момента t.
		return entry.value
	}
	return fmt.Sprintf("срок действия значения истек или его нет в кеше")
}

func main() {
	cashe := NewInMemoryCache(15 * time.Second)

	cashe.Set("первый ключ", "первое значение")
	fmt.Println(cashe.Get("первый ключ"))

	time.Sleep(20 * time.Second)
	fmt.Println(cashe.Get("первый ключ"))

}
