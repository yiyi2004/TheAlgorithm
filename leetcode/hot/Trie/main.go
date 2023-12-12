package main

func main() {

}

type Trie struct {
	trees map[byte]*node
}

type node struct {
	v      byte
	isLeaf bool
	next   *node
}

func Constructor() Trie {
	return Trie{
		trees: make(map[byte]*node),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	// 已经没有耐心了，明天再写吧。
	root := this.trees[word[0]]
	for i := 0; i < len(word); i++ {
		if root == nil {
			root = &node{
				v:    word[i],
				next: nil,
			}
		}
		if root.v == word[i] {
			root = root.next

		}
	}
}

func (this *Trie) Search(word string) bool {

}

func (this *Trie) StartsWith(prefix string) bool {

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
