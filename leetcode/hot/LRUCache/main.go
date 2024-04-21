package main

import (
	"fmt"
	"sync"
)

func main() {
	sync.Mutex{}
	lruCache := Constructor(1)
	lruCache.Put(2, 1)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(3, 2)
	fmt.Println(lruCache.Get(2))
	fmt.Println(lruCache.Get(3))
}

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key   int
	value int
	pre   *DLinkNode
	next  *DLinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkNode, capacity),
		head:     new(DLinkNode),
		tail:     new(DLinkNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		if node == this.head {
			return node.value
		}
		if node == this.tail {
			this.tail = this.tail.pre
			node.next = this.head
			this.head.pre = node
			this.head = node
			return node.value
		}
		if node.pre != nil {
			node.pre.next = node.next
		}
		if node.next != nil {
			node.next.pre = node.pre
		}
		this.head.pre = node
		node.next = this.head
		this.head = node
		return node.value

	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// 如果存在
		// 修改 cache 中的值
		node.value = value
		if node == this.head {
			return
		}
		if node == this.tail {
			this.tail = node.pre
			this.tail.next = nil
			node.next = this.head
			this.head.pre = node
			this.head = node
			return
		}
		// 移到前面
		if node.pre != nil {
			node.pre.next = node.next
		}
		if node.next != nil {
			node.next.pre = node.pre
		}
		this.head.pre = node
		node.next = this.head
		this.head = node
		return
	}
	// 判断目前 LRU 是否已经满了 如果满了，那就淘汰。
	// 如果没满，插入第一个位置
	if this.capacity == this.size {
		// 删除尾部节点
		delete(this.cache, this.tail.key)
		if this.size == 1 {
			newHead := &DLinkNode{
				key:   key,
				value: value,
				next:  this.head,
			}
			this.head = newHead
			this.tail = newHead
			this.cache[key] = newHead
			return
		}
		node := this.tail
		this.tail = node.pre
		this.tail.next = nil
		node.pre = nil
		// 添加头节点
		newHead := &DLinkNode{
			key:   key,
			value: value,
			next:  this.head,
		}
		this.head.pre = newHead
		this.head = newHead
		this.cache[key] = newHead
	} else {
		newHead := &DLinkNode{
			key:   key,
			value: value,
		}
		if this.size == 0 {
			this.head = newHead
			this.tail = newHead
		} else {
			newHead.next = this.head
			this.head.pre = newHead
			this.head = newHead
		}
		this.cache[key] = newHead
		this.size++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 实现一个 LRUCache 并不是困难的事情

//type LRUCache struct {
//	size       int
//	cap        int
//	cache      map[int]*DNode
//	head, tail *DNode
//}
//
//type DNode struct {
//	key, val  int
//	pre, next *DNode
//}
//
//func Constructor(capacity int) LRUCache {
//	head, tail := new(DNode), new(DNode)
//	head.next = tail
//	tail.pre = head
//
//	return LRUCache{
//		cap:   capacity,
//		cache: make(map[int]*DNode),
//		head:  head,
//		tail:  tail,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	res, ok := this.cache[key]
//	if !ok {
//		return -1
//	}
//	resVal := res.val
//	res.pre.next = res.next
//	res.next.pre = res.pre
//	res.pre = this.tail.pre
//	this.tail.pre.next = res
//	this.tail.pre = res
//	this.tail.next = nil
//	return resVal
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	c, ok := this.cache[key]
//	if ok {
//		c.val = value
//		c.pre.next = c.next
//		c.next.pre = c.pre
//		this.tail.pre.next = c
//		c.pre = this.tail.pre
//		c.next = this.tail
//		this.tail.pre = c
//		return
//	} else if !ok && this.size == this.cap {
//		// 移除头节点
//		head := this.head.next
//		delete(this.cache, head.key)
//
//		this.head.next = head.next
//		head.key = key
//		head.val = value
//		this.cache[head.key] = head
//
//		head.pre = this.tail.pre
//		head.next = this.tail
//		this.tail.pre = head
//		this.tail.next = nil
//		return
//	}
//	this.size++
//	n := &DNode{
//		key:  key,
//		val:  value,
//		pre:  nil,
//		next: nil,
//	}
//	this.cache[key] = n
//	this.tail.pre.next = n
//	n.pre = this.tail.pre
//	n.next = this.tail
//	this.tail.pre = n
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 对比看一下你的垃圾代码
//type LRUCache struct {
//	size int
//	capacity int
//	cache map[int]*DLinkedNode
//	head, tail *DLinkedNode
//}
//
//type DLinkedNode struct {
//	key, value int
//	prev, next *DLinkedNode
//}
//
//func initDLinkedNode(key, value int) *DLinkedNode {
//	return &DLinkedNode{
//		key: key,
//		value: value,
//	}
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		cache: map[int]*DLinkedNode{},
//		head: initDLinkedNode(0, 0),
//		tail: initDLinkedNode(0, 0),
//		capacity: capacity,
//	}
//	l.head.next = l.tail
//	l.tail.prev = l.head
//	return l
//}
//
//func (this *LRUCache) Get(key int) int {
//	if _, ok := this.cache[key]; !ok {
//		return -1
//	}
//	node := this.cache[key]
//	this.moveToHead(node)
//	return node.value
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	if _, ok := this.cache[key]; !ok {
//		node := initDLinkedNode(key, value)
//		this.cache[key] = node
//		this.addToHead(node)
//		this.size++
//		if this.size > this.capacity {
//			removed := this.removeTail()
//			delete(this.cache, removed.key)
//			this.size--
//		}
//	} else {
//		node := this.cache[key]
//		node.value = value
//		this.moveToHead(node)
//	}
//}
//
//func (this *LRUCache) addToHead(node *DLinkedNode) {
//	node.prev = this.head
//	node.next = this.head.next
//	this.head.next.prev = node
//	this.head.next = node
//}
//
//func (this *LRUCache) removeNode(node *DLinkedNode) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (this *LRUCache) moveToHead(node *DLinkedNode) {
//	this.removeNode(node)
//	this.addToHead(node)
//}
//
//func (this *LRUCache) removeTail() *DLinkedNode {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}
