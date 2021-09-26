package main

import "fmt"

type Trie struct {
	dic   [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range []byte(word) {
		index := c - 'a'
		if cur.dic[index] == nil {
			newTrie := Constructor()
			cur.dic[index] = &newTrie
		}
		cur = cur.dic[index]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range []byte(word) {
		index := c - 'a'
		if cur.dic[index] == nil {
			return false
		}
		cur = cur.dic[index]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range []byte(prefix) {
		index := c - 'a'
		if cur.dic[index] == nil {
			return false
		}
		cur = cur.dic[index]
	}
	return true
}

func main() {
	obj := Constructor()
	obj.Insert("trie")
	param_2 := obj.Search("trie")
	param_3 := obj.StartsWith("tr")
	fmt.Println(param_2, param_3)
}
