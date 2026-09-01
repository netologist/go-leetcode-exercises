package lrucache

// Node represents a doubly linked list node holding key-value pairs.
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache implements a Least Recently Used (LRU) cache with O(1) Get and Put operations.
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // Dummy head (most recently used side)
	tail     *Node // Dummy tail (least recently used side)
}

// Constructor initializes the LRUCache with positive capacity.
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

// Get retrieves the value of the key if present, moving it to the front (MRU).
// Returns -1 if the key does not exist.
func (c *LRUCache) Get(key int) int {
	node, exists := c.cache[key]
	if !exists {
		return -1
	}
	c.moveToHead(node)
	return node.value
}

// Put inserts or updates the key-value pair.
// If capacity is exceeded, it evicts the least recently used item (before dummy tail).
func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		node.value = value
		c.moveToHead(node)
		return
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	c.cache[key] = newNode
	c.addToHead(newNode)

	if len(c.cache) > c.capacity {
		lru := c.removeTail()
		delete(c.cache, lru.key)
	}
}

// addToHead inserts a node right after the dummy head.
func (c *LRUCache) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

// removeNode unlinks an existing node from the list.
func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// moveToHead moves an existing node to the MRU position (after dummy head).
func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

// removeTail removes and returns the LRU node right before the dummy tail.
func (c *LRUCache) removeTail() *Node {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
