package leetcode

type QueueItem struct {
	index int
	steps int
}

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{0, 1})

	sameValueIndices := map[int][]int{}
	for i, v := range arr {
		if i > 0 && i < len(arr)-1 {
			if arr[i] == arr[i-1] && arr[i] == arr[i+1] {
				continue
			}
		}
		sameValueIndices[v] = append(sameValueIndices[v], i)
	}

	visited := make(map[int]bool)
	visited[0] = true

	getNextIndices := func(i int) []int {
		result := sameValueIndices[arr[i]]

		// 这个没有会超时
		sameValueIndices[arr[i]] = nil
		if i > 1 {
			result = append(result, i-1)
		}
		if i < len(arr)-1 {
			result = append(result, i+1)
		}

		return result
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, nextIndex := range getNextIndices(current.index) {
			if nextIndex == len(arr)-1 {
				return current.steps
			}

			if visited[nextIndex] {
				continue
			}
			visited[nextIndex] = true

			queue = append(queue, QueueItem{nextIndex, current.steps + 1})
		}

	}

	return -1
}
