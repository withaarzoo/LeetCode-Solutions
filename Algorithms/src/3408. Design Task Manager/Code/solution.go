package main

import (
	"container/heap"
)

// Entry represents a task in the priority queue.
type Entry struct {
	priority int
	taskId   int
}

// MaxHeap implements a max-heap for Entries.
// Higher priority first; if tie, larger taskId first.
type MaxHeap []Entry

func (h MaxHeap) Len() int      { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId > h[j].taskId
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Entry))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// TaskManager structure
type TaskManager struct {
	mp   map[int][2]int // taskId -> [userId, priority]
	heap MaxHeap
}

// Constructor initializes TaskManager with given tasks
func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		mp:   make(map[int][2]int),
		heap: MaxHeap{},
	}
	for _, t := range tasks {
		if len(t) < 3 {
			continue
		}
		userId, taskId, pr := t[0], t[1], t[2]
		tm.mp[taskId] = [2]int{userId, pr}
		heap.Push(&tm.heap, Entry{priority: pr, taskId: taskId})
	}
	return tm
}

// Add a new task
func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.mp[taskId] = [2]int{userId, priority}
	heap.Push(&this.heap, Entry{priority: priority, taskId: taskId})
}

// Edit modifies the priority of an existing task
func (this *TaskManager) Edit(taskId int, newPriority int) {
	rec := this.mp[taskId]
	this.mp[taskId] = [2]int{rec[0], newPriority}
	heap.Push(&this.heap, Entry{priority: newPriority, taskId: taskId})
}

// Rmv removes a task
func (this *TaskManager) Rmv(taskId int) {
	delete(this.mp, taskId)
}

// ExecTop executes the highest priority task and returns userId
func (this *TaskManager) ExecTop() int {
	for this.heap.Len() > 0 {
		top := heap.Pop(&this.heap).(Entry)
		rec, ok := this.mp[top.taskId]
		if !ok {
			continue // task removed
		}
		if rec[1] != top.priority {
			continue // stale entry
		}
		delete(this.mp, top.taskId)
		return rec[0]
	}
	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
