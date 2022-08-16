package lru_cache_146

type Cache interface {
	Put(k, v string)
	Get(k string) (*string, error)
}

type cache struct {
	currentSize int
	capacity    int
	listMap     map[string]*listNode
	lru         *listNode // pointer to least recently used node
	mru         *listNode // pointer to most recently used node
}

func New(capacity int) Cache {
	if capacity < 1 {
		panic("capacity must be at least 1")
	}
	return &cache{
		capacity: capacity,
		listMap:  make(map[string]*listNode, capacity),
	}
}

func (c *cache) Put(k, v string) {
	newNode := &listNode{
		val:  &v,
		next: nil,
	}

	if c.mru != nil {
		c.mru.next = newNode
	}
	c.mru = newNode

	if c.currentSize == 0 { // initialise
		c.lru = newNode
	}

	if c.currentSize == c.capacity {
		if next := c.lru.next; next != nil {
			//*c.lru = *c.lru.next
			temp := c.lru
			c.lru = next
			*temp = listNode{}
		}
	}

	if node, isPresent := c.listMap[k]; isPresent {
		node.val = &v
		c.mru = node
	}

	c.listMap[k] = newNode
	if c.currentSize != c.capacity {
		c.currentSize++
	}
}

func (c *cache) Get(k string) (*string, error) {
	if node, isPresent := c.listMap[k]; isPresent {
		if node.val == nil {
			delete(c.listMap, k) // clean up
			return nil, noElementError{}
		}
		c.mru.next = node
		c.mru = node
		// update the lru node when you Get() the lru node
		if next := c.lru.next; next != nil && node == c.lru {
			c.lru = next
		}
		c.mru.next = nil
		strCopy := *node.val
		return &strCopy, nil
	}
	return nil, noElementError{}
}

type noElementError struct{}

func (e noElementError) Error() string {
	return "no element exists for the given key"
}

type listNode struct {
	val  *string
	next *listNode
}
