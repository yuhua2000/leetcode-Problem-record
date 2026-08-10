package leetcode

func remainingMethods(n int, k int, invocations [][]int) []int {
	callGraph := make(map[int][]int)
	for _, invocation := range invocations {
		callGraph[invocation[0]] = append(callGraph[invocation[0]], invocation[1])
	}

	suspiciousMethods := make(map[int]bool)
	suspiciousMethods[k] = true
	queue := []int{k}
	for len(queue) > 0 {
		methodID := queue[0]
		queue = queue[1:]

		for _, calledMethodID := range callGraph[methodID] {
			if suspiciousMethods[calledMethodID] {
				continue
			}
			suspiciousMethods[calledMethodID] = true
			queue = append(queue, calledMethodID)
		}
	}

	result := make([]int, 0, n)
	for _, invocation := range invocations {
		callerID := invocation[0]
		calledID := invocation[1]

		_, callerSuspicious := suspiciousMethods[callerID]
		_, calledSuspicious := suspiciousMethods[calledID]

		if !callerSuspicious && calledSuspicious {
			for i := 0; i < n; i++ {
				result = append(result, i)
			}
			return result
		}

	}

	for i := 0; i < n; i++ {
		if suspiciousMethods[i] {
			continue
		}
		result = append(result, i)
	}

	return result

}
