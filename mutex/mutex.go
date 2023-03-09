package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    *sync.Mutex // Без инициализации = nil
}

func (c *counter) inc() {
	c.mu.Lock() // Разблокируем
	c.count++
	c.mu.Unlock() //Блокируем
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
	/*если без defer то
	c.mu.Lock() // разблокируем для чтения
	value:= c.count // читаем
	c.mu.Unlock() //блокируем
	return value // возвращаем
	*/

}

func main() {
	c := counter{
		mu: new(sync.Mutex), // Инициализируем Mutex
	}

	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}
	//time.Sleep(time.Second)
	fmt.Println(c.value())
}
