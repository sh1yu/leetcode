package main

import "github.com/sh1yu/assertion"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	dis := make(map[string]int)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := make([]string, 1, len(wordSet))
	queue[0] = beginWord
	wordSet[beginWord] = false
	dis[beginWord] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == endWord {
			return dis[cur]
		}

		for i := 0; i < len(cur); i++ {
			for c := 'a'; c <= 'z'; c++ {
				newCur := cur[:i] + string(c) + cur[i+1:]
				if wordSet[newCur] {
					queue = append(queue, newCur)
					wordSet[newCur] = false
					dis[newCur] = dis[cur] + 1
				}
			}
		}
	}

	return 0
}

func main() {
	assertion.Equals(5, ladderLength("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
