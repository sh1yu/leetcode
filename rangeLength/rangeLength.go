package main

import (
	"sort"

	"github.com/sh1yu/assertion"
)

type Task struct {
	Start int
	End   int
}

type TaskSlice []*Task

func (t TaskSlice) Len() int {
	return len(t)
}

func (t TaskSlice) Less(i, j int) bool {
	return t[i].Start < t[j].Start
}

func (t TaskSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	assertion.Equals(0, getSum([]*Task{}))
	assertion.Equals(3, getSum([]*Task{{Start: 1, End: 4}}))
	assertion.Equals(5, getSum([]*Task{{Start: 1, End: 4}, {Start: 6, End: 8}}))
	assertion.Equals(4, getSum([]*Task{{Start: 1, End: 4}, {Start: 2, End: 5}}))
	assertion.Equals(6, getSum([]*Task{{Start: 6, End: 8}, {Start: 1, End: 4}, {Start: 2, End: 5}}))
}

func getSum(tasks []*Task) int {

	if len(tasks) == 0 {
		return 0
	}
	if len(tasks) == 1 {
		return tasks[0].End - tasks[0].Start
	}

	sort.Sort(TaskSlice(tasks))

	sum := 0
	preMinStart := tasks[0].Start
	preMaxEnd := tasks[0].End

	for i := 1; i < len(tasks); i++ {

		// 有交叉
		if tasks[i].Start <= preMaxEnd {
			preMinStart = min(preMinStart, tasks[i].Start)
			preMaxEnd = max(preMaxEnd, tasks[i].End)
			continue
		}

		// 没有交叉
		sum += preMaxEnd - preMinStart
		preMinStart = tasks[i].Start
		preMaxEnd = tasks[i].End
	}

	return sum + (preMaxEnd - preMinStart)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
