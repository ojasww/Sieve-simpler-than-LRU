# Sieve-simpler-than-LRU

SIEVE is Simpler than LRU: an Efficient Turn-Key Eviction Algorithm for Web Caches

This repository contains a **Golang** implementation for the web caching algorithm - **Sieve**.

[Paper Link](https://www.usenix.org/conference/nsdi24/presentation/zhang-yazhuo)

## Description

The cache is a Double Linked List. It's double because the hand in the sieve which searches for the element
to evict moves in the direction from tail to head, hence easy to navigate backwards.

The cache struct looks like the following:

```go
    type Cache struct {
        head    *Node
        tail    *Node
        hand    *Node
        size    int
        maxSize int
    }
``

## Testing

Just run the main file:

```go
go run main.go
```

## Resources

Animated explanation: https://cachemon.github.io/SIEVE-website/

