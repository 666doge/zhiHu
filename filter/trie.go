package filter

import (
	"strings"
)

type Node struct{
	char rune
	deepth int
	children map[rune]*Node
	isTerminal bool
}

type Trie struct{
	root *Node
}

func newNode() *Node{
	return &Node{
		children: make(map[rune]*Node, 32),
	}
}

func newTrie() *Trie{
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) addWord (word string) {
	str := strings.TrimSpace(word)
	chars := []rune(str)
	node := t.root
	for _, c := range chars {
		child, ok := node.children[c]
		if !ok {
			child = newNode()
			child.char = c
			child.deepth = node.deepth + 1
			node.children[c] = child
		}
		node = child
	}
	node.isTerminal = true
	return
}

func (t *Trie) filterSensitiveWord (str string) (result string){
	if t.root == nil {
		return
	}

	chars := []rune(str)
	newChars := []rune{}
	node := t.root
	for _, c := range chars {
		child, ok := node.children[c]
		if ok {
			if child.isTerminal == true {
				newChars = append(newChars, []rune("***")...)
			}
			node = child
		} else {
			newChars = append(newChars, c)
			node = t.root
		}
	}
	result = string(newChars)
	return
}

