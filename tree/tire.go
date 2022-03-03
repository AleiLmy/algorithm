/**
 * @Author: xuelei02
 * @Description:
 * @File:  tire
 * @Version: 1.0.0
 * @Date: 2022/3/3 20:41
 */

package tree

type TrieNode struct {
	Children [26]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}

func (t *TrieNode) insert(word string) {
	node := t
	for _, v := range word {
		idx := v - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &TrieNode{}
		}
		node = node.Children[idx]
	}
}

// SearchPrefix
// todo 搜索前缀啊
func (t *TrieNode) SearchPrefix(s string) {

}
