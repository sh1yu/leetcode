package main

import "github.com/sh1yu/assertion"

func minMutation(start string, end string, bank []string) int {

	libSet := make(map[string]bool)
	for _, s := range bank {
		libSet[s] = true
	}

	if !libSet[end] {
		return -1
	}

	dis := make(map[string]int)
	queue := make([]string, 1, len(bank))
	queue[0] = start
	libSet[start] = false
	dis[start] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			return dis[cur]
		}

		for i := 0; i < len(cur); i++ {
			for _, c := range []byte{'A', 'C', 'G', 'T'} {
				newCur := cur[:i] + string(c) + cur[i+1:]
				if libSet[newCur] {
					queue = append(queue, newCur)
					libSet[newCur] = false
					dis[newCur] = dis[cur] + 1
				}
			}
		}
	}

	return -1
}

func main() {
	assertion.Equals(2, minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}
