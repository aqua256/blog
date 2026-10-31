package p0208

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		i := ch - 'a'
		if node.children[i] == nil {
			node.children[i] = &Trie{}
		}
		node = node.children[i]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.walk(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

// walk 沿路径走到底，走不通返回 nil
func (t *Trie) walk(s string) *Trie {
	node := t
	for _, ch := range s {
		node = node.children[ch-'a']
		if node == nil {
			return nil
		}
	}
	return node
}
