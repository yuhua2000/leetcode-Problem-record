package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=2512 lang=golang
 *
 * [2512] 奖励最顶尖的 K 名学生
 */

// @lc code=start
func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	type pair struct{ score, id int }
	score := make([]pair, len(student_id))
	words := make(map[string]int)

	for _, s := range positive_feedback {
		words[s] = 3
	}
	for _, s := range negative_feedback {
		words[s] = -1
	}

	for i := 0; i < len(report); i++ {
		score[i] = pair{score: 0, id: student_id[i]}
		for _, s := range strings.Split(report[i], " ") {
			score[i].score += words[s]
		}
	}

	sort.Slice(score, func(i, j int) bool {
		if score[i].score == score[j].score {
			return score[i].id < score[j].id
		}
		return score[i].score > score[j].score
	})

	for i, p := range score[:k] {
		student_id[i] = p.id
	}
	return student_id[:k]
}

// @lc code=end
