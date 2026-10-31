package p0146

// 哈希表 + 双向链表：表头最新，表尾最久未用
type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node // 哨兵，head.next 是最新
	tail     *node // 哨兵，tail.prev 是最旧
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next, tail.prev = tail, head
	return LRUCache{capacity: capacity, cache: map[int]*node{}, head: head, tail: tail}
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.moveToFront(n)
	return n.val
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.cache[key]; ok {
		n.val = value
		c.moveToFront(n)
		return
	}
	n := &node{key: key, val: value}
	c.cache[key] = n
	c.pushFront(n)
	if len(c.cache) > c.capacity {
		old := c.tail.prev
		c.remove(old)
		delete(c.cache, old.key) // 节点里存 key 就是为了这里
	}
}

func (c *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) pushFront(n *node) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) moveToFront(n *node) {
	c.remove(n)
	c.pushFront(n)
}
