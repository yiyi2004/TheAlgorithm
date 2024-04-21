package main

import (
	"container/list"
)

func main() {

}

//type entry struct {
//	key, value, freq int // freq 表示这本书被看的次数
//}
//
//type LFUCache struct {
//	capacity   int
//	minFreq    int
//	keyToNode  map[int]*list.Element
//	freqToList map[int]*list.List
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		capacity:   capacity,
//		keyToNode:  map[int]*list.Element{},
//		freqToList: map[int]*list.List{},
//	}
//}
//
//func (c *LFUCache) pushFront(e *entry) {
//	if _, ok := c.freqToList[e.freq]; !ok {
//		c.freqToList[e.freq] = list.New() // 双向链表
//	}
//	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
//}
//
//func (c *LFUCache) getEntry(key int) *entry {
//	// 拿到这个节点的缓存
//	node := c.keyToNode[key]
//	if node == nil { // 没有这本书
//		return nil
//	}
//	// 拿到节点和对应列表，先从原来的列表中删除
//	// 然后添加到新的列表上
//	// 需要注意 minFreq
//	e := node.Value.(*entry)
//	lst := c.freqToList[e.freq]
//	lst.Remove(node)    // 把这本书抽出来
//	if lst.Len() == 0 { // 抽出来后，这摞书是空的
//		delete(c.freqToList, e.freq) // 移除空链表
//		if c.minFreq == e.freq {     // 这摞书是最左边的
//			c.minFreq++
//		}
//	}
//
//	e.freq++       // 看书次数 +1
//	c.pushFront(e) // 放在右边这摞书的最上面
//	return e
//}
//
//func (c *LFUCache) Get(key int) int {
//	if e := c.getEntry(key); e != nil { // 有这本书
//		return e.value
//	}
//	return -1 // 没有这本书
//}
//
//func (c *LFUCache) Put(key, value int) {
//	if e := c.getEntry(key); e != nil { // 有这本书
//		e.value = value // 更新 value
//		return
//	}
//	if len(c.keyToNode) == c.capacity { // 书太多了
//		lst := c.freqToList[c.minFreq]                           // 最左边那摞书
//		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key) // 移除这摞书的最下面的书
//		if lst.Len() == 0 {                                      // 这摞书是空的
//			delete(c.freqToList, c.minFreq) // 移除空链表
//		}
//	}
//	c.pushFront(&entry{key, value, 1}) // 新书放在「看过 1 次」的最上面
//	c.minFreq = 1
//}

type LFUCache struct {
	minFreq    int
	capacity   int
	keyToNode  map[int]*list.Element
	freqToList map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		keyToNode:  make(map[int]*list.Element),
		freqToList: make(map[int]*list.List),
	}
}

type entry struct {
	key, val, freq int
}

func (this *LFUCache) pushFront(e *entry) {
	if _, ok := this.freqToList[e.freq]; !ok {
		this.freqToList[e.freq] = list.New()
	}
	this.keyToNode[e.key] = this.freqToList[e.freq].PushFront(e)
}

func (this *LFUCache) getEntry(key int) *entry {
	if node, ok := this.keyToNode[key]; ok {
		e := node.Value.(*entry)
		lst := this.freqToList[e.freq]
		lst.Remove(node)
		if lst.Len() == 0 {
			delete(this.freqToList, e.freq)
			if this.minFreq == e.freq {
				this.minFreq++
			}
		}
		e.freq++
		this.pushFront(e)
		return e
	}
	return nil
}

func (this *LFUCache) Get(key int) int {
	node := this.getEntry(key)
	if node == nil {
		return -1
	}
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	node := this.getEntry(key)
	if node != nil {
		node.val = value
		return
	}
	if len(this.keyToNode) == this.capacity {
		lst := this.freqToList[this.minFreq]
		delete(this.keyToNode, lst.Remove(lst.Back()).(*entry).key)
		if lst.Len() == 0 {
			delete(this.freqToList, this.minFreq)
		}
	}
	this.pushFront(&entry{key, value, 1})
	this.minFreq = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
