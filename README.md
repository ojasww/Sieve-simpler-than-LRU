# Sieve-simpler-than-LRU

``SIEVE`` is Simpler than LRU: an Efficient Turn-Key Eviction Algorithm for Web Caches

This repository contains a **Golang** implementation for the web caching algorithm - ``SIEVE``.

[Paper Link](https://www.usenix.org/conference/nsdi24/presentation/zhang-yazhuo)

I [tweeted](https://x.com/ojasw_/status/1856209687676326286) about this.

## Description

The cache is a **Double Linked List**. It's double because the hand in the sieve which searches for the element
to evict moves in the direction from tail to head, hence easy to navigate backwards.

Any cache algorithm operates on two policies:
1. **Admission Policy** - Rules to insert an element (usually at the start of the queue)
2. **Eviction Policy** - Rules to evict element.

Sieve is different from other algorithms like ``FIFO``, ``FIFO-Reinsertion`` and ``LRU`` in one particular way - it applies **lazy promotion** and **quick demotion**. For this, it has to maintain a **visited** flag for each node that keeps the element in place after a cache hit (lazy promotion). The fast movement of the hand in search for the element to be evicted results in quick demotion.

The cache struct looks like the following:

```go
type Cache struct {
    head    *Node
    tail    *Node
    hand    *Node
    size    int
    maxSize int
}
```

``hand`` is a pointer in the list that searches for the first element with visited flag as **zero** while moving backwards. If it encounters an element with flag as **one** it flips it and continues searching.

## Testing

Tests cases for testing the `Initialization`, `Add`, `Evict`, `Hit`, `EvictionWithVisited` and `EmptyEviction` are present in `main_test.go`. To test the exhaustive test cases, run:

```go
go test
```

## Resources

- Animated explanation: https://cachemon.github.io/SIEVE-website/
- [Cache Replacement Poicies](https://en.wikipedia.org/wiki/Cache_replacement_policies)