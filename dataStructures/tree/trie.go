package tree

type trieNode struct {
	isWord bool // 是否是单词结尾
	next   map[rune]*trieNode
}
type trie struct {
	size int // 节点数量
	root *trieNode
}

func NewTrie() *trie {
	return &trie{
		0,
		&trieNode{false, make(map[rune]*trieNode)},
	}
}
