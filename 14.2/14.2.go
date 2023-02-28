package main

import (
	"fmt"
)

const salt int = 1893574

type hashmap struct {
	hashMap map[uint64]string
}

func main() {

	// fmt.Println(hashstr("avav")) //1247902518561550
	// fmt.Println(hashstr("vava")) //1517750106487008

	hash := hashmap{}

	hash.Set("Первый ключ", "первое значение ")
	fmt.Println(hash.Get("Первый Ключ ")) // Значение отсутствует! false
	fmt.Println(hash.Get("Первый ключ"))
	hash.Set("Второй ключ", "второе значение ")
	fmt.Println(hash)
	hash.Delete("Первый ключ")
	fmt.Println(hash)
}

// hashstr - хеш-функция возвращает хеш типа uint64 от строки, используя остаток от деления на 1000 (по условию задачи)
func hashstr(val string) uint64 {
	res := 0
	for i, v := range val {
		res = ((i + int(v) + res) * salt) / 1000
	}
	return uint64(res)
}

// Set - добовляет значение
func (h *hashmap) Set(key, val string) {
	//инициализируем мапу
	if h.hashMap == nil {
		h.hashMap = make(map[uint64]string)
	}
	h.hashMap[hashstr(key)] = val
}

// Get - возвращает значение по ключу
func (h *hashmap) Get(key string) (value string, ok bool) {
	//проверяем мапу на валидность значения
	value, ok = h.hashMap[hashstr(key)]
	if !ok {
		return "Значение отсутствует!", false
	}
	return value, true
}

// Delete - удаляем значение по ключу
func (h *hashmap) Delete(key string) {
	delete(h.hashMap, hashstr(key))
}
