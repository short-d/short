package cache

// Cache stores key-value pair in memory,
// it uses classic implementation of LRU cache
type Cache struct {
	Head     *Node
	Tail     *Node
	Mapping  map[interface{}]*Node
	Capacity int
}

// Node defines a node in doubly linked list
type Node struct {
	Key  interface{}
	Val  []interface{}
	Prev *Node
	Next *Node
}

// Get retrieves the value inside cache
func (c *Cache) Get(key interface{}) []interface{} {
	node, ok := c.Mapping[key]
	if ok {
		c.Remove(node)
		c.Add(node)
		return node.Val
	}
	return []interface{}{}
}

// Set sets the key-value pair into cache
func (c *Cache) Set(key interface{}, value []interface{}) {
	node, ok := c.Mapping[key]
	if ok {
		node.Val = value
		c.Remove(node)
		c.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		c.Mapping[key] = node
		c.Add(node)
	}
	if len(c.Mapping) > c.Capacity {
		delete(c.Mapping, c.Tail.Key)
		c.Remove(c.Tail)
	}
}

// Add adds a new Node into the front of the doubly linked list
func (c *Cache) Add(node *Node) {
	node.Prev = nil
	node.Next = c.Head
	if c.Head != nil {
		c.Head.Prev = node
	}
	c.Head = node
	if c.Tail == nil {
		c.Tail = node
	}
}

// Remove removes a Node inside the doubly linked list
func (c *Cache) Remove(node *Node) {
	if node != c.Head {
		node.Prev.Next = node.Next
	} else {
		c.Head = node.Next
	}
	if node != c.Tail {
		node.Next.Prev = node.Prev
	} else {
		c.Tail = node.Prev
	}
}

// NewCache creates a Cache object
func NewCache(capacity int) Cache {
	return Cache{
		Mapping:  make(map[interface{}]*Node),
		Capacity: capacity,
	}
}
