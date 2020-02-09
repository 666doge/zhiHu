package filter

import (
	"os"
	"io"
	"bufio"
)

var trie *Trie

func Init(filename string) (err error) {
	trie = newTrie()
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, e := reader.ReadString('\n')
		if e == io.EOF {
			return
		}

		if e != nil {
			err = e
			return
		}
		trie.addWord(str)
	}
	return
}
func Filter(str string) (newStr string) {
	return trie.filterSensitiveWord(str)
}