package leetcode

import (
	"container/heap"
	"slices"
	"strconv"
	"strings"
)

type offlineUser struct {
	onlineAt int
	userID   int
}

type offlineUserHeap []offlineUser

func (o offlineUserHeap) Len() int { return len(o) }

func (o offlineUserHeap) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

func (o offlineUserHeap) Less(i, j int) bool { return o[i].onlineAt < o[j].onlineAt }

func (o *offlineUserHeap) Push(x any) {
	*o = append(*o, x.(offlineUser))
}

func (o *offlineUserHeap) Pop() any {
	v := (*o)[o.Len()-1]
	*o = (*o)[:len(*o)-1]
	return v
}

func countMentions(numberOfUsers int, events [][]string) []int {
	slices.SortFunc(events, func(a, b []string) int {
		aNum, _ := strconv.Atoi(a[1])
		bNum, _ := strconv.Atoi(b[1])
		if aNum == bNum {
			return strings.Compare(b[0], a[0])
		}
		return aNum - bNum
	})

	offlineUntil := make(map[int]struct{}, numberOfUsers/6)
	offlineHeap := &offlineUserHeap{}
	result := make([]int, numberOfUsers)
	allMentionCount := 0

	for _, event := range events {
		timestamp, _ := strconv.Atoi(event[1])
		switch event[0] {
		case "MESSAGE":
			switch event[2] {
			case "ALL":
				allMentionCount++
			case "HERE":
				for offlineHeap.Len() > 0 && (*offlineHeap)[0].onlineAt <= timestamp {
					user := heap.Pop(offlineHeap).(offlineUser)
					delete(offlineUntil, user.userID)
				}

				for userID := range result {
					if _, ok := offlineUntil[userID]; ok {
						continue
					}
					result[userID]++
				}
			default:
				for _, token := range strings.Split(event[2], " ") {
					id, _ := strconv.Atoi(token[2:])

					result[id]++
				}
			}
		case "OFFLINE":
			for offlineHeap.Len() > 0 && (*offlineHeap)[0].onlineAt <= timestamp {
				user := heap.Pop(offlineHeap).(offlineUser)
				delete(offlineUntil, user.userID)
			}

			userId, _ := strconv.Atoi(event[2])
			offlineUntil[userId] = struct{}{}
			heap.Push(offlineHeap, offlineUser{userID: userId, onlineAt: timestamp + 60})
		}
	}

	for i := 0; i < len(result) && allMentionCount > 0; i++ {
		result[i] += allMentionCount
	}

	return result
}
