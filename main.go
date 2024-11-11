package main

import (
	"errors"
	"fmt"
)

type Sieve interface {
	Evict() error
	Add(value int) error
}

type Node struct {
	value   int
	visited bool
	next    *Node
	prev    *Node
}

type Cache struct {
	head    *Node
	tail    *Node
	hand    *Node
	size    int
	maxSize int
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		head:    nil,
		tail:    nil,
		hand:    nil,
		size:    0,
		maxSize: maxSize,
	}
}

func (c *Cache) Add(value int) error {
	if c.size == c.maxSize {
		if err := c.Evict(); err != nil {
			return err
		}
	}

	newNode := &Node{value: value, visited: false}

	if c.head == nil {
		c.head = newNode
		c.tail = newNode
	} else {
		newNode.next = c.head
		c.head.prev = newNode
		c.head = newNode
	}

	c.size++
	return nil
}

func (c *Cache) Evict() error {
	if c.size == 0 {
		return errors.New("cache is empty")
	}

	if c.hand == nil {
		c.hand = c.tail
	}

	for c.hand != nil {
		if c.hand.visited {
			c.hand.visited = false
			c.hand = c.hand.prev
		} else {
			if c.hand.prev == nil && c.hand.next == nil {
				c.head = nil
				c.tail = nil
			} else if c.hand.prev == nil {
				c.head = c.hand.next
				c.head.prev = nil
			} else if c.hand.next == nil {
				c.tail = c.hand.prev
				c.tail.next = nil
			} else {
				c.hand.prev.next = c.hand.next
				c.hand.next.prev = c.hand.prev
			}

			c.size--
			return nil
		}
	}

	return nil
}

func (c *Cache) Hit(value int) error {
	node := c.head
	for node != nil {
		if node.value == value {
			node.visited = true
			return nil
		}
		node = node.next
	}

	return c.Add(value)
}

func (c *Cache) PrintCache() {
	node := c.head
	for node != nil {
		fmt.Printf("Value: %d, Visited: %v\n", node.value, node.visited)
		node = node.next
	}
	fmt.Println("\n")
}

func main() {
	cache := NewCache(3)

	cache.Hit(1)
	cache.Hit(2)
	cache.Hit(3)
	cache.PrintCache()

	cache.Hit(1)
	cache.PrintCache()

	cache.Hit(4)
	cache.PrintCache()

	cache.Hit(5)
	cache.PrintCache()
}
