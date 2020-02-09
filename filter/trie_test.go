package filter

import (
	"testing"
	"fmt"
)

func TestAddWord(t *testing.T) {
	trie := newTrie()
	trie.addWord("二狗")
	trie.addWord("呵呵")
	result := trie.filterSensitiveWord("hi,二狗, 吃饭啦！！！呵呵")
	fmt.Println(result)
}
