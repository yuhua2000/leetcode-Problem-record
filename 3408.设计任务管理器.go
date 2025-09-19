package leetcode

import (
	"container/heap"
)

type Task struct {
	priority int
	taskId   int
}

type TaskHp []*Task

func (a TaskHp) Len() int { return len(a) }

func (a TaskHp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a TaskHp) Less(i, j int) bool {
	if a[i].priority == a[j].priority {
		return a[i].taskId > a[j].taskId
	}
	return a[i].priority > a[j].priority
}

func (a *TaskHp) Pop() any {
	n := a.Len()
	result := (*a)[n-1]
	*a = (*a)[:n-1]
	return result
}

func (a *TaskHp) Push(item any) {
	task := item.(*Task)
	*a = append(*a, task)
}

type TaskManager struct {
	taskHeap *TaskHp
	taskMap  map[int][2]int
}

func Constructor(tasks [][]int) TaskManager {
	taskHeap := make(TaskHp, 0)
	taskMap := make(map[int][2]int)
	for _, item := range tasks {
		user, taskId, priority := item[0], item[1], item[2]
		task := &Task{priority: priority, taskId: taskId}
		taskMap[taskId] = [2]int{priority, user}
		heap.Push(&taskHeap, task)
	}
	return TaskManager{taskHeap: &taskHeap, taskMap: taskMap}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{priority: priority, taskId: taskId}
	this.taskMap[taskId] = [2]int{priority, userId}
	heap.Push(this.taskHeap, task)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	if info, ok := this.taskMap[taskId]; ok {
		info[0] = newPriority
		this.taskMap[taskId] = info
		heap.Push(this.taskHeap, &Task{priority: newPriority, taskId: taskId})
	}
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskMap, taskId)
}

func (this *TaskManager) ExecTop() int {
	for this.taskHeap.Len() > 0 {
		task := heap.Pop(this.taskHeap).(*Task)
		if value, ok := this.taskMap[task.taskId]; ok && value[0] == task.priority {
			delete(this.taskMap, task.taskId)
			return value[1]
		}
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
